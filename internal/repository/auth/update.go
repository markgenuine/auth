package auth

import (
	"context"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/markgenuine/auth/internal/model"
	"github.com/markgenuine/platform_common/pkg/db"
)

func (r *repo) Update(ctx context.Context, user *model.UserUpdate) error {
	builder := r.sq.Update(users)

	if user.Name != nil {
		builder = builder.Set(usersName, user.Name)
	}

	if user.Email != nil {
		builder = builder.Set(usersEmail, user.Email)
	}

	if user.Role != nil && *user.Role != "UNKNOWN" {
		builder = builder.Set(usersRole, user.Role)
	}

	builder = builder.Set(usersUpdatedAt, time.Now())
	builder = builder.Where(sq.Eq{usersID: user.ID})
	query, args, err := builder.ToSql()
	if err != nil {
		log.Printf("failed to build query update user: %s", err.Error())
		return err
	}

	q := db.Query{
		Name:     "auth_repository.Update",
		QueryRaw: query,
	}
	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		log.Printf("failed to update user: %s", err.Error())
		return err
	}

	return nil
}
