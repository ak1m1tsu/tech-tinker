package account

import (
	"cmp"
	"context"
	"errors"
	"fmt"
	"slices"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/ak1m1tsu/tech-tinker/internal/domain/interfaces"
	"github.com/ak1m1tsu/tech-tinker/internal/domain/model"
	"github.com/hashicorp/golang-lru/v2/expirable"
)

var _ interfaces.AccountService = (*Service)(nil)

type Service struct {
	employeeRepo interfaces.EmployeeRepo
	customerRepo interfaces.CustomerRepo
	orderRepo    interfaces.OrderRepo
	errors       *expirable.LRU[string, error]
	employees    *expirable.LRU[string, *model.Employee]
	stats        *expirable.LRU[string, *model.Stat]
}

func New(cfg *Config,
	employeeRepo interfaces.EmployeeRepo,
	customerRepo interfaces.CustomerRepo,
	orderRepo interfaces.OrderRepo,
) *Service {
	return &Service{
		employeeRepo: employeeRepo,
		customerRepo: customerRepo,
		orderRepo:    orderRepo,
		errors:       expirable.NewLRU[string, error](cfg.Cache.Size, nil, cfg.Cache.TTL),
		employees:    expirable.NewLRU[string, *model.Employee](cfg.Cache.Size, nil, cfg.Cache.TTL),
		stats:        expirable.NewLRU[string, *model.Stat](cfg.Cache.Size, nil, cfg.Cache.TTL),
	}
}

// GetAccount implements interfaces.AccountService.
func (s *Service) GetAccount(ctx context.Context, id string) (*model.Employee, error) {
	employee, ok := s.employees.Get(id)
	if ok {
		return employee, nil
	}

	employee, err := s.employeeRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	orders, err := s.orderRepo.GetByEmployeeID(ctx, id)
	if err != nil {
		return nil, err
	}

	slices.SortFunc(orders, func(a, b model.Order) int {
		return cmp.Compare(a.Status, b.Status)
	})

	customersIDs := make([]string, 0, len(orders))
	for _, order := range orders {
		customersIDs = append(customersIDs, order.CustomerID)
	}

	customers, err := s.customerRepo.GetManyByOrderIDs(ctx, customersIDs)
	if err != nil {
		return nil, err
	}

	for _, order := range orders {
		for _, customer := range customers {
			if order.CustomerID == customer.ID {
				order.Customer = &customer
				break
			}
		}
	}

	employee.Orders = orders

	s.employees.Add(id, employee)

	return employee, nil
}

func (s *Service) GetOrders(ctx context.Context, id string) ([]model.Order, error) {
	employee, ok := s.employees.Get(id)
	if ok {
		return employee.Orders, nil
	}

	orders, err := s.orderRepo.GetByEmployeeID(ctx, id)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (s *Service) GetOrder(ctx context.Context, id string) (*model.Order, error) {
	if err, ok := s.errors.Get(id); ok {
		return nil, err
	}

	order, err := s.orderRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = errors.Join(err, ErrOrderNotFound)
			s.errors.Add(id, err)
			return nil, err
		}

		return nil, err
	}

	return order, nil
}

func (s *Service) GetStatistic(ctx context.Context, id string, from, to time.Time) (*model.Stat, error) {
	if stat, ok := s.stats.Get(fmt.Sprintf("%s-%s", from, to)); ok {
		return stat, nil
	}

	if to.Before(from) {
		return nil, ErrInvalidDateRange
	}

	if to.Sub(from).Hours() > 24*31 {
		return nil, ErrInvalidDateRange
	}

	orders, err := s.orderRepo.GetByEmployeeID(ctx, id)
	if err != nil {
		return nil, err
	}

	budgets := model.Budgets{
		{Type: model.BudgetTypeLowerThan50K, Count: 0},
		{Type: model.BudgetTypeBetween50KAnd100K, Count: 0},
		{Type: model.BudgetTypeBetween100KAnd500K, Count: 0},
		{Type: model.BudgetTypeGreaterThan500K, Count: 0},
	}
	total := 0
	for _, order := range orders {
		switch {
		case order.PriceLimit < 50000:
			budgets[0].Count++
		case order.PriceLimit >= 50000 && order.PriceLimit < 100000:
			budgets[1].Count++
		case order.PriceLimit >= 100000 && order.PriceLimit < 500000:
			budgets[2].Count++
		case order.PriceLimit >= 500000:
			budgets[3].Count++
		}
		total += order.PriceLimit
	}

	stat := &model.Stat{
		From:    from,
		To:      to,
		Budgets: budgets,
		Total:   total,
	}

	s.stats.Add(fmt.Sprintf("%s-%s", from, to), stat)

	return stat, nil
}
