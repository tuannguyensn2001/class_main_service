package auth_repository

import (
	auth_model "class_main_service/src/features/auth/model"
	repository2 "class_main_service/src/repository"
	"context"
	"gorm.io/gorm"
)

type repository struct {
	*repository2.Repository
	db *gorm.DB
}

func New(db *gorm.DB) *repository {
	return &repository{db: db}
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
