package converter

import (
	modelRepo "github.com/markgenuine/auth/internal/repository/auth/model"
	desc "github.com/markgenuine/auth/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ToUserFromRepo ...
func ToUserFromRepo(user *modelRepo.User) *desc.GetResponse {

	var updatedAt *timestamppb.Timestamp
	if user.UpdatedAt.Valid {
		updatedAt = timestamppb.New(user.UpdatedAt.Time)
	}

	return &desc.GetResponse{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      ToUserRole(user.Role),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

// ToUserRole ...
func ToUserRole(role string) desc.Role {
	return desc.Role(desc.Role_value[role])
}
