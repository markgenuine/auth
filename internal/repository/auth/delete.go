package auth

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/markgenuine/auth/internal/client/db"
)

func (r *repo) Delete(ctx context.Context, id int64) error {
	query, args, err := r.sq.Delete(users).Where(sq.Eq{
		usersID: id,
	}).ToSql()

	if err != nil {
		log.Printf("failed to build query delete user: %s", err.Error())
		return err
	}

	q := db.Query{
		Name:     "auth_repository.Delete",
		QueryRaw: query,
	}
	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		log.Printf("failed to delete user: %s", err.Error())
		return err
	}

	return nil
}
