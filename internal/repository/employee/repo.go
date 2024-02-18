package employee

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/ak1m1tsu/go-libs/connector/postgresql"
	"github.com/insan1a/tech-tinker/internal/domain/interfaces"
	"github.com/insan1a/tech-tinker/internal/domain/model"
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
		Select("id", "email", "first_name", "last_name", "role", "hashed_password", "created_at", "updated_at", "deleted_at").
		From("employee").
		Where(sq.Eq{"email": email}).
		ToSql()
	if err != nil {
		return nil, err
	}

	var e model.Employee
	if err = r.conn.QueryRow(ctx, sql, args).Scan(
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
