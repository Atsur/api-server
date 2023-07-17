package persistence

import (
	"strconv"

	"github.com/atsur/api-server/internal/pkg/db"
	models "github.com/atsur/api-server/internal/pkg/models"
)

type EntryRepository struct{}

var entryRepository *EntryRepository

func GetEntryRepository() *EntryRepository {
	if entryRepository == nil {
		entryRepository = &EntryRepository{}
	}
	return entryRepository
}

func (r *EntryRepository) Get(id string) (*models.Entry, error) {
	var Entry models.Entry
	where := models.Entry{}
	where.ID, _ = strconv.ParseUint(id, 10, 64)
	_, err := First(&where, &Entry, []string{"Entry"})
	if err != nil {
		return nil, err
	}
	return &Entry, err
}

func (r *EntryRepository) GetProvenanceByEntry(EntryId string) (*[]models.Provenance, error) {
	var Provenance []models.Provenance
	// where := models.Provenance{}
	// EntryId, _ := strconv.ParseUint(id, 10, 64)
	err := Find("Entry_id = ?", &Provenance, []string{EntryId}, "id asc")
	if err != nil {
		return nil, err
	}
	return &Provenance, err
}

func (r *EntryRepository) GetExhibitionByEntry(EntryId string) (*[]models.Exhibition, error) {
	var Exhibition []models.Exhibition
	// where := models.Provenance{}
	// EntryId, _ := strconv.ParseUint(id, 10, 64)
	err := Find("Entry_id = ?", &Exhibition, []string{EntryId}, "id asc")
	if err != nil {
		return nil, err
	}
	return &Exhibition, err
}

func (r *EntryRepository) GetAuctionEntry(EntryId string) (*[]models.Auction, error) {
	var Auction []models.Auction
	// where := models.Provenance{}
	// EntryId, _ := strconv.ParseUint(id, 10, 64)
	err := Find("Entry_id = ?", &Auction, []string{EntryId}, "id asc")
	if err != nil {
		return nil, err
	}
	return &Auction, err
}

func (r *EntryRepository) All() (*[]models.Entry, error) {
	var Entrys []models.Entry
	err := Find(&models.Entry{}, &Entrys, []string{"Entry"}, "id asc")
	return &Entrys, err
}

func (r *EntryRepository) Query(q *models.Entry) (*[]models.Entry, error) {
	var Entrys []models.Entry
	err := Find(&q, &Entrys, []string{"Entry"}, "id asc")
	return &Entrys, err
}

func (r *EntryRepository) Add(Entry *models.Entry) error {
	err := Create(&Entry)
	err = Save(&Entry)
	return err
}

func (r *EntryRepository) Update(Entry *models.Entry) error {
	return db.GetDB().Omit("Entry").Save(&Entry).Error
}

func (r *EntryRepository) Delete(Entry *models.Entry) error {
	return db.GetDB().Unscoped().Delete(&Entry).Error
}
