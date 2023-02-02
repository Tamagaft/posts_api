package entity

type User struct {
	Id          int    `db:"is" json:"id"`
	Username    string `db:"username" json:"username"`
	Description string `db:"description" json:"description"`
	Password    string `db:"password" json:"password"`
}
