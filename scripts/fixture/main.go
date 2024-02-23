package main

import (
	"context"
	"os"

	"github.com/Masterminds/squirrel"
	"github.com/ak1m1tsu/go-libs/connector/postgresql"
	"github.com/ak1m1tsu/tech-tinker/internal/domain/model"
	"github.com/go-faker/faker/v4"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		logrus.Fatal("DB_URL is not set")
	}

	password := faker.Password()
	logrus.Info("Password: ", password)

	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	account := model.Employee{
		FirstName:      faker.FirstName(),
		LastName:       faker.LastName(),
		Email:          faker.Email(),
		Role:           model.EmployeeRoleTechnician,
		HashedPassword: hash,
	}

	conn, err := postgresql.Connect(dsn, nil)
	if err != nil {
		logrus.Fatal(err)
	}
	defer conn.Close()

	logrus.Info("Connected to database")

	sql, args, err := squirrel.Insert("employee").
		Columns("first_name", "last_name", "email", "role", "password").
		Values(account.FirstName, account.LastName, account.Email, account.Role, string(account.HashedPassword)).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("SQL query: ", sql)
	logrus.Info("Query args: ", args)

	if _, err = conn.Exec(context.Background(), sql, args...); err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("Account created")
}
