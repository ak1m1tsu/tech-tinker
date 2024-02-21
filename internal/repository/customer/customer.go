package customer

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/ak1m1tsu/go-libs/connector/postgresql"
	"github.com/insan1a/tech-tinker/internal/domain/interfaces"
	"github.com/insan1a/tech-tinker/internal/domain/model"
)

var _ interfaces.CustomerRepo = &Repo{}

type Repo struct {
	conn *postgresql.Connection
}

func New(conn *postgresql.Connection) *Repo {
	return &Repo{conn: conn}
}

func (r *Repo) GetManyByOrderIDs(ctx context.Context, ids []string) ([]model.Customer, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	sql, args, err := psql.
		Select("id", "first_name", "last_name", "email", "phone_number").
		From("customer").
		Where(sq.Eq{"id": ids}).
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.conn.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []model.Customer
	var customer model.Customer
	for rows.Next() {
		if err = rows.Scan(
			&customer.ID,
			&customer.FirstName,
			&customer.LastName,
			&customer.Email,
			&customer.PhoneNumber,
		); err != nil {
			break
		}

		customers = append(customers, customer)
	}

	if err != nil {
		return nil, err
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return customers, nil
}
