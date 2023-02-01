package entity

type User struct {
	id          int    `db:"is"`
	username    string `db:"username"`
	description string `db:"description"`
	password    string `db:"password"`
}
