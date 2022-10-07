package internal_service

import (
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
}

func New() *service {
	return &service{}
}

func (s *service) GenerateHashText(text []byte, cost int) ([]byte, error) {
	return bcrypt.GenerateFromPassword(text, cost)
}

func (s *service) CompareHashText(text string, hashText string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashText), []byte(text))
	if err != nil {
		return false
	}

	return true
}

func (s *service) GenerateJwtToken(secretKey string, claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return result, nil
}
