package persistence

import (
	"strconv"

	"github.com/atsur/api-server/internal/pkg/db"
	models "github.com/atsur/api-server/internal/pkg/models/artifacts"
)

type ProvenanceRepository struct{}

var provenanceRepository *ProvenanceRepository

func GetProvenanceRepository() *ProvenanceRepository {
	if provenanceRepository == nil {
		provenanceRepository = &ProvenanceRepository{}
	}
	return provenanceRepository
}

func (r *ProvenanceRepository) Get(id string) (*models.Provenance, error) {
	var Provenance models.Provenance
	where := models.Provenance{}
	where.ID, _ = strconv.ParseUint(id, 10, 64)
	_, err := First(&where, &Provenance, []string{"Provenance"})
	if err != nil {
		return nil, err
	}
	return &Provenance, err
}

func (r *ProvenanceRepository) All() (*[]models.Provenance, error) {
	var Provenances []models.Provenance
	err := Find(&models.Provenance{}, &Provenances, []string{"Provenance"}, "id asc")
	return &Provenances, err
}

func (r *ProvenanceRepository) Query(q *models.Provenance) (*[]models.Provenance, error) {
	var Provenances []models.Provenance
	err := Find(&q, &Provenances, []string{"Provenance"}, "id asc")
	return &Provenances, err
}

func (r *ProvenanceRepository) Add(Provenance *models.Provenance) error {
	err := Create(&Provenance)
	err = Save(&Provenance)
	return err
}

func (r *ProvenanceRepository) Update(Provenance *models.Provenance) error {
	return db.GetDB().Omit("Provenance").Save(&Provenance).Error
}

func (r *ProvenanceRepository) Delete(Provenance *models.Provenance) error {
	return db.GetDB().Unscoped().Delete(&Provenance).Error
}
