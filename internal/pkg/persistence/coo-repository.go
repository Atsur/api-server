package persistence

import (
	"strconv"

	"github.com/atsur/api-server/internal/pkg/db"
	models "github.com/atsur/api-server/internal/pkg/models"
)

type CertificateOfOwnershipRepository struct{}

var cooRepository *CertificateOfOwnershipRepository

func GetCertificateOfOwnershipRepository() *CertificateOfOwnershipRepository {
	if cooRepository == nil {
		cooRepository = &CertificateOfOwnershipRepository{}
	}
	return cooRepository
}

func (r *CertificateOfOwnershipRepository) Get(id string) (*models.CertificateOfOwnership, error) {
	var CertificateOfOwnership models.CertificateOfOwnership
	where := models.CertificateOfOwnership{}
	where.ID, _ = strconv.ParseUint(id, 10, 64)
	_, err := First(&where, &CertificateOfOwnership, []string{"CertificateOfOwnership"})
	if err != nil {
		return nil, err
	}
	return &CertificateOfOwnership, err
}

func (r *CertificateOfOwnershipRepository) All() (*[]models.CertificateOfOwnership, error) {
	var CertificateOfOwnerships []models.CertificateOfOwnership
	err := Find(&models.CertificateOfOwnership{}, &CertificateOfOwnerships, []string{"CertificateOfOwnership"}, "id asc")
	return &CertificateOfOwnerships, err
}

func (r *CertificateOfOwnershipRepository) Query(q *models.CertificateOfOwnership) (*[]models.CertificateOfOwnership, error) {
	var CertificateOfOwnerships []models.CertificateOfOwnership
	err := Find(&q, &CertificateOfOwnerships, []string{"CertificateOfOwnership"}, "id asc")
	return &CertificateOfOwnerships, err
}

func (r *CertificateOfOwnershipRepository) Add(CertificateOfOwnership *models.CertificateOfOwnership) error {
	err := Create(&CertificateOfOwnership)
	err = Save(&CertificateOfOwnership)
	return err
}

func (r *CertificateOfOwnershipRepository) Update(CertificateOfOwnership *models.CertificateOfOwnership) error {
	return db.GetDB().Omit("CertificateOfOwnership").Save(&CertificateOfOwnership).Error
}

func (r *CertificateOfOwnershipRepository) Delete(CertificateOfOwnership *models.CertificateOfOwnership) error {
	return db.GetDB().Unscoped().Delete(&CertificateOfOwnership).Error
}
