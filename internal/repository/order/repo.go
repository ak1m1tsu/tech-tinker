package order

import (
	"context"
	"github.com/sirupsen/logrus"

	sq "github.com/Masterminds/squirrel"
	"github.com/ak1m1tsu/go-libs/connector/postgresql"
	"github.com/insan1a/tech-tinker/internal/domain/interfaces"
	"github.com/insan1a/tech-tinker/internal/domain/model"
)

var _ interfaces.OrderRepo = &Repo{}

type Repo struct {
	conn *postgresql.Connection
}

func New(conn *postgresql.Connection) *Repo {
	return &Repo{conn: conn}
}

func (r *Repo) GetByEmployeeID(ctx context.Context, employeeID string) ([]model.Order, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	sql, args, err := psql.
		Select("id", "customer_id", "number", "price_limit", "comment", "address", "status", "created_at").
		From("\"order\"").
		Where(sq.And{
			sq.Eq{"employee_id": employeeID},
			sq.Eq{"deleted_at": nil},
		}).
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.conn.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []model.Order
	var order model.Order
	for rows.Next() {
		if err = rows.Scan(
			&order.ID,
			&order.CustomerID,
			&order.Number,
			&order.PriceLimit,
			&order.Comment,
			&order.Address,
			&order.Status,
			&order.CreatedAt,
		); err != nil {
			break
		}

		orders = append(orders, order)
	}

	if err != nil {
		return nil, err
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return orders, nil
}

func (r *Repo) GetByID(ctx context.Context, id string) (*model.Order, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	sql, args, err := psql.
		Select("id", "employee_id", "customer_id", "number", "price_limit", "comment", "address", "status", "created_at").
		From("\"order\"").
		Where(sq.And{
			sq.Eq{"id": id},
			sq.Eq{"deleted_at": nil},
		}).
		ToSql()
	if err != nil {
		return nil, err
	}

	logrus.Debugf("Query: %s, args: %v", sql, args)

	var order model.Order
	if err = r.conn.QueryRow(ctx, sql, args...).Scan(
		&order.ID,
		&order.EmployeeID,
		&order.CustomerID,
		&order.Number,
		&order.PriceLimit,
		&order.Comment,
		&order.Address,
		&order.Status,
		&order.CreatedAt,
	); err != nil {
		return nil, err
	}

	return &order, nil
}
