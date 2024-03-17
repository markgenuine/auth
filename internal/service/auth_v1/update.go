package auth_v1

import (
	"context"
	"errors"
	"log"

	"github.com/markgenuine/auth/internal/model"
)

func (s *service) Update(ctx context.Context, user *model.UserUpdate) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		errTx := s.userRepository.Update(ctx, user)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		log.Print(err)
		return errors.New("failed updated user by id")
	}
	return nil
}
