package util

import (
	"context"

	"github.com/ak1m1tsu/tech-tinker/internal/domain/model"
)

func EmployeeFromContext(ctx context.Context) *model.Employee {
	employee, ok := ctx.Value("employee").(*model.Employee)
	if !ok {
		return nil
	}

	return employee
}
