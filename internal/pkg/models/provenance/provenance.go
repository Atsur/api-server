package provenance

import (
	"time"

	"github.com/atsur/api-server/internal/pkg/models"
	"github.com/atsur/api-server/internal/pkg/models/users"
)

type Provenance struct {
	models.Model
	Name   string     `gorm:"column:name;not null;" json:"name" form:"name"`
	Text   string     `gorm:"column:text;not null;" json:"text" form:"text"`
	UserID uint64     `gorm:"column:user_id;unique_index:user_id;not null;" json:"user_id" form:"user_id"`
	User   users.User `json:"user"`
	Auctions   []Auction `json:"auctions"`
	Exhibitions   []Exhibition `json:"exhibitions"`
	History   History `json:"history"`
}

func (m *Provenance) BeforeCreate() error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

func (m *Provenance) BeforeUpdate() error {
	m.UpdatedAt = time.Now()
	return nil
}
