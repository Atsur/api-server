package controllers

import (
	"errors"
	"log"
	"net/http"

	models "github.com/atsur/api-server/internal/pkg/models"
	"github.com/atsur/api-server/internal/pkg/persistence"
	// "github.com/atsur/api-server/pkg/crypto"
	"github.com/atsur/api-server/pkg/http_err"
	"github.com/gin-gonic/gin"
)

type ProvenanceInput struct {
	ProvenanceName string `json:"ProvenanceName" binding:"required"`
	Lastname       string `json:"lastname"`
	Firstname      string `json:"firstname"`
	Password       string `json:"password" binding:"required"`
	Role           string `json:"role"`
}

// GetProvenanceById godoc
// @Summary Retrieves Provenance based on given ID
// @Description get Provenance by ID
// @Produce json
// @Param id path integer true "Provenance ID"
// @Success 200 {object} Provenances.Provenance
// @Router /api/Provenances/{id} [get]
// @Security Authorization Token
func GetProvenanceById(c *gin.Context) {
	s := persistence.GetProvenanceRepository()
	id := c.Param("id")
	if Provenance, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("Provenance not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, Provenance)
	}
}

// GetProvenances godoc
// @Summary Retrieves Provenances based on query
// @Description Get Provenances
// @Produce json
// @Param Provenancename query string false "Provenancename"
// @Param firstname query string false "Firstname"
// @Param lastname query string false "Lastname"
// @Success 200 {array} []Provenances.Provenance
// @Router /api/Provenances [get]
// @Security Authorization Token
func GetProvenances(c *gin.Context) {
	s := persistence.GetProvenanceRepository()
	var q models.Provenance
	_ = c.Bind(&q)
	if Provenances, err := s.Query(&q); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("Provenances not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, Provenances)
	}
}

func CreateProvenance(c *gin.Context) {
	s := persistence.GetProvenanceRepository()
	var ProvenanceInput ProvenanceInput
	_ = c.BindJSON(&ProvenanceInput)
	Provenance := models.Provenance{
		// Provenancename:  ProvenanceInput.Provenancename,
		// Firstname: ProvenanceInput.Firstname,
		// Lastname:  ProvenanceInput.Lastname,
		// Hash:      crypto.HashAndSalt([]byte(ProvenanceInput.Password)),
		// Role:      models.ProvenanceRole{RoleName: ProvenanceInput.Role},
	}
	if err := s.Add(&Provenance); err != nil {
		http_err.NewError(c, http.StatusBadRequest, err)
		log.Println(err)
	} else {
		c.JSON(http.StatusCreated, Provenance)
	}
}

func UpdateProvenance(c *gin.Context) {
	s := persistence.GetProvenanceRepository()
	id := c.Params.ByName("id")
	var ProvenanceInput ProvenanceInput
	_ = c.BindJSON(&ProvenanceInput)
	if Provenance, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("Provenance not found"))
		log.Println(err)
	} else {
		// Provenance.Provenancename = ProvenanceInput.Provenancename
		// Provenance.Lastname = ProvenanceInput.Lastname
		// Provenance.Firstname = ProvenanceInput.Firstname
		// Provenance.Hash = crypto.HashAndSalt([]byte(ProvenanceInput.Password))
		// Provenance.Role = models.ProvenanceRole{RoleName: ProvenanceInput.Role}
		if err := s.Update(Provenance); err != nil {
			http_err.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, Provenance)
		}
	}
}

func DeleteProvenance(c *gin.Context) {
	s := persistence.GetProvenanceRepository()
	id := c.Params.ByName("id")
	var ProvenanceInput ProvenanceInput
	_ = c.BindJSON(&ProvenanceInput)
	if Provenance, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("Provenance not found"))
		log.Println(err)
	} else {
		if err := s.Delete(Provenance); err != nil {
			http_err.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusNoContent, "")
		}
	}
}
