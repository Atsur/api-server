package users

import (
	"time"

	"github.com/atsur/api-server/internal/pkg/models"
)

type Profile struct {
	models.Model
	FirstName string `gorm:"column:firstName;not null;" json:"firstName" form:"firstName"`
	LastName  string `gorm:"column:lastName;not null;" json:"lastName" form:"lastName"`
	Phone     string `gorm:"column:phone;not null;" json:"phone" form:"phone"`
	Bio       string `gorm:"column:bio;" json:"bio" form:"bio"`
	Avatar    string `gorm:"column:avatar;" json:"avatar" form:"avatar"`
	Address   string `gorm:"column:address;" json:"address" form:"address"`
	Type      string `gorm:"column:type;" json:"type" form:"type"`
	UserID    uint64 `gorm:"column:user_id;unique_index:user_id;" json:"user_id" form:"user_id"`
	User      User   `json:"user"`
}

func (m *Profile) BeforeCreate() error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

func (m *Profile) BeforeUpdate() error {
	m.UpdatedAt = time.Now()
	return nil
}
