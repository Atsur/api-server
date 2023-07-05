package users

import (
	"time"

	"github.com/atsur/api-server/internal/pkg/models"
)

type User struct {
	models.Model
	UUID         string   `gorm:"column:uuid;not null;unique_index:uuid" json:"uuid" form:"uuid"`
	Email        string   `gorm:"column:email;not null;" json:"email" form:"email"`
	PasswordHash string   `gorm:"column:hash;not null;" json:"hash"`
	Role         UserRole `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (m *User) BeforeCreate() error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

func (m *User) BeforeUpdate() error {
	m.UpdatedAt = time.Now()
	return nil
}
