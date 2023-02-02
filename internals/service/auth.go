package service

import (
	"crypto/sha256"
	"encoding/hex"
	"posts/internals/entity"
	"posts/internals/repository"
)

type AuthService struct {
	rep repository.Authorization
}

var (
	salt = []byte{175, 92, 205, 93, 46, 123, 82, 117, 134, 164, 43, 33, 129, 115, 60, 103}
)

func newAuthService(rep repository.Authorization) *AuthService {
	return &AuthService{rep: rep}
}

func (s AuthService) CreateUser(user entity.User) (int, error) {
	user.Password = generateHash(user.Password)
	return s.rep.CreateUser(user)
}
func (s AuthService) GetUser(mail, password string) (entity.User, error) { return entity.User{}, nil }

func generateHash(s string) string {
	var Hasher = sha256.New()
	passwordBytes := []byte(s)
	passwordBytes = append(passwordBytes, salt...)
	Hasher.Write(passwordBytes)
	hashedPassword := Hasher.Sum(nil)

	return hex.EncodeToString(hashedPassword)
}
