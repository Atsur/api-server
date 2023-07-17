package persistence

import (
	"strconv"

	"github.com/atsur/api-server/internal/pkg/db"
	models "github.com/atsur/api-server/internal/pkg/models"
)

type RecordRepository struct{}

var recordRepository *RecordRepository

func GetRecordRepository() *RecordRepository {
	if recordRepository == nil {
		recordRepository = &RecordRepository{}
	}
	return recordRepository
}

func (r *RecordRepository) Get(id string) (*models.Record, error) {
	var Record models.Record
	where := models.Record{}
	where.ID, _ = strconv.ParseUint(id, 10, 64)
	_, err := First(&where, &Record, []string{"Record"})
	if err != nil {
		return nil, err
	}
	return &Record, err
}

func (r *RecordRepository) GetProvenanceByRecord(RecordId string) (*[]models.Provenance, error) {
	var Provenance []models.Provenance
	// where := models.Provenance{}
	// RecordId, _ := strconv.ParseUint(id, 10, 64)
	err := Find("Record_id = ?", &Provenance, []string{RecordId}, "id asc")
	if err != nil {
		return nil, err
	}
	return &Provenance, err
}

func (r *RecordRepository) GetExhibitionByRecord(RecordId string) (*[]models.Exhibition, error) {
	var Exhibition []models.Exhibition
	// where := models.Provenance{}
	// RecordId, _ := strconv.ParseUint(id, 10, 64)
	err := Find("Record_id = ?", &Exhibition, []string{RecordId}, "id asc")
	if err != nil {
		return nil, err
	}
	return &Exhibition, err
}

func (r *RecordRepository) GetAuctionRecord(RecordId string) (*[]models.Auction, error) {
	var Auction []models.Auction
	// where := models.Provenance{}
	// RecordId, _ := strconv.ParseUint(id, 10, 64)
	err := Find("Record_id = ?", &Auction, []string{RecordId}, "id asc")
	if err != nil {
		return nil, err
	}
	return &Auction, err
}

func (r *RecordRepository) All() (*[]models.Record, error) {
	var Records []models.Record
	err := Find(&models.Record{}, &Records, []string{"Record"}, "id asc")
	return &Records, err
}

func (r *RecordRepository) Query(q *models.Record) (*[]models.Record, error) {
	var Records []models.Record
	err := Find(&q, &Records, []string{"Record"}, "id asc")
	return &Records, err
}

func (r *RecordRepository) Add(Record *models.Record) error {
	err := Create(&Record)
	err = Save(&Record)
	return err
}

func (r *RecordRepository) Update(Record *models.Record) error {
	return db.GetDB().Omit("Record").Save(&Record).Error
}

func (r *RecordRepository) Delete(Record *models.Record) error {
	return db.GetDB().Unscoped().Delete(&Record).Error
}
