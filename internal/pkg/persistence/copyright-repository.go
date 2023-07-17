package persistence

import (
	"strconv"

	"github.com/atsur/api-server/internal/pkg/db"
	models "github.com/atsur/api-server/internal/pkg/models"
)

type CopyrightRepository struct{}

var copyrightRepository *CopyrightRepository

func GetCopyrightRepository() *CopyrightRepository {
	if copyrightRepository == nil {
		copyrightRepository = &CopyrightRepository{}
	}
	return copyrightRepository
}

func (r *CopyrightRepository) Get(id string) (*models.Copyright, error) {
	var Copyright models.Copyright
	where := models.Copyright{}
	where.ID, _ = strconv.ParseUint(id, 10, 64)
	_, err := First(&where, &Copyright, []string{"Copyright"})
	if err != nil {
		return nil, err
	}
	return &Copyright, err
}

func (r *CopyrightRepository) All() (*[]models.Copyright, error) {
	var Copyrights []models.Copyright
	err := Find(&models.Copyright{}, &Copyrights, []string{"Copyright"}, "id asc")
	return &Copyrights, err
}

func (r *CopyrightRepository) Query(q *models.Copyright) (*[]models.Copyright, error) {
	var Copyrights []models.Copyright
	err := Find(&q, &Copyrights, []string{"Copyright"}, "id asc")
	return &Copyrights, err
}

func (r *CopyrightRepository) Add(Copyright *models.Copyright) error {
	err := Create(&Copyright)
	err = Save(&Copyright)
	return err
}

func (r *CopyrightRepository) Update(Copyright *models.Copyright) error {
	return db.GetDB().Omit("Copyright").Save(&Copyright).Error
}

func (r *CopyrightRepository) Delete(Copyright *models.Copyright) error {
	return db.GetDB().Unscoped().Delete(&Copyright).Error
}
