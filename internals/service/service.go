package service

import (
	"posts/internals/entity"
	"posts/internals/repository"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	SignIn(username, password string) (string, error)
	GetUser(id int) (entity.User, error)
	ParseToken(token string) (int, error)
}

type UserPost interface {
	CreatePost(userId int, post entity.Post) error
	GetPostById(postId int) (*entity.Post, error)
	GetUserPostsRange(userId int, part int) ([]entity.Post, error)
	GetPostAnswers(postId int) ([]entity.Post, error)
}

type Service struct {
	Authorization
	UserPost
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		Authorization: newAuthService(rep.Authorization),
		UserPost:      newUserPostService(rep.UserPost),
	}
}
