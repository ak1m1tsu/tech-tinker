package interfaces

import (
	"context"

	"github.com/ak1m1tsu/tech-tinker/internal/domain/model"
)

type OrderRepo interface {
	GetByEmployeeID(ctx context.Context, employeeID string) ([]model.Order, error)
	GetByID(ctx context.Context, id string) (*model.Order, error)
}
