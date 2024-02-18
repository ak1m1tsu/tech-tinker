package model

type Order struct {
	Base
	Number     int         `db:"number"`
	PriceLimit int         `db:"price_limit"`
	Comment    string      `db:"comment"`
	Address    string      `db:"address"`
	Status     OrderStatus `db:"status"`

	CustomerID string `db:"customer_id"`
	Customer   *Customer

	EmployeeID string `db:"employee_id"`
	Employee   *Employee

	Configurations Configurations
}

type Orders []Order

type OrderStatus uint8

func (s OrderStatus) String() string {
	return orderTypeNames[s]
}

const (
	OrderStatusInProcess OrderStatus = iota
	OrderStatusCompleted
)

var orderTypeNames = map[OrderStatus]string{
	OrderStatusInProcess: "In process",
	OrderStatusCompleted: "Completed",
}
