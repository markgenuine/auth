package service

import (
	"context"

	"github.com/markgenuine/auth/internal/model"
)

// AuthService ...
type AuthService interface {
	Create(ctx context.Context, user *model.UserCreate) (int64, error)
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*model.User, error)
	Update(ctx context.Context, user *model.UserUpdate) error
}
