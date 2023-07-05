package controllers

import (
	"errors"
	"log"
	"net/http"

	models "github.com/atsur/api-server/internal/pkg/models/artifacts"
	"github.com/atsur/api-server/internal/pkg/persistence"
	// "github.com/atsur/api-server/pkg/crypto"
	"github.com/atsur/api-server/pkg/http_err"
	"github.com/gin-gonic/gin"
)

type AuctionInput struct {
	AuctionName string `json:"AuctionName" binding:"required"`
	Lastname    string `json:"lastname"`
	Firstname   string `json:"firstname"`
	Password    string `json:"password" binding:"required"`
	Role        string `json:"role"`
}

// GetAuctionById godoc
// @Summary Retrieves Auction based on given ID
// @Description get Auction by ID
// @Produce json
// @Param id path integer true "Auction ID"
// @Success 200 {object} Auctions.Auction
// @Router /api/Auctions/{id} [get]
// @Security Authorization Token
func GetAuctionById(c *gin.Context) {
	s := persistence.GetAuctionRepository()
	id := c.Param("id")
	if Auction, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("Auction not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, Auction)
	}
}

// GetAuctions godoc
// @Summary Retrieves Auctions based on query
// @Description Get Auctions
// @Produce json
// @Param Auctionname query string false "Auctionname"
// @Param firstname query string false "Firstname"
// @Param lastname query string false "Lastname"
// @Success 200 {array} []Auctions.Auction
// @Router /api/Auctions [get]
// @Security Authorization Token
func GetAuctions(c *gin.Context) {
	s := persistence.GetAuctionRepository()
	var q models.Auction
	_ = c.Bind(&q)
	if Auctions, err := s.Query(&q); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("Auctions not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, Auctions)
	}
}

func CreateAuction(c *gin.Context) {
	s := persistence.GetAuctionRepository()
	var AuctionInput AuctionInput
	_ = c.BindJSON(&AuctionInput)
	Auction := models.Auction{
		// Auctionname:  AuctionInput.Auctionname,
		// Firstname: AuctionInput.Firstname,
		// Lastname:  AuctionInput.Lastname,
		// Hash:      crypto.HashAndSalt([]byte(AuctionInput.Password)),
		// Role:      models.AuctionRole{RoleName: AuctionInput.Role},
	}
	if err := s.Add(&Auction); err != nil {
		http_err.NewError(c, http.StatusBadRequest, err)
		log.Println(err)
	} else {
		c.JSON(http.StatusCreated, Auction)
	}
}

func UpdateAuction(c *gin.Context) {
	s := persistence.GetAuctionRepository()
	id := c.Params.ByName("id")
	var AuctionInput AuctionInput
	_ = c.BindJSON(&AuctionInput)
	if Auction, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("Auction not found"))
		log.Println(err)
	} else {
		// Auction.Auctionname = AuctionInput.Auctionname
		// Auction.Lastname = AuctionInput.Lastname
		// Auction.Firstname = AuctionInput.Firstname
		// Auction.Hash = crypto.HashAndSalt([]byte(AuctionInput.Password))
		// Auction.Role = models.AuctionRole{RoleName: AuctionInput.Role}
		if err := s.Update(Auction); err != nil {
			http_err.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, Auction)
		}
	}
}

func DeleteAuction(c *gin.Context) {
	s := persistence.GetAuctionRepository()
	id := c.Params.ByName("id")
	var AuctionInput AuctionInput
	_ = c.BindJSON(&AuctionInput)
	if Auction, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("Auction not found"))
		log.Println(err)
	} else {
		if err := s.Delete(Auction); err != nil {
			http_err.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusNoContent, "")
		}
	}
}
