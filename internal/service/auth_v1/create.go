package authService

import (
	"context"
	"errors"
	"log"

	"github.com/markgenuine/auth/internal/model"
)

func (s *service) Create(ctx context.Context, user *model.UserCreate) (int64, error) {
	id, err := s.userRepository.Create(ctx, user)
	if err != nil {
		log.Print(err)
		return 0, errors.New("failed to create user")
	}

	return id, nil
}
