package helpers

import (
	"time"
)

type RateLimiter struct {
	Limit  int
	Cursor string
	Reset  time.Time
}
