package auth_v1

import (
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	desc "github.com/markgenuine/auth/pkg/auth_v1"
)

const (
	// Users ...
	Users = "users"

	// UsersID ...
	UsersID = "id"

	// UsersName ...
	UsersName = "name"

	// UsersEmail ...
	UsersEmail = "email"

	// UsersPassword ...
	UsersPassword = "password"

	// UsersPasswordConfirm ...
	UsersPasswordConfirm = "password_confirm"

	// UsersRole ...
	UsersRole = "role"

	// UsersCreatedAt ...
	UsersCreatedAt = "created_at"

	// UsersUpdatedAt ...
	UsersUpdatedAt = "updated_at"
)

// User - type for proto implementation
type User struct {
	desc.UnimplementedUserV1Server

	poolDB *pgx.Conn
	sq     squirrel.StatementBuilderType
}

// NewUserService create proto interface implementation
func NewUserService(conn *pgx.Conn, sqIn squirrel.StatementBuilderType) *User {
	return &User{poolDB: conn, sq: sqIn}
}
