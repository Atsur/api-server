package artifacts

import (
	"time"

	"github.com/atsur/api-server/internal/pkg/models"
	"github.com/atsur/api-server/internal/pkg/models/users"
)

type Artifact struct {
	models.Model
	Title           string        `gorm:"column:title;not null;" json:"title" form:"title"`
	Description     string        `gorm:"column:description;not null;" json:"description" form:"description"`
	Type            string        `gorm:"column:type;not null;" json:"type" form:"type"`
	Label           string        `gorm:"column:_label;not null;" json:"_label" form:"_label"`
	ArtistProfileID uint64        `gorm:"column:artist_profile_id;unique_index:artist_profile_id;not null;" json:"artist_profile_id" form:"artist_profile_id"`
	ArtistProfile   users.Profile `json:"artistProfile"`
	Provenance      []Provenance  `json:"provenance"`
	Exhibitions     []Exhibition  `json:"exhibitions"`
	Auctions        []Auction     `json:"auctions"`
}

func (m *Artifact) BeforeCreate() error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

func (m *Artifact) BeforeUpdate() error {
	m.UpdatedAt = time.Now()
	return nil
}
