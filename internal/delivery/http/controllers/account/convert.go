package account

import (
	"time"

	"github.com/ak1m1tsu/tech-tinker/internal/domain/model"
)

func FillInfo(employee *model.Employee) InfoResponse {
	return InfoResponse{
		ID:        employee.ID,
		FirstName: employee.FirstName,
		LastName:  employee.LastName,
		Email:     employee.Email,
		Role:      employee.Role.String(),
		CreatedAt: employee.CreatedAt.Format(time.RFC3339),
		Orders:    FillOrders(employee.Orders),
	}
}

func FillOrder(order *model.Order) OrderResponse {
	return OrderResponse{
		ID:         order.ID,
		Number:     order.Number,
		PriceLimit: order.PriceLimit,
		Comment:    order.Comment,
		Address:    order.Address,
		Status:     order.Status.String(),
		CreatedAt:  order.CreatedAt.Format(time.RFC3339),
	}
}

func FillOrders(orders []model.Order) OrderResponses {
	responses := make(OrderResponses, 0, len(orders))
	for _, order := range orders {
		responses = append(responses, FillOrder(&order))
	}
	return responses
}

func FillStatistic(stat *model.Stat) StatisticResponse {
	return StatisticResponse{
		From:    stat.From.Format(time.RFC3339),
		To:      stat.From.Format(time.RFC3339),
		Total:   stat.Total,
		Budgets: FillBudgets(stat.Budgets),
	}
}

func FillBudgets(budgets model.Budgets) BudgetResponses {
	responses := make(BudgetResponses, 0, len(budgets))
	for _, budget := range budgets {
		responses = append(responses, FillBudget(budget))
	}
	return responses
}

func FillBudget(budget model.Budget) BudgetResponse {
	return BudgetResponse{
		Type:  budget.Type.String(),
		Count: budget.Count,
	}
}
