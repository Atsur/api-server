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

type ExhibitionInput struct {
	ExhibitionName string `json:"ExhibitionName" binding:"required"`
	Lastname       string `json:"lastname"`
	Firstname      string `json:"firstname"`
	Password       string `json:"password" binding:"required"`
	Role           string `json:"role"`
}

// GetExhibitionById godoc
// @Summary Retrieves Exhibition based on given ID
// @Description get Exhibition by ID
// @Produce json
// @Param id path integer true "Exhibition ID"
// @Success 200 {object} Exhibitions.Exhibition
// @Router /api/Exhibitions/{id} [get]
// @Security Authorization Token
func GetExhibitionById(c *gin.Context) {
	s := persistence.GetExhibitionRepository()
	id := c.Param("id")
	if Exhibition, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("Exhibition not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, Exhibition)
	}
}

// GetExhibitions godoc
// @Summary Retrieves Exhibitions based on query
// @Description Get Exhibitions
// @Produce json
// @Param Exhibitionname query string false "Exhibitionname"
// @Param firstname query string false "Firstname"
// @Param lastname query string false "Lastname"
// @Success 200 {array} []Exhibitions.Exhibition
// @Router /api/Exhibitions [get]
// @Security Authorization Token
func GetExhibitions(c *gin.Context) {
	s := persistence.GetExhibitionRepository()
	var q models.Exhibition
	_ = c.Bind(&q)
	if Exhibitions, err := s.Query(&q); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("Exhibitions not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, Exhibitions)
	}
}

func CreateExhibition(c *gin.Context) {
	s := persistence.GetExhibitionRepository()
	var ExhibitionInput ExhibitionInput
	_ = c.BindJSON(&ExhibitionInput)
	Exhibition := models.Exhibition{
		// Exhibitionname:  ExhibitionInput.Exhibitionname,
		// Firstname: ExhibitionInput.Firstname,
		// Lastname:  ExhibitionInput.Lastname,
		// Hash:      crypto.HashAndSalt([]byte(ExhibitionInput.Password)),
		// Role:      models.ExhibitionRole{RoleName: ExhibitionInput.Role},
	}
	if err := s.Add(&Exhibition); err != nil {
		http_err.NewError(c, http.StatusBadRequest, err)
		log.Println(err)
	} else {
		c.JSON(http.StatusCreated, Exhibition)
	}
}

func UpdateExhibition(c *gin.Context) {
	s := persistence.GetExhibitionRepository()
	id := c.Params.ByName("id")
	var ExhibitionInput ExhibitionInput
	_ = c.BindJSON(&ExhibitionInput)
	if Exhibition, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("Exhibition not found"))
		log.Println(err)
	} else {
		// Exhibition.Exhibitionname = ExhibitionInput.Exhibitionname
		// Exhibition.Lastname = ExhibitionInput.Lastname
		// Exhibition.Firstname = ExhibitionInput.Firstname
		// Exhibition.Hash = crypto.HashAndSalt([]byte(ExhibitionInput.Password))
		// Exhibition.Role = models.ExhibitionRole{RoleName: ExhibitionInput.Role}
		if err := s.Update(Exhibition); err != nil {
			http_err.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, Exhibition)
		}
	}
}

func DeleteExhibition(c *gin.Context) {
	s := persistence.GetExhibitionRepository()
	id := c.Params.ByName("id")
	var ExhibitionInput ExhibitionInput
	_ = c.BindJSON(&ExhibitionInput)
	if Exhibition, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("Exhibition not found"))
		log.Println(err)
	} else {
		if err := s.Delete(Exhibition); err != nil {
			http_err.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusNoContent, "")
		}
	}
}
