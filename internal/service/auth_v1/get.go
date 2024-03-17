package auth_v1

import (
	"context"
	"errors"
	"log"

	"github.com/markgenuine/auth/internal/converter"
	"github.com/markgenuine/auth/internal/model"
	modelrepo "github.com/markgenuine/auth/internal/repository/auth/model"
)

func (s *service) Get(ctx context.Context, id int64) (*model.User, error) {
	var user *modelrepo.User
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		user, errTx = s.userRepository.Get(ctx, id)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		log.Print(err)
		return nil, errors.New("failed get user by id")
	}

	return converter.GetToServiceFromRepo(user), nil
}
