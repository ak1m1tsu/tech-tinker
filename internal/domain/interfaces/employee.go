package interfaces

import (
	"context"

	"github.com/insan1a/tech-tinker/internal/domain/model"
)

type EmployeeRepo interface {
	GetByEmail(ctx context.Context, email string) (*model.Employee, error)
	GetByID(ctx context.Context, id string) (*model.Employee, error)
}

type EmployeeService interface {
	GetByEmail(ctx context.Context, email string) (*model.Employee, error)
	GetByID(ctx context.Context, id string) (*model.Employee, error)
}
