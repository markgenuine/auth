package authService

import (
	"context"
	"errors"
	"log"

	"github.com/markgenuine/auth/internal/model"
)

func (s *service) Update(ctx context.Context, user *model.UserUpdate) error {
	err := s.userRepository.Update(ctx, user)
	if err != nil {
		log.Print(err)
		return errors.New("failed updated user by id")
	}
	return nil
}
