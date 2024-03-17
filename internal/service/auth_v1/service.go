package auth_v1

import (
	"github.com/markgenuine/auth/internal/client/db"
	"github.com/markgenuine/auth/internal/repository"
	def "github.com/markgenuine/auth/internal/service"
)

var _ def.AuthService = (*service)(nil)

type service struct {
	userRepository repository.UserRepository
	txManager      db.TxManager
}

// NewService ...
func NewService(userRepository repository.UserRepository, txManager db.TxManager) *service {
	return &service{
		userRepository: userRepository,
		txManager:      txManager,
	}
}
