package account

import (
	"cmp"
	"context"
	"slices"

	"github.com/hashicorp/golang-lru/v2/expirable"
	"github.com/insan1a/tech-tinker/internal/domain/interfaces"
	"github.com/insan1a/tech-tinker/internal/domain/model"
)

var _ interfaces.AccountService = (*Service)(nil)

type Service struct {
	employeeRepo interfaces.EmployeeRepo
	customerRepo interfaces.CustomerRepo
	orderRepo    interfaces.OrderRepo
	cache        *expirable.LRU[string, *model.Employee]
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
		cache:        expirable.NewLRU[string, *model.Employee](cfg.Cache.Size, nil, cfg.Cache.TTL),
	}
}

// GetAccount implements interfaces.AccountService.
func (s *Service) GetAccount(ctx context.Context, id string) (*model.Employee, error) {
	employee, ok := s.cache.Get(id)
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

	s.cache.Add(id, employee)

	return employee, nil
}

func (s *Service) GetOrders(ctx context.Context, id string) ([]model.Order, error) {
	employee, ok := s.cache.Get(id)
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
	order, err := s.orderRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (s *Service) GetStatistic(ctx context.Context, id string) (*model.Stat, error) {
	return nil, nil
}
