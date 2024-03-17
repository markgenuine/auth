package auth

import (
	"context"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	desc "github.com/markgenuine/auth/pkg/auth_v1"
)

// Delete user from auth-service
func (s *Implementation) Delete(ctx context.Context, request *desc.DeleteRequest) (*empty.Empty, error) {
	err := s.authService.Delete(ctx, request.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("deleted user with id: %d", request.GetId())

	return &empty.Empty{}, nil
}
