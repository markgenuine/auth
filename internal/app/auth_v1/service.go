package auth_v1

import (
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	desc "github.com/markgenuine/auth/pkg/auth_v1"
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
