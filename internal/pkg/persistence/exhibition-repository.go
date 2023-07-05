package persistence

import (
	"strconv"

	"github.com/atsur/api-server/internal/pkg/db"
	models "github.com/atsur/api-server/internal/pkg/models/artifacts"
)

type ExhibitionRepository struct{}

var exhibitionRepository *ExhibitionRepository

func GetExhibitionRepository() *ExhibitionRepository {
	if exhibitionRepository == nil {
		exhibitionRepository = &ExhibitionRepository{}
	}
	return exhibitionRepository
}

func (r *ExhibitionRepository) Get(id string) (*models.Exhibition, error) {
	var Exhibition models.Exhibition
	where := models.Exhibition{}
	where.ID, _ = strconv.ParseUint(id, 10, 64)
	_, err := First(&where, &Exhibition, []string{"Exhibition"})
	if err != nil {
		return nil, err
	}
	return &Exhibition, err
}

func (r *ExhibitionRepository) All() (*[]models.Exhibition, error) {
	var Exhibitions []models.Exhibition
	err := Find(&models.Exhibition{}, &Exhibitions, []string{"Exhibition"}, "id asc")
	return &Exhibitions, err
}

func (r *ExhibitionRepository) Query(q *models.Exhibition) (*[]models.Exhibition, error) {
	var Exhibitions []models.Exhibition
	err := Find(&q, &Exhibitions, []string{"Exhibition"}, "id asc")
	return &Exhibitions, err
}

func (r *ExhibitionRepository) Add(Exhibition *models.Exhibition) error {
	err := Create(&Exhibition)
	err = Save(&Exhibition)
	return err
}

func (r *ExhibitionRepository) Update(Exhibition *models.Exhibition) error {
	return db.GetDB().Omit("Exhibition").Save(&Exhibition).Error
}

func (r *ExhibitionRepository) Delete(Exhibition *models.Exhibition) error {
	return db.GetDB().Unscoped().Delete(&Exhibition).Error
}
