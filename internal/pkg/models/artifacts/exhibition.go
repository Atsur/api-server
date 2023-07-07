package artifacts

import (
	"time"

	"github.com/atsur/api-server/internal/pkg/models"
	"github.com/atsur/api-server/internal/pkg/models/users"
)

type Exhibition struct {
	models.Model
	Name                string        `gorm:"column:name;not null;" json:"name" form:"name"`
	Type                string        `gorm:"column:type;not null;" json:"type" form:"type"`
	Label               string        `gorm:"column:_label;not null;" json:"_label" form:"_label"`
	ArtifactID          uint64        `gorm:"column:artifact_id;not null;" json:"artifact_id" form:"artifact_id"`
	ExhibitionProfileID uint64        `gorm:"column:exhibition_profile_id;" json:"exhibition_profile_id" form:"exhibition_profile_id"`
	AmountSold          float64       `gorm:"column:amount_sold;" json:"amount_sold" form:"amount_sold"`
	DateStarted         time.Time     `gorm:"column:date_started;type:datetime;not null;" json:"date_started"`
	DateEnded           time.Time     `gorm:"column:date_ended;type:datetime;not null;" json:"date_ended"`
	Artifact            Artifact      `gorm:"references:Artifact" json:"artifact"`
	ExhibitionProfile   users.Profile `json:"auction_profile"`
}

func (m *Exhibition) BeforeCreate() error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

func (m *Exhibition) BeforeUpdate() error {
	m.UpdatedAt = time.Now()
	return nil
}
