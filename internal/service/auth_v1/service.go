package authService

import (
	"github.com/markgenuine/auth/internal/repository"
	def "github.com/markgenuine/auth/internal/service"
)

var _ def.AuthService = (*service)(nil)

type service struct {
	userRepository repository.UserRepository
}

// NewService ...
func NewService(userRepository repository.UserRepository) *service {
	return &service{
		userRepository: userRepository,
	}
}

func NewMockService(deps ...interface{}) *service {
	srv := service{}

	for _, v := range deps {
		switch s := v.(type) {
		case repository.UserRepository:
			srv.userRepository = s
		}
	}

	return &srv
}
