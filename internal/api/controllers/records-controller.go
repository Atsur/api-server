package controllers

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	models "github.com/atsur/api-server/internal/pkg/models"
	"github.com/atsur/api-server/internal/pkg/persistence"
	// "github.com/atsur/api-server/pkg/crypto"
	"github.com/atsur/api-server/pkg/http_err"
	"github.com/gin-gonic/gin"
)

type RecordInput struct {
	Recordname string `json:"recordname" binding:"required"`
	Lastname   string `json:"lastname"`
	Firstname  string `json:"firstname"`
	Password   string `json:"password" binding:"required"`
	Role       string `json:"role"`
}

// GetRecordById godoc
// @Summary Retrieves Record based on given ID
// @Description get Record by ID
// @Produce json
// @Param id path integer true "Record ID"
// @Success 200 {object} Records.Record
// @Router /api/Records/{id} [get]
// @Security Authorization Token
func GetRecordById(c *gin.Context) {
	s := persistence.GetRecordRepository()
	id := c.Param("id")
	if Record, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("Record not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, Record)
	}
}

// GetRecordById godoc
// @Summary Retrieves Record based on given ID
// @Description get Record by ID
// @Produce json
// @Param id path integer true "Record ID"
// @Success 200 {object} Records.Record
// @Router /api/Records/{id} [get]
// @Security Authorization Token
func GetProvenanceByRecordId(c *gin.Context) {
	s := persistence.GetRecordRepository()
	Record_id := c.Param("id")
	if provenance, err := s.GetProvenanceByRecord(Record_id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New(fmt.Sprintf("Provenance not found for Record id %s", Record_id)))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, provenance)
	}
}

// GetRecordById godoc
// @Summary Retrieves Record based on given ID
// @Description get Record by ID
// @Produce json
// @Param id path integer true "Record ID"
// @Success 200 {object} Records.Record
// @Router /api/Records/{id} [get]
// @Security Authorization Token
func GetExhibitionByRecordId(c *gin.Context) {
	s := persistence.GetRecordRepository()
	Record_id := c.Param("id")
	if exhibitions, err := s.GetExhibitionByRecord(Record_id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New(fmt.Sprintf("Exhibitions not found for Record id %s", Record_id)))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, exhibitions)
	}
}

// GetRecordById godoc
// @Summary Retrieves Record based on given ID
// @Description get Record by ID
// @Produce json
// @Param id path integer true "Record ID"
// @Success 200 {object} Records.Record
// @Router /api/Records/{id} [get]
// @Security Authorization Token
func GetAuctionsByRecordId(c *gin.Context) {
	s := persistence.GetRecordRepository()
	Record_id := c.Param("id")
	if auctions, err := s.GetAuctionRecord(Record_id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New(fmt.Sprintf("Auctions not found for Record id %s", Record_id)))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, auctions)
	}
}

// GetRecords godoc
// @Summary Retrieves Records based on query
// @Description Get Records
// @Produce json
// @Param Recordname query string false "Recordname"
// @Param firstname query string false "Firstname"
// @Param lastname query string false "Lastname"
// @Success 200 {array} []Records.Record
// @Router /api/Records [get]
// @Security Authorization Token
func GetRecords(c *gin.Context) {
	s := persistence.GetRecordRepository()
	var q models.Record
	_ = c.Bind(&q)
	if Records, err := s.Query(&q); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("Records not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, Records)
	}
}

func CreateRecord(c *gin.Context) {
	s := persistence.GetRecordRepository()
	var RecordInput RecordInput
	_ = c.BindJSON(&RecordInput)
	Record := models.Record{
		// Recordname:  RecordInput.Recordname,
		// Firstname: RecordInput.Firstname,
		// Lastname:  RecordInput.Lastname,
		// Hash:      crypto.HashAndSalt([]byte(RecordInput.Password)),
		// Role:      models.RecordRole{RoleName: RecordInput.Role},
	}
	if err := s.Add(&Record); err != nil {
		http_err.NewError(c, http.StatusBadRequest, err)
		log.Println(err)
	} else {
		c.JSON(http.StatusCreated, Record)
	}
}

func UpdateRecord(c *gin.Context) {
	s := persistence.GetRecordRepository()
	id := c.Params.ByName("id")
	var RecordInput RecordInput
	_ = c.BindJSON(&RecordInput)
	if Record, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("Record not found"))
		log.Println(err)
	} else {
		// Record.Recordname = RecordInput.Recordname
		// Record.Lastname = RecordInput.Lastname
		// Record.Firstname = RecordInput.Firstname
		// Record.Hash = crypto.HashAndSalt([]byte(RecordInput.Password))
		// Record.Role = models.RecordRole{RoleName: RecordInput.Role}
		if err := s.Update(Record); err != nil {
			http_err.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, Record)
		}
	}
}

func DeleteRecord(c *gin.Context) {
	s := persistence.GetRecordRepository()
	id := c.Params.ByName("id")
	var RecordInput RecordInput
	_ = c.BindJSON(&RecordInput)
	if Record, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("Record not found"))
		log.Println(err)
	} else {
		if err := s.Delete(Record); err != nil {
			http_err.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusNoContent, "")
		}
	}
}
