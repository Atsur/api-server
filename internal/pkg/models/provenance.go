package models

import (
	"time"
)

type Provenance struct {
	Model
	Description    string  `gorm:"column:description;not null;" json:"description" form:"description"`
	Type           string  `gorm:"column:type;not null;" json:"type" form:"type"`
	Label          string  `gorm:"column:_label;not null;" json:"_label" form:"_label"`
	ArtifactTagNo  string  `gorm:"column:artifact_tag_no;not null;" json:"artifact_tag_no" form:"artifact_tag_no"`
	OwnerProfileID uint64  `gorm:"column:owner_profile_id;not null;" json:"owner_profile_id" form:"owner_profile_id"`
	Record         Record  `gorm:"references:Artifact" json:"artifact"`
	OwnerProfile   Profile `gorm:"references:Profile" json:"ownerProfile"`
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
