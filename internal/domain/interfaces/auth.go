package interfaces

import (
	"context"

	"github.com/insan1a/tech-tinker/internal/domain/model"
)

type AuthService interface {
	EmployeeRepo
	CreateToken(ctx context.Context, e *model.Employee) (string, error)
}
