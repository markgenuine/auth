package auth

import (
	"github.com/Masterminds/squirrel"
	"github.com/markgenuine/auth/internal/repository"
	"github.com/markgenuine/platform_common/pkg/db"
)

const (
	users          = "users"
	usersID        = "id"
	usersName      = "name"
	usersEmail     = "email"
	usersPassword  = "password"
	usersRole      = "role"
	usersCreatedAt = "created_at"
	usersUpdatedAt = "updated_at"
)

type repo struct {
	db db.Client
	sq squirrel.StatementBuilderType
}

// NewRepository ...
func NewRepository(db db.Client) repository.UserRepository {
	return &repo{
		db: db,
		sq: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
