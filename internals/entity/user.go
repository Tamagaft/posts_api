package entity

type User struct {
	Id          int    `db:"is"`
	Username    string `db:"username"`
	Description string `db:"description"`
	Password    string `db:"password"`
}
