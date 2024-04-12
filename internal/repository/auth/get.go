package auth

import (
	"context"
	"errors"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	modelRepo "github.com/markgenuine/auth/internal/repository/auth/model"
	"github.com/markgenuine/platform_common/pkg/db"
)

func (r *repo) Get(ctx context.Context, id int64) (*modelRepo.User, error) {
	query, args, err := r.sq.Select(
		usersID, usersName, usersEmail,
		usersRole, usersCreatedAt, usersUpdatedAt).
		From(users).
		Where(sq.Eq{usersID: id}).
		Limit(1).
		ToSql()

	if err != nil {
		log.Printf("failed to build query get user: %s", err.Error())
		return nil, err
	}

	q := db.Query{
		Name:     "auth_repository.Get",
		QueryRaw: query,
	}
	var user modelRepo.User
	err = r.db.DB().ScanOneContext(ctx, &user, q, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		log.Printf("failed to get user: %s", err.Error())
		return nil, err
	}

	return &user, nil
}
