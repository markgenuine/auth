package auth

import (
	"context"
	"fmt"
	"log"

	"github.com/markgenuine/auth/internal/client/db"
	"github.com/markgenuine/auth/internal/model"
)

func (r *repo) Create(ctx context.Context, user *model.UserCreate) (int64, error) {
	query, args, err := r.sq.Insert(users).
		Columns(usersName, usersEmail, usersPassword, usersRole).
		Values(user.Name, user.Email, user.Password, user.Role).
		Suffix(fmt.Sprintf("RETURNING %s", usersID)).ToSql()

	if err != nil {
		log.Printf("failed to build query create user: %s", err.Error())
		return 0, err
	}

	q := db.Query{
		Name:     "auth_repository.Create",
		QueryRaw: query,
	}

	var userID int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&userID)
	if err != nil {
		log.Printf("failed to insert user: %s", err.Error())
		return 0, err
	}

	return userID, err
}
