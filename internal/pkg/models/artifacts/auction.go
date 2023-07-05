package artifacts

import (
	"time"

	"github.com/atsur/api-server/internal/pkg/models"
	"github.com/atsur/api-server/internal/pkg/models/users"
)

type Auction struct {
	models.Model
	Name   string     `gorm:"column:name;not null;" json:"name" form:"name"`
	Type   string     `gorm:"column:type;not null;" json:"type" form:"type"`
	Label  string     `gorm:"column:_label;not null;" json:"_label" form:"_label"`
	UserID uint64     `gorm:"column:user_id;unique_index:user_id;not null;" json:"user_id" form:"user_id"`
	User   users.User `json:"user"`
}

func (m *Auction) BeforeCreate() error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

func (m *Auction) BeforeUpdate() error {
	m.UpdatedAt = time.Now()
	return nil
}
