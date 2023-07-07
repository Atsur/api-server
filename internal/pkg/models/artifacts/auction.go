package artifacts

import (
	"time"

	"github.com/atsur/api-server/internal/pkg/models"
	"github.com/atsur/api-server/internal/pkg/models/users"
)

type Auction struct {
	models.Model
	Title            string        `gorm:"column:title;not null;" json:"title" form:"title"`
	Type             string        `gorm:"column:type;not null;" json:"type" form:"type"`
	Description      string        `gorm:"column:description;not null;" json:"description" form:"description"`
	ArtifactID       uint64        `gorm:"column:artifact_id;not null;" json:"artifact_id" form:"artifact_id"`
	AuctionProfileID uint64        `gorm:"column:auction_profile_id;" json:"auction_profile_id" form:"auction_profile_id"`
	AmountSold       float64       `gorm:"column:amount_sold;" json:"amount_sold" form:"amount_sold"`
	DateStarted      time.Time     `gorm:"column:date_started;type:datetime;not null;" json:"date_started"`
	DateEnded        time.Time     `gorm:"column:date_ended;type:datetime;not null;" json:"date_ended"`
	Artifact         Artifact      `json:"artifact"`
	AuctionProfile   users.Profile `json:"auction_profile"`
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
