package auth_usecase

import (
	"class_main_service/src/app"
	auth_model "class_main_service/src/features/auth/model"
	auth_struct "class_main_service/src/features/auth/struct"
	"context"
	"errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (u *usecase) Register(ctx context.Context, input auth_struct.RegisterInput) error {
	ctx, span := tracer.Start(ctx, "start register")
	defer span.End()

	ctx, span = tracer.Start(ctx, "find by email", trace.WithAttributes(attribute.String("email", input.Email)))
	user, err := u.repository.FindByEmail(ctx, input.Email)
	span.End()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if user != nil {
		err = app.BadRequestHttpError("user existed", errors.New("user existed"))
		zap.S().Error(err)
		return err
	}

	ctx, span = tracer.Start(ctx, "generate password", trace.WithAttributes(attribute.String("password", input.Password)))
	password, err := u.internalService.GenerateHashText([]byte(input.Password), cost)
	span.End()
	if err != nil {
		zap.S().Error(err)
	}

	newUser := auth_model.User{
		Username: input.Username,
		Email:    input.Email,
		Password: string(password),
	}

	ctx, span = tracer.Start(ctx, "insert user to db")
	err = u.repository.Create(ctx, &newUser)
	span.End()
	if err != nil {
		zap.S().Error(err)
	}

	return nil

}
