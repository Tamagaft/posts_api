package service

import (
	"crypto/sha256"
	"encoding/hex"
	"posts/internals/entity"
	"posts/internals/repository"
	"time"

	"github.com/golang-jwt/jwt"
)

type AuthService struct {
	rep repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

var (
	tokenTTL     = 12 * time.Hour
	salt         = []byte{175, 92, 205, 93, 46, 123, 82, 117, 134, 164, 43, 33, 129, 115, 60, 103}
	jwtSignature = []byte{30, 35, 64, 136, 241, 119, 92, 15, 1, 120, 2, 245, 91, 253, 29, 129}
)

func newAuthService(rep repository.Authorization) *AuthService {
	return &AuthService{rep: rep}
}

func (s AuthService) CreateUser(user entity.User) (int, error) {
	user.Password = generateHash(user.Password)
	return s.rep.CreateUser(user)
}

func (s AuthService) SignIn(username, password string) (string, error) {
	user, err := s.rep.GetUser(username, generateHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString(jwtSignature)
}

func (s AuthService) GetUser(id int) (entity.User, error) { return entity.User{}, nil }

func generateHash(s string) string {
	var Hasher = sha256.New()
	passwordBytes := []byte(s)
	passwordBytes = append(passwordBytes, salt...)
	Hasher.Write(passwordBytes)
	hashedPassword := Hasher.Sum(nil)

	return hex.EncodeToString(hashedPassword)
}
