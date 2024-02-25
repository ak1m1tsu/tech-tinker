package employee

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/ak1m1tsu/go-libs/connector/postgresql"
	"github.com/ak1m1tsu/tech-tinker/internal/domain/interfaces"
	"github.com/ak1m1tsu/tech-tinker/internal/domain/model"
	"github.com/sirupsen/logrus"
)

var _ interfaces.EmployeeRepo = &Repo{}

type Repo struct {
	conn *postgresql.Connection
}

func New(conn *postgresql.Connection) *Repo {
	return &Repo{conn: conn}
}

// GetByEmail implements interfaces.EmployeeRepo.
func (r *Repo) GetByEmail(ctx context.Context, email string) (*model.Employee, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	sql, args, err := psql.
		Select("id", "email", "first_name", "last_name", "role", "password", "created_at", "updated_at", "deleted_at").
		From("employee").
		Where(sq.Eq{"email": email}).
		ToSql()
	if err != nil {
		return nil, err
	}

	logrus.Debug(sql, args)

	var e model.Employee
	if err = r.conn.QueryRow(ctx, sql, args...).Scan(
		&e.ID,
		&e.Email,
		&e.FirstName,
		&e.LastName,
		&e.Role,
		&e.HashedPassword,
		&e.CreatedAt,
		&e.UpdatedAt,
		&e.DeletedAt,
	); err != nil {
		return nil, err
	}

	return &e, nil
}

func (r *Repo) GetByID(ctx context.Context, id string) (*model.Employee, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	sql, args, err := psql.
		Select("id", "email", "first_name", "last_name", "role", "password", "created_at").
		From("employee").
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return nil, err
	}

	var employee model.Employee
	if err = r.conn.QueryRow(ctx, sql, args...).Scan(
		&employee.ID,
		&employee.Email,
		&employee.FirstName,
		&employee.LastName,
		&employee.Role,
		&employee.HashedPassword,
		&employee.CreatedAt,
	); err != nil {
		return nil, err
	}

	return &employee, nil
}
