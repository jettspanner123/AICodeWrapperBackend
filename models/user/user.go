package user

import (
	"github.com/google/uuid"
)

type UserRole string

const (
	RoleUser UserRole = "user"
	RoleDev  UserRole = "dev"
)

type BaseUser struct {
	ID           uuid.UUIDs `gorm:"type:uuid;primaryKey"`
	Email        string
	Username     string
	PasswordHash string
	Role         UserRole
}
