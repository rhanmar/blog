package article

import (
	"time"
)

type Article struct {
	ID        int64
	Title     string
	Text      string
	CreatedAt time.Time
}
