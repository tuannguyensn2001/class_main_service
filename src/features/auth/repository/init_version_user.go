package auth_repository

import (
	"context"
	"fmt"
)

func (r *repository) InitVersionUser(ctx context.Context, userId int) error {
	return r.rd.Set(ctx, fmt.Sprintf("auth_user_%d", userId), 1, 0).Err()
}
