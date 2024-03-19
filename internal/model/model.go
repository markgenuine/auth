package model

import (
	"database/sql"
	"time"
)

// UserCreate ...
type UserCreate struct {
	Name            string
	Email           string
	Password        string
	PasswordConfirm string
	Role            string
}

// UserUpdate ...
type UserUpdate struct {
	ID    int64
	Name  *string
	Email *string
	Role  *string
}

// User ...
type User struct {
	ID        int64
	Name      string
	Email     string
	Role      string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
