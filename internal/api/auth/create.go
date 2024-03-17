package auth

import (
	"context"
	"log"

	"github.com/markgenuine/auth/internal/converter"
	desc "github.com/markgenuine/auth/pkg/auth_v1"
)

// Create user in auth-service
func (s *Implementation) Create(ctx context.Context, request *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := s.authService.Create(ctx, converter.CreateUserToServiceFromUser(request))
	if err != nil {
		return nil, err
	}

	log.Printf("inserted user with id: %d", id)

	return converter.CreateUserToUserFromService(id), nil
}
