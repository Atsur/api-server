package persistence

import (
	"strconv"

	"github.com/atsur/api-server/internal/pkg/db"
	models "github.com/atsur/api-server/internal/pkg/models"
)

type CertificateOfAuthenticityRepository struct{}

var coaRepository *CertificateOfAuthenticityRepository

func GetCertificateOfAuthenticityRepository() *CertificateOfAuthenticityRepository {
	if coaRepository == nil {
		coaRepository = &CertificateOfAuthenticityRepository{}
	}
	return coaRepository
}

func (r *CertificateOfAuthenticityRepository) Get(id string) (*models.CertificateOfAuthenticity, error) {
	var CertificateOfAuthenticity models.CertificateOfAuthenticity
	where := models.CertificateOfAuthenticity{}
	where.ID, _ = strconv.ParseUint(id, 10, 64)
	_, err := First(&where, &CertificateOfAuthenticity, []string{"CertificateOfAuthenticity"})
	if err != nil {
		return nil, err
	}
	return &CertificateOfAuthenticity, err
}

func (r *CertificateOfAuthenticityRepository) All() (*[]models.CertificateOfAuthenticity, error) {
	var CertificateOfAuthenticitys []models.CertificateOfAuthenticity
	err := Find(&models.CertificateOfAuthenticity{}, &CertificateOfAuthenticitys, []string{"CertificateOfAuthenticity"}, "id asc")
	return &CertificateOfAuthenticitys, err
}

func (r *CertificateOfAuthenticityRepository) Query(q *models.CertificateOfAuthenticity) (*[]models.CertificateOfAuthenticity, error) {
	var CertificateOfAuthenticitys []models.CertificateOfAuthenticity
	err := Find(&q, &CertificateOfAuthenticitys, []string{"CertificateOfAuthenticity"}, "id asc")
	return &CertificateOfAuthenticitys, err
}

func (r *CertificateOfAuthenticityRepository) Add(CertificateOfAuthenticity *models.CertificateOfAuthenticity) error {
	err := Create(&CertificateOfAuthenticity)
	err = Save(&CertificateOfAuthenticity)
	return err
}

func (r *CertificateOfAuthenticityRepository) Update(CertificateOfAuthenticity *models.CertificateOfAuthenticity) error {
	return db.GetDB().Omit("CertificateOfAuthenticity").Save(&CertificateOfAuthenticity).Error
}

func (r *CertificateOfAuthenticityRepository) Delete(CertificateOfAuthenticity *models.CertificateOfAuthenticity) error {
	return db.GetDB().Unscoped().Delete(&CertificateOfAuthenticity).Error
}
