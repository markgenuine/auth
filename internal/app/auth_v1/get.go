package auth_v1

import (
	"context"
	"fmt"

	"github.com/brianvoe/gofakeit"
	desc "github.com/markgenuine/auth/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Get user of ID
func (s *User) Get(_ context.Context, request *desc.GetRequest) (*desc.GetResponse, error) {
	fmt.Printf("Get user with ID: %d", request.GetId())

	return &desc.GetResponse{
		Id:        request.GetId(),
		Name:      gofakeit.Name(),
		Email:     gofakeit.Email(),
		Role:      0,
		CreatedAt: timestamppb.New(gofakeit.Date()),
		UpdatedAt: timestamppb.New(gofakeit.Date()),
	}, nil
}
