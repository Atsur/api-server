package models

import (
	"time"
)

type Exhibition struct {
	Model
	Name                string    `gorm:"column:name;not null;" json:"name" form:"name"`
	Type                string    `gorm:"column:type;not null;" json:"type" form:"type"`
	Label               string    `gorm:"column:_label;not null;" json:"_label" form:"_label"`
	ArtifactTagNo       string    `gorm:"column:artifact_tag_no;not null;" json:"artifact_tag_no" form:"artifact_tag_no"`
	ExhibitionProfileID uint64    `gorm:"column:exhibition_profile_id;" json:"exhibition_profile_id" form:"exhibition_profile_id"`
	AmountSold          float64   `gorm:"column:amount_sold;" json:"amount_sold" form:"amount_sold"`
	DateStarted         time.Time `gorm:"column:date_started;type:datetime;not null;" json:"date_started"`
	DateEnded           time.Time `gorm:"column:date_ended;type:datetime;not null;" json:"date_ended"`
	Record              Record    `gorm:"references:Artifact" json:"artifact"`
	ExhibitionProfile   Profile   `json:"auction_profile"`
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
