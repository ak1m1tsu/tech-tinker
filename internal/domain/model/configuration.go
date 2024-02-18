package model

type Configuration struct {
	Base
	Price int

	OrderID string
	Order   *Order

	EmployeeID string
	Employee   *Employee

	Components Components
}

type Configurations []Configuration
