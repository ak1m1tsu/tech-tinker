package model

type Order struct {
	Base
	Number     uint
	PriceLimit uint
	Comment    string
	Address    string
	Status     OrderStatus

	Customer       *Customer
	Employee       *Employee
	Configurations Configurations
}

type OrderStatus uint

func (s OrderStatus) String() string {
	return orderTypeNames[s]
}

const (
	OrderStatus_InProcess OrderStatus = iota
	OrderStatus_Completed
)

var orderTypeNames = map[OrderStatus]string{
	OrderStatus_InProcess: "In process",
	OrderStatus_Completed: "Completed",
}
