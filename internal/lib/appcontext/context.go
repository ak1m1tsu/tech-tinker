package appcontext

import "context"

type key string

const (
	EmployeeID key = "employee_id"
)

// WithEmployeeID store employee ID as a string into context.
func WithEmployeeID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, EmployeeID, id)
}

// GetEmployeeID returns employee ID as a string from context if it exists.
func GetEmployeeID(ctx context.Context) string {
	id, ok := ctx.Value(EmployeeID).(string)
	if !ok {
		return ""
	}

	return id
}
