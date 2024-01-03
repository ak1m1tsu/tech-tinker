package model

import "golang.org/x/crypto/bcrypt"

type Employee struct {
	Base
	FirstName      string
	LastName       string
	Email          string
	Role           EmployeeRole
	HashedPassword HashedPassword

	Orders Orders
}

type Employees []Employee

type HashedPassword []byte

func (hp *HashedPassword) FromPassword(password string) {
	*hp, _ = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (hp HashedPassword) Compare(password string) bool {
	err := bcrypt.CompareHashAndPassword(hp, []byte(password))
	return err == nil
}

type EmployeeRole uint8

func (e EmployeeRole) String() string {
	return employeeRoleNames[e]
}

const (
	EmployeeRoleManager EmployeeRole = iota
	EmployeeRoleTechnician
	EmployeeRoleAdministrator
)

var employeeRoleNames = map[EmployeeRole]string{
	EmployeeRoleManager:       "Manager",
	EmployeeRoleTechnician:    "Technician",
	EmployeeRoleAdministrator: "Administrator",
}
