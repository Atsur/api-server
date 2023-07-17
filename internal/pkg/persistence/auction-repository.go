package persistence

import (
	"strconv"

	"github.com/atsur/api-server/internal/pkg/db"
	models "github.com/atsur/api-server/internal/pkg/models"
)

type AuctionRepository struct{}

var auctionRepository *AuctionRepository

func GetAuctionRepository() *AuctionRepository {
	if auctionRepository == nil {
		auctionRepository = &AuctionRepository{}
	}
	return auctionRepository
}

func (r *AuctionRepository) Get(id string) (*models.Auction, error) {
	var Auction models.Auction
	where := models.Auction{}
	where.ID, _ = strconv.ParseUint(id, 10, 64)
	_, err := First(&where, &Auction, []string{"Auction"})
	if err != nil {
		return nil, err
	}
	return &Auction, err
}

func (r *AuctionRepository) All() (*[]models.Auction, error) {
	var Auctions []models.Auction
	err := Find(&models.Auction{}, &Auctions, []string{"Auction"}, "id asc")
	return &Auctions, err
}

func (r *AuctionRepository) Query(q *models.Auction) (*[]models.Auction, error) {
	var Auctions []models.Auction
	err := Find(&q, &Auctions, []string{"Auction"}, "id asc")
	return &Auctions, err
}

func (r *AuctionRepository) Add(Auction *models.Auction) error {
	err := Create(&Auction)
	err = Save(&Auction)
	return err
}

func (r *AuctionRepository) Update(Auction *models.Auction) error {
	return db.GetDB().Omit("Auction").Save(&Auction).Error
}

func (r *AuctionRepository) Delete(Auction *models.Auction) error {
	return db.GetDB().Unscoped().Delete(&Auction).Error
}
