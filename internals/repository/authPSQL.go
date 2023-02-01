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

func (sp AuthPSQL) CreateUser(user entity.User) error                  { return nil }
func (sp AuthPSQL) GetUser(mail, password string) (entity.User, error) { return entity.User{}, nil }
