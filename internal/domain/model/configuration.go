package model

type Configuration struct {
	Base
	Price int `db:"price"`

	OrderID string `db:"order_id"`
	Order   *Order

	EmployeeID string `db:"employee_id"`
	Employee   *Employee

	Components Components
}

type Configurations []Configuration
