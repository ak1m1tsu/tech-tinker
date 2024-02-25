package interfaces

import (
	"context"

	"github.com/ak1m1tsu/tech-tinker/internal/domain/model"
)

type AuthService interface {
	GetByEmail(ctx context.Context, email string) (*model.Employee, error)
	CreateToken(ctx context.Context, e *model.Employee) (string, error)
}
