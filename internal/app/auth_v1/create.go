package auth_v1

import (
	"context"
	"fmt"

	desc "github.com/markgenuine/auth/pkg/auth_v1"
)

// Create user in auth-service
func (s *Auth) Create(ctx context.Context, request *desc.CreateRequest) (*desc.CreateResponse, error) {
	_ = ctx
	fmt.Printf("Create User: %s", request.GetName())
	fmt.Printf("Create User email: %s", request.GetEmail())
	fmt.Printf("Create User role: %s", request.GetRole())

	return &desc.CreateResponse{
		Id: 0,
	}, nil
}
