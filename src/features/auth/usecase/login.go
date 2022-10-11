package auth_usecase

import (
	"class_main_service/src/app"
	auth_struct "class_main_service/src/features/auth/struct"
	"context"
	"errors"
	"github.com/go-redis/redis/v9"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"time"
)

func (u *usecase) Login(ctx context.Context, input auth_struct.LoginInput) (*auth_struct.LoginOutput, error) {
	ctx, span := tracer.Start(ctx, "start login")
	defer span.End()

	ctx, span = tracer.Start(ctx, "query in email")
	user, err := u.repository.FindByEmail(ctx, input.Email)
	span.End()
	if err != nil {
		zap.S().Error(err)
		return nil, err
	}

	ctx, span = tracer.Start(ctx, "compare password")
	checkPassword := u.internalService.CompareHashText(input.Password, user.Password)
	span.End()
	if !checkPassword {
		err = app.BadRequestHttpError("username or password not valid", errors.New("username or password not valid"))
		zap.S().Error(err)
		return nil, err
	}

	version, err := u.repository.GetVersionUser(ctx, user.Id)
	if err != nil && errors.Is(err, redis.Nil) {
		zap.S().Error("nil")
	} else {
		zap.S().Error(err)
		return nil, err
	}

	claims := userClaims{
		user.Id,
		version,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(60 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        string(user.Id),
		},
	}
	ctx, span = tracer.Start(ctx, "generate jwt token")
	token, err := u.internalService.GenerateJwtToken(u.secretKey, claims)
	span.End()
	if err != nil {
		zap.S().Error(err)
		return nil, err
	}

	output := auth_struct.LoginOutput{
		AccessToken: token,
		User:        user,
	}

	return &output, nil
}
