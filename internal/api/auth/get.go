package auth

import (
	"context"
	"log"

	"github.com/markgenuine/auth/internal/converter"
	desc "github.com/markgenuine/auth/pkg/auth_v1"
)

// Get user of ID
func (s *Implementation) Get(ctx context.Context, request *desc.GetRequest) (*desc.GetResponse, error) {
	user, err := s.authService.Get(ctx, request.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("getted user with id: %d", request.GetId())

	return converter.GetToUserFromService(user), nil //add converter
}
