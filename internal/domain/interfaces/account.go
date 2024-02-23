package interfaces

import (
	"context"

	"github.com/ak1m1tsu/tech-tinker/internal/domain/model"
)

type AccountService interface {
	GetAccount(ctx context.Context, id string) (*model.Employee, error)
	GetOrders(ctx context.Context, id string) ([]model.Order, error)
	GetOrder(ctx context.Context, id string) (*model.Order, error)
	GetStatistic(ctx context.Context, id string) (*model.Stat, error)
}
