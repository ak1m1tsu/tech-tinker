package interfaces

import (
	"context"

	"github.com/insan1a/tech-tinker/internal/domain/model"
)

type CustomerRepo interface {
	GetManyByOrderIDs(ctx context.Context, ids []string) ([]model.Customer, error)
}
