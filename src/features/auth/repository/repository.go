package auth_repository

import (
	auth_model "class_main_service/src/features/auth/model"
	repository2 "class_main_service/src/repository"
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"
	"strconv"
)

type repository struct {
	*repository2.Repository
	db *gorm.DB
	rd *redis.Client
}

func New(db *gorm.DB, rd *redis.Client) *repository {
	return &repository{db: db, rd: rd}
}

func (r *repository) Create(ctx context.Context, user *auth_model.User) error {
	db := r.GetDBFromContext(ctx, r.db)
	return db.Create(user).WithContext(ctx).Error
}

func (r *repository) FindByEmail(ctx context.Context, email string) (*auth_model.User, error) {
	db := r.GetDBFromContext(ctx, r.db)
	var result auth_model.User
	if err := db.WithContext(ctx).Where("email = ?", email).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *repository) GetVersionUser(ctx context.Context, userId int) (int, error) {
	val, err := r.rd.Get(ctx, fmt.Sprintf("auth_user_%d", userId)).Result()
	if err != nil {
		return -1, err
	}
	result, err := strconv.Atoi(val)
	if err != nil {
		return -1, err
	}

	return result, nil
}
