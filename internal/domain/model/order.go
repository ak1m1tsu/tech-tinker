package model

type Order struct {
	Base
	Number     int         `json:"number"`
	PriceLimit int         `json:"price_limit"`
	Comment    string      `json:"comment"`
	Address    string      `json:"address"`
	Status     OrderStatus `json:"status"`

	CustomerID string    `json:"customer_id"`
	Customer   *Customer `json:"-"`

	EmployeeID string    `json:"employee_id"`
	Employee   *Employee `json:"-"`

	Configurations Configurations `json:"configurations,omitempty"`
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
