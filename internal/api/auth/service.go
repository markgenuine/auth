package auth

import (
	"github.com/markgenuine/auth/internal/service"
	desc "github.com/markgenuine/auth/pkg/auth_v1"
)

// Implementation - type for proto implementation
type Implementation struct {
	desc.UnimplementedUserV1Server
	authService service.AuthService
}

// NewImplementation create proto interface implementation
func NewImplementation(authService service.AuthService) *Implementation {
	return &Implementation{
		authService: authService,
	}
}
