package controllers

import (
	"errors"
	"log"
	"net/http"

	models "github.com/atsur/api-server/internal/pkg/models"
	"github.com/atsur/api-server/internal/pkg/persistence"
	"github.com/atsur/api-server/pkg/http_err"
	"github.com/gin-gonic/gin"
)

type EntryInput struct {
	Title          string                `json:"title,omitempty" binding:"required"`
	Description    string                `json:"description,omitempty"`
	Edition        string                `json:"edition,omitempty"`
	Series         string                `json:"series,omitempty"`
	YearCreated    uint16                `json:"year_created,omitempty" binding:"required"`
	SellingPrice   float64               `json:"selling_price,omitempty"`
	Metadata       string                `json:"metadata,omitempty"`
	Frame          string                `json:"frame,omitempty"`
	Dimensions     []uint16              `json:"dimensions,omitempty" binding:"required"`
	ArtistProfiles []uint64              `json:"artist_profiles,omitempty"`
	Classification models.Classification `json:"classification,omitempty" binding:"required"`
	EntryType      models.EntryType      `json:"entry_type,omitempty" binding:"required"`
	RecordType     models.RecordType     `json:"record_type,omitempty" binding:"required"`
	Materials      models.Materials      `json:"materials,omitempty" binding:"required"`
	Medium         models.Medium         `json:"medium,omitempty" binding:"required"`
}

func GetEntryById(c *gin.Context) {
	s := persistence.GetEntryRepository()
	id := c.Param("id")
	if Entry, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("Entry not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, Entry)
	}
}

func GetEntries(c *gin.Context) {
	s := persistence.GetEntryRepository()
	var q models.Entry
	_ = c.Bind(&q)
	if Entrys, err := s.Query(&q); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("Entrys not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, Entrys)
	}
}

func CreateEntry(c *gin.Context) {
	s := persistence.GetEntryRepository()
	var EntryInput EntryInput
	_ = c.BindJSON(&EntryInput)
	Entry := models.Entry{
		Title:            EntryInput.Title,
		Description:      EntryInput.Description,
		Edition:          EntryInput.Edition,
		Series:           EntryInput.Series,
		YearCreated:      EntryInput.YearCreated,
		SellingPrice:     EntryInput.SellingPrice,
		Metadata:         EntryInput.Metadata,
		Frame:            EntryInput.Frame,
		Dimensions:       EntryInput.Dimensions,
		ArtistProfileIDs: EntryInput.ArtistProfiles,
		Classification:   EntryInput.Classification,
		EntryType:        EntryInput.EntryType,
		RecordType:       EntryInput.RecordType,
		Materials:        EntryInput.Materials,
		Medium:           EntryInput.Medium,
	}
	if err := s.Add(&Entry); err != nil {
		http_err.NewError(c, http.StatusBadRequest, err)
		log.Println(err)
	} else {
		c.JSON(http.StatusCreated, Entry)
	}
}

func UpdateEntry(c *gin.Context) {
	s := persistence.GetEntryRepository()
	id := c.Params.ByName("id")
	var EntryInput EntryInput
	_ = c.BindJSON(&EntryInput)
	if Entry, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("Entry not found"))
		log.Println(err)
	} else {
		Entry.Title = EntryInput.Title
		Entry.Description = EntryInput.Description
		Entry.Edition = EntryInput.Edition
		Entry.Series = EntryInput.Series
		Entry.YearCreated = EntryInput.YearCreated
		Entry.SellingPrice = EntryInput.SellingPrice
		Entry.Metadata = EntryInput.Metadata
		Entry.Frame = EntryInput.Frame
		Entry.Dimensions = EntryInput.Dimensions
		Entry.ArtistProfileIDs = EntryInput.ArtistProfiles
		Entry.Classification = EntryInput.Classification
		Entry.EntryType = EntryInput.EntryType
		Entry.RecordType = EntryInput.RecordType
		Entry.Materials = EntryInput.Materials
		Entry.Medium = EntryInput.Medium
		if err := s.Update(Entry); err != nil {
			http_err.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, Entry)
		}
	}
}

func DeleteEntry(c *gin.Context) {
	s := persistence.GetEntryRepository()
	id := c.Params.ByName("id")
	var EntryInput EntryInput
	_ = c.BindJSON(&EntryInput)
	if Entry, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("Entry not found"))
		log.Println(err)
	} else {
		if err := s.Delete(Entry); err != nil {
			http_err.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusNoContent, "")
		}
	}
}
