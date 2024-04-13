package authservice

import (
	"context"
	"errors"
	"log"

	"github.com/markgenuine/auth/internal/converter"
	"github.com/markgenuine/auth/internal/model"
)

func (s *service) Get(ctx context.Context, id int64) (*model.User, error) {
	user, err := s.userRepository.Get(ctx, id)
	if err != nil {
		log.Print(err)
		return nil, errors.New("failed get user by id")
	}

	return converter.GetToServiceFromRepo(user), nil
}
