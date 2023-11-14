package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	FirstName   string    `gorm:"not null"`
	LastName    string    `gorm:"not null"`
	Username    string    `gorm:"uniqueIndex;not null"`
	Email       string    `gorm:"uniqueIndex;not null"`
	Password    string    `gorm:"not null"`
	LastLogin   time.Time
	CreatedAt   time.Time
	UpdateAt    time.Time
	Roles       []Role       `gorm:"many2many:user_roles;"`
	Permissions []Permission `gorm:"many2many:user_permissions;"`
}

type UserPermission struct {
	UserID       uuid.UUID
	PermissionID uuid.UUID
	CreatedAt    time.Time
	UpdateAt     time.Time
}

// Set the time the user role was created.
type UserRoles struct {
	UserID    uuid.UUID
	RoleID    uuid.UUID
	CreatedAt time.Time
	UpdateAt  time.Time
}
