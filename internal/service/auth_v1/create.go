package auth_v1

import (
	"context"
	"errors"
	"log"

	"github.com/markgenuine/auth/internal/model"
)

func (s *service) Create(ctx context.Context, user *model.UserCreate) (int64, error) {
	var id int64
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.userRepository.Create(ctx, user)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		log.Print(err)
		return 0, errors.New("failed to create user")
	}

	return id, nil
}
