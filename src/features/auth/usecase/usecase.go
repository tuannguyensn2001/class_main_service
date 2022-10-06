package auth_usecase

import (
	auth_model "class_main_service/src/features/auth/model"
	"context"
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
}

type usecase struct {
	repository      IRepository
	internalService IInternalService
}

func New(repository IRepository, internalService IInternalService) *usecase {
	return &usecase{repository: repository, internalService: internalService}
}
