package internal_service

import "golang.org/x/crypto/bcrypt"

type service struct {
}

func New() *service {
	return &service{}
}

func (s *service) GenerateHashText(text []byte, cost int) ([]byte, error) {
	return bcrypt.GenerateFromPassword(text, cost)
}
