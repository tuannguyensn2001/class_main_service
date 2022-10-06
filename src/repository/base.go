package repository

import (
	"context"
	"gorm.io/gorm"
)

const key = "db"

type Repository struct {
}

func (r *Repository) GetDBFromContext(ctx context.Context, db *gorm.DB) *gorm.DB {
	if val, ok := ctx.Value(key).(*gorm.DB); ok {
		return val
	}
	return db
}
