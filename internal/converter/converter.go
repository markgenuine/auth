package converter

import (
	"github.com/markgenuine/auth/internal/model"
	modelRepo "github.com/markgenuine/auth/internal/repository/auth/model"
	desc "github.com/markgenuine/auth/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// CreateUserToServiceFromUser ...
func CreateUserToServiceFromUser(request *desc.CreateRequest) *model.UserCreate {
	return &model.UserCreate{
		Name:            request.GetName(),
		Email:           request.GetEmail(),
		Password:        request.GetPassword(),
		PasswordConfirm: request.GetPasswordConfirm(),
		Role:            request.GetRole().String(),
	}
}

// CreateUserToUserFromService ...
func CreateUserToUserFromService(id int64) *desc.CreateResponse {
	return &desc.CreateResponse{Id: id}
}

// GetToServiceFromRepo ...
func GetToServiceFromRepo(user *modelRepo.User) *model.User {
	return &model.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// GetToUserFromService ...
func GetToUserFromService(user *model.User) *desc.GetResponse {
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

// UpdateToServiceFromUser ...
func UpdateToServiceFromUser(request *desc.UpdateRequest) *model.UserUpdate {
	name := request.GetName().String()
	email := request.GetEmail().String()
	role := request.GetRole().String()
	return &model.UserUpdate{
		ID:    request.GetId(),
		Name:  &name,
		Email: &email,
		Role:  &role,
	}
}
