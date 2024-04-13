package authservice

import (
	"context"
	"errors"
	"log"
)

func (s *service) Delete(ctx context.Context, id int64) error {
	err := s.userRepository.Delete(ctx, id)
	if err != nil {
		log.Print(err)
		return errors.New("failed to delete user")
	}

	return nil
}
