package repository

import (
	"posts/internals/entity"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user entity.User) error
	GetUser(mail, password string) (entity.User, error)
}

type UserPost interface {
	CreatePost(post entity.Post) error
	GetPostById(postId string) (*entity.Post, error)
	GetUserPostsRange(userId, part int) ([]entity.Post, error)
}

type Repository struct {
	Authorization
	UserPost
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPSQL(db),
		UserPost:      NewUserPostPSQL(db),
	}
}
