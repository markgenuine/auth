package converter

import (
	desc "github.com/markgenuine/auth/pkg/auth_v1"
)

// ToUserRole ...
func ToUserRole(role string) desc.Role {
	return desc.Role(desc.Role_value[role])
}
