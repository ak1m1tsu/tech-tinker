package account

import "errors"

var (
	ErrOrderNotFound    = errors.New("the order not found")
	ErrInvalidDateRange = errors.New("invalid date range")
)
