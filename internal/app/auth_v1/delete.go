package auth_v1

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	desc "github.com/markgenuine/auth/pkg/auth_v1"
)

// Delete user from auth-service
func (s *Auth) Delete(ctx context.Context, request *desc.DeleteRequest) (*empty.Empty, error) {
	_ = ctx
	fmt.Printf("User delete with ID: %d", request.GetId())

	return &empty.Empty{}, nil
}
