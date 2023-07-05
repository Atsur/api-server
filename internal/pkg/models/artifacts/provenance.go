package artifacts

import (
	"time"

	"github.com/atsur/api-server/internal/pkg/models"
	// "github.com/atsur/api-server/internal/pkg/models/artifacts"
	"github.com/atsur/api-server/internal/pkg/models/users"
)

type Provenance struct {
	models.Model
	Description    string     `gorm:"column:description;not null;" json:"description" form:"description"`
	Type           string     `gorm:"column:type;not null;" json:"type" form:"type"`
	Label          string     `gorm:"column:_label;not null;" json:"_label" form:"_label"`
	ArtifactID     uint64     `gorm:"column:artifact_id;unique_index:artifact_id;not null;" json:"artifact_id" form:"artifact_id"`
	OwnerProfileID uint64     `gorm:"column:owner_id;unique_index:owner_id;not null;" json:"owner_id" form:"owner_id"`
	Artifact       Artifact   `json:"artifact"`
	OwnerProfile   users.User `json:"ownerProfile"`
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
