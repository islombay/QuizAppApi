package service

import (
	"QuizAppApi/pkg/repository"
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	salt       = "vj39ghJDFS3huwgs"
	tokenTTL   = 12 * time.Hour
	signingKey = "23g83bfw8fbwbbfifiwnbhwiwri33492hd828dh8f1"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserRole string `json:"user-role"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) GenerateAdminToken(user, password string) (string, error) {
	_, err := s.repo.GetAdmin(user, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		"admin",
	})
	return token.SignedString([]byte(signingKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
