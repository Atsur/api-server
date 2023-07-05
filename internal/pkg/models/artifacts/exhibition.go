package artifacts

import (
	"time"

	"github.com/atsur/api-server/internal/pkg/models"
	// "github.com/atsur/api-server/internal/pkg/models/users"
)

type Exhibition struct {
	models.Model
	Name       string   `gorm:"column:name;not null;" json:"name" form:"name"`
	Type       string   `gorm:"column:type;not null;" json:"type" form:"type"`
	Label      string   `gorm:"column:_label;not null;" json:"_label" form:"_label"`
	ArtifactID uint64   `gorm:"column:artifact_id;not null;" json:"artifact_id" form:"artifact_id"`
	Artifact   Artifact `json:"artifact"`
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
