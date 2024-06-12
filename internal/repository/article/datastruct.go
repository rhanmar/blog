package article

import "time"

type Article struct {
	ID        int64     `db:"id"`
	Title     string    `db:"title"`
	Text      string    `db:"text"`
	CreatedAt time.Time `db:"created_at"`
}
