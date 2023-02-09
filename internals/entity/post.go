package entity

import "time"

type Post struct {
	Id     int       `db:"id"`
	Text   string    `db:"text"`
	Author int       `db:"author_pk"`
	Date   time.Time `db:"date"`
}
