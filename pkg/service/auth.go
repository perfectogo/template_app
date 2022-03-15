package service

import (
	"crypto/sha1"
	"fmt"

	template "github.com/perfectogo/template_app"
	"github.com/perfectogo/template_app/pkg/repository"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

// service create user
func (s *AuthService) CreateUser(user template.User) (int, error) {
	fmt.Println(user)
	user.Password = GeneratePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

// password hesher func
func GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte("fizpongpingbuz")))
}
