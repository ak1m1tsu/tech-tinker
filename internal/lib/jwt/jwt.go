package jwt

import (
	"crypto/rsa"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type Employee struct {
	ID   string `json:"id"`
	Role string `json:"role"`
}

func GenerateToken(employee *Employee, ttl time.Duration, key *rsa.PrivateKey) (string, error) {
	now := time.Now().UTC()
	claims := &Claims{
		EmployeeID:   employee.ID,
		EmployeeRole: employee.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ID:        uuid.NewString(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return "", err
	}

	return token, nil
}

func ValidateToken(token string, key *rsa.PublicKey) (*Claims, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return key, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(*Claims)
	if !ok || !parsedToken.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}
