package auth_usecase

import "github.com/golang-jwt/jwt/v4"

type userClaims struct {
	Id      int
	Version int
	jwt.RegisteredClaims
}
