package main

import (
	"context"
	"github.com/urfave/cli/v2"
	"os"
	"sort"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/ak1m1tsu/go-libs/connector/postgresql"
	"github.com/ak1m1tsu/tech-tinker/internal/domain/model"
	"github.com/go-faker/faker/v4"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	dsn := os.Getenv("POSTGRESQL_URL")
	if dsn == "" {
		logrus.Fatal("POSTGRESQL_URL is not set")
	}

	fixture, err := New(dsn)
	if err != nil {
		logrus.Fatal(err)
	}
	defer fixture.Close()
	logrus.Info("Connected to database")

	if fixture.Run() != nil {
		logrus.Fatal(err)
	}
}

func run() error {
	app := &cli.App{}

	return app.Run(os.Args)
}

type Fixture struct {
	cli  *cli.App
	conn *postgresql.Connection
	log  *logrus.Logger
}

func New(dsn string) (*Fixture, error) {
	conn, err := postgresql.Connect(dsn, nil)
	if err != nil {
		return nil, err
	}

	f := &Fixture{
		cli: &cli.App{
			Name:  "fixture",
			Usage: "generate fixtures",
		},
		log:  logrus.New(),
		conn: conn,
	}

	f.cli.Commands = []*cli.Command{
		f.employee(),
	}

	return f, nil
}

func (f *Fixture) Run() error {
	sort.Sort(cli.FlagsByName(f.cli.Flags))
	sort.Sort(cli.CommandsByName(f.cli.Commands))
	return f.cli.Run(os.Args)
}

func (f *Fixture) Close() {
	f.conn.Close()
}

func (f *Fixture) employee() *cli.Command {
	return &cli.Command{
		Name:  "employee",
		Usage: "generate a random employee",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "role",
				Aliases: []string{"r"},
				Value:   model.EmployeeRoleTechnician.String(),
				Usage:   "set employee role",
			},
			&cli.StringFlag{
				Name:    "password",
				Aliases: []string{"p"},
				Usage:   "set employee password",
			},
		},
		Action: func(ctx *cli.Context) error {
			role := strings.ToLower(ctx.String("role"))
			if _, ok := model.EmployeeRoleValues[role]; !ok {
				role = model.EmployeeRoleTechnician.String()
			}

			password := ctx.String("password")
			if password == "" {
				password = faker.Password()
			}

			f.log.Info("Password: ", password)

			hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			account := model.Employee{
				FirstName:      faker.FirstName(),
				LastName:       faker.LastName(),
				Email:          faker.Email(),
				Role:           model.EmployeeRoleValues[role],
				HashedPassword: hash,
			}

			sql, args, err := squirrel.Insert("employee").
				Columns("first_name", "last_name", "email", "role", "password").
				Values(account.FirstName, account.LastName, account.Email, account.Role, string(account.HashedPassword)).
				PlaceholderFormat(squirrel.Dollar).
				ToSql()
			if err != nil {
				return err
			}

			f.log.Info("SQL query: ", sql)
			f.log.Info("Query args: ", args)

			if _, err = f.conn.Exec(context.Background(), sql, args...); err != nil {
				return err
			}

			logrus.Info("Account created")

			return nil
		},
	}
}
