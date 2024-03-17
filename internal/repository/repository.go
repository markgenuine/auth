package repository

import (
	"context"

	"github.com/markgenuine/auth/internal/model"
	modelRepo "github.com/markgenuine/auth/internal/repository/auth/model"
)

// UserRepository ...
type UserRepository interface {
	Create(ctx context.Context, request *model.UserCreate) (int64, error)
	Get(ctx context.Context, id int64) (*modelRepo.User, error)
	Update(ctx context.Context, request *model.UserUpdate) error
	Delete(ctx context.Context, id int64) error
}
