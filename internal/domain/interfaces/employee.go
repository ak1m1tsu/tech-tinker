package interfaces

import (
	"context"

	"github.com/ak1m1tsu/tech-tinker/internal/domain/model"
)

type EmployeeRepo interface {
	GetByEmail(ctx context.Context, email string) (*model.Employee, error)
	GetByID(ctx context.Context, id string) (*model.Employee, error)
}

type EmployeeService interface {
	GetByEmail(ctx context.Context, email string) (*model.Employee, error)
	GetByID(ctx context.Context, id string) (*model.Employee, error)
}
