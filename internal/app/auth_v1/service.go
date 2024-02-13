package auth_v1

import desc "github.com/markgenuine/auth/pkg/auth_v1"

// Auth - type for proto implementation
type Auth struct {
	desc.UnimplementedAuthV1Server
}

// NewAuth create proto interface implementation
func NewAuth() *Auth {
	return &Auth{}
}
