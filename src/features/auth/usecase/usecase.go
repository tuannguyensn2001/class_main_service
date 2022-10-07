package auth_usecase

import (
	auth_model "class_main_service/src/features/auth/model"
	"context"
	"github.com/golang-jwt/jwt/v4"
	"go.opentelemetry.io/otel"
)

const cost = 5

var tracer = otel.Tracer("auth_usecase")

type IRepository interface {
	Create(ctx context.Context, user *auth_model.User) error
	FindByEmail(ctx context.Context, email string) (*auth_model.User, error)
}

type IInternalService interface {
	GenerateHashText(text []byte, cost int) ([]byte, error)
	CompareHashText(text string, hashText string) bool
	GenerateJwtToken(secretKey string, claims jwt.Claims) (string, error)
}

type usecase struct {
	repository      IRepository
	internalService IInternalService
	secretKey       string
}

func New(repository IRepository, internalService IInternalService, secretKey string) *usecase {
	return &usecase{repository: repository, internalService: internalService, secretKey: secretKey}
}
