package auth_v1

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	desc "github.com/markgenuine/auth/pkg/auth_v1"
)

// Update user of ID
func (s *Auth) Update(ctx context.Context, request *desc.UpdateRequest) (*empty.Empty, error) {
	_ = ctx
	fmt.Printf("Update user with ID: %d", request.GetId())

	return &empty.Empty{}, nil
}
