package model

import "golang.org/x/crypto/bcrypt"

type Employee struct {
	Base
	FirstName      string
	LastName       string
	Email          string
	Role           EmployeeRole
	HashedPassword HashedPassword
}

type HashedPassword []byte

func (hp HashedPassword) FromPassword(password string) {
	bcryptPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hp = HashedPassword(bcryptPassword)
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
	Manager EmployeeRole = iota
	Technician
	Administrator
)

var employeeRoleNames = map[EmployeeRole]string{
	Manager:       "Manager",
	Technician:    "Technician",
	Administrator: "Administrator",
}
