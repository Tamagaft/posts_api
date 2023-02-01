package service

import (
	"posts/internals/entity"
	"posts/internals/repository"
)

type AuthService struct {
	rep repository.Authorization
}

func newAuthService(rep repository.Authorization) *AuthService {
	return &AuthService{rep: rep}
}

func (as AuthService) CreateUser(user entity.User) error                  { return nil }
func (as AuthService) GetUser(mail, password string) (entity.User, error) { return entity.User{}, nil }
