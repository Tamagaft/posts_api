package repository

import (
	"posts/internals/entity"

	"github.com/jmoiron/sqlx"
)

type AuthPSQL struct {
	db *sqlx.DB
}

func NewAuthPSQL(db *sqlx.DB) *AuthPSQL {
	return &AuthPSQL{db: db}
}

func (r AuthPSQL) CreateUser(user entity.User) (int, error) {
	var userId int
	stmt, err := r.db.Prepare("INSERT INTO users(username,password) VALUES($1,$2) RETURNING id")
	if err != nil {
		return userId, err
	}
	row := stmt.QueryRow(user.Username, user.Password)
	if err := row.Scan(&userId); err != nil {
		return userId, err
	}
	return userId, nil
}
func (r AuthPSQL) ChangeDescription(user entity.User) error {
	stmt, err := r.db.Prepare("UPDATE users SET description=$1 where id=$2")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.Description, user.Id)
	if err != nil {
		return err
	}
	return nil
}
func (r AuthPSQL) GetUser(username, password string) (*entity.User, error) {
	var user entity.User
	stmt, err := r.db.Prepare("SELECT id,Username,COALESCE(description,'') FROM users WHERE username=$1 AND password=$2")
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow(username, password)
	if err = row.Scan(&user.Id, &user.Username, &user.Description); err != nil {
		return nil, err
	}
	return &user, nil
}
