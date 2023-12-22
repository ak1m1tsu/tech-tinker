package model

type Configuration struct {
	Base
	Price      int
	Order      *Order
	Employee   *Employee
	Components Components
}

type Configurations []Configuration
