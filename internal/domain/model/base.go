package model

import "time"

type Base struct {
	ID        string
	CreatedAt time.Time
	UpdateAt  time.Time
	DeletedAt time.Time
}
