package interfaces

import (
	"context"

	"github.com/ak1m1tsu/tech-tinker/internal/domain/model"
)

type CustomerRepo interface {
	GetManyByOrderIDs(ctx context.Context, ids []string) ([]model.Customer, error)
}
