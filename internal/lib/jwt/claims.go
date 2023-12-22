package jwt

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	EmployeeID   string `json:"employeeID"`
	EmployeeRole string `json:"employeeRole"`
	jwt.RegisteredClaims
}
