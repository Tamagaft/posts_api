package service

import (
	"posts/internals/entity"
	"posts/internals/repository"
)

type UserPostService struct {
	rep repository.UserPost
}

func newUserPostService(rep repository.UserPost) *UserPostService {
	return &UserPostService{rep: rep}
}

func (ups UserPostService) CreatePost(userId string, post entity.Post) error { return nil }
func (ups UserPostService) GetPostById(postId string) (entity.Post, error)   { return entity.Post{}, nil }
func (ups UserPostService) GetUserPostsRange(part int) ([]entity.Post, error) {
	return []entity.Post{}, nil
}
func (ups UserPostService) GetPostAnswers(postId string) ([]entity.Post, error) {
	return []entity.Post{}, nil
}
