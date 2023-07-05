package persistence

import (
	"strconv"

	"github.com/atsur/api-server/internal/pkg/db"
	models "github.com/atsur/api-server/internal/pkg/models/users"
)

type ProfileRepository struct{}

var profileRepository *ProfileRepository

func GetProfileRepository() *ProfileRepository {
	if profileRepository == nil {
		profileRepository = &ProfileRepository{}
	}
	return profileRepository
}

func (r *ProfileRepository) Get(id string) (*models.Profile, error) {
	var Profile models.Profile
	where := models.Profile{}
	where.ID, _ = strconv.ParseUint(id, 10, 64)
	_, err := First(&where, &Profile, []string{"User"})
	if err != nil {
		return nil, err
	}
	return &Profile, err
}

func (r *ProfileRepository) GetByUserId(userId string) (*models.Profile, error) {
	var Profile models.Profile
	where := models.Profile{}
	where.UserID, _ = strconv.ParseUint(userId, 10, 64)
	_, err := First(&where, &Profile, []string{"User"})
	if err != nil {
		return nil, err
	}
	return &Profile, err
}

func (r *ProfileRepository) All() (*[]models.Profile, error) {
	var Profiles []models.Profile
	err := Find(&models.Profile{}, &Profiles, []string{"User"}, "id asc")
	return &Profiles, err
}

func (r *ProfileRepository) Query(q *models.Profile) (*[]models.Profile, error) {
	var Profiles []models.Profile
	err := Find(&q, &Profiles, []string{"User"}, "id asc")
	return &Profiles, err
}

func (r *ProfileRepository) Add(Profile *models.Profile) error {
	err := Create(&Profile)
	err = Save(&Profile)
	return err
}

func (r *ProfileRepository) Update(Profile *models.Profile) error {
	return db.GetDB().Omit("User").Save(&Profile).Error
}

func (r *ProfileRepository) Delete(Profile *models.Profile) error {
	return db.GetDB().Unscoped().Delete(&Profile).Error
}
