package util

import (
	"context"
	"github.com/insan1a/tech-tinker/internal/domain/model"
)

func EmployeeFromContext(ctx context.Context) *model.Employee {
	employee, ok := ctx.Value("employee").(*model.Employee)
	if !ok {
		return nil
	}

	return employee
}
