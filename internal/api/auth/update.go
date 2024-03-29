package auth

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/markgenuine/auth/internal/converter"
	desc "github.com/markgenuine/auth/pkg/auth_v1"
)

// Update user of ID
func (s *Implementation) Update(ctx context.Context, request *desc.UpdateRequest) (*empty.Empty, error) {
	err := s.authService.Update(ctx, converter.UpdateToServiceFromUser(request))
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
