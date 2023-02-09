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

func (s UserPostService) CreatePost(userId int, post entity.Post) error {
	post.Author = userId
	return s.rep.CreatePost(post)
}
func (s UserPostService) GetPostById(postId int) (*entity.Post, error) {
	return s.rep.GetPostById(postId)
}
func (s UserPostService) GetUserPostsRange(userId int, part int) ([]entity.Post, error) {
	return s.rep.GetUserPostsRange(userId, part)
}
func (s UserPostService) GetPostAnswers(postId int) ([]entity.Post, error) {
	return []entity.Post{}, nil
}
