package account

import "time"

type Config struct {
	Cache struct {
		Size int
		TTL  time.Duration
	}
}
