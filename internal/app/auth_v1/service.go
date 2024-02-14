package auth_v1

import desc "github.com/markgenuine/auth/pkg/auth_v1"

// User - type for proto implementation
type User struct {
	desc.UnimplementedUserV1Server
}

// NewUserService create proto interface implementation
func NewUserService() *User {
	return &User{}
}
