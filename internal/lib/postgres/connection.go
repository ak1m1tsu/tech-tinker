package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Conn struct {
	maxPoolSize  int
	connAttempts int
	connTimeout  time.Duration

	*pgxpool.Pool
}

func New(url string, opts ...Option) (*Conn, error) {
	conn := &Conn{
		maxPoolSize:  defaultPoolSize,
		connAttempts: defaultAttempts,
		connTimeout:  defaultTimeout,
	}

	for _, opt := range opts {
		opt(conn)
	}

	cfg, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, err
	}

	cfg.MaxConns = int32(conn.maxPoolSize)

	for i := 0; i < conn.connAttempts; i++ {
		conn.Pool, err = pgxpool.NewWithConfig(context.Background(), cfg)
		if err == nil {
			return conn, nil
		}

		time.Sleep(conn.connTimeout)
		continue
	}

	return nil, err
}
