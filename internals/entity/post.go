package entity

import "time"

type Post struct {
	id     int       `db:"id"`
	Text   string    `db:"text"`
	Author string    `db:"author_pk"`
	Date   time.Time `db:"date"`
}
