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

// GetProfileById godoc
// @Summary Retrieves Profile based on given ID
// @Description get Profile by ID
// @Produce json
// @Param id path integer true "Profile ID"
// @Success 200 {object} Profiles.Profile
// @Router /api/Profiles/{id} [get]
// @Security Authorization Token
func GetProfileByUserId(c *gin.Context) {
	s := persistence.GetProfileRepository()
	userId := c.Param("user_id")
	if profile, err := s.GetByUserId(userId); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("Profile not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, profile)
	}
}

// GetProfileById godoc
// @Summary Retrieves Profile based on given ID
// @Description get Profile by ID
// @Produce json
// @Param id path integer true "Profile ID"
// @Success 200 {object} Profiles.Profile
// @Router /api/Profiles/{id} [get]
// @Security Authorization Token
func GetProfileById(c *gin.Context) {
	s := persistence.GetProfileRepository()
	id := c.Param("id")
	if Profile, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("Profile not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, Profile)
	}
}

// GetProfiles godoc
// @Summary Retrieves Profiles based on query
// @Description Get Profiles
// @Produce json
// @Param Profilename query string false "Profilename"
// @Param firstname query string false "Firstname"
// @Param lastname query string false "Lastname"
// @Success 200 {array} []Profiles.Profile
// @Router /api/Profiles [get]
// @Security Authorization Token
func GetProfiles(c *gin.Context) {
	s := persistence.GetProfileRepository()
	var q models.Profile
	_ = c.Bind(&q)
	if Profiles, err := s.Query(&q); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("Profiles not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, Profiles)
	}
}

func CreateProfile(c *gin.Context) {
	s := persistence.GetProfileRepository()
	var ProfileInput models.Profile
	_ = c.BindJSON(&ProfileInput)
	if err := s.Add(&ProfileInput); err != nil {
		http_err.NewError(c, http.StatusBadRequest, err)
		log.Println(err)
	} else {
		c.JSON(http.StatusCreated, ProfileInput)
	}
}

func UpdateProfile(c *gin.Context) {
	s := persistence.GetProfileRepository()
	id := c.Params.ByName("id")
	var ProfileInput models.Profile
	_ = c.BindJSON(&ProfileInput)
	if _, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("Profile not found"))
		log.Println(err)
	} else {
		if err := s.Update(&ProfileInput); err != nil {
			http_err.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, ProfileInput)
		}
	}
}

func DeleteProfile(c *gin.Context) {
	s := persistence.GetProfileRepository()
	id := c.Params.ByName("id")
	var ProfileInput models.Profile
	_ = c.BindJSON(&ProfileInput)
	if Profile, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("Profile not found"))
		log.Println(err)
	} else {
		if err := s.Delete(Profile); err != nil {
			http_err.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusNoContent, "")
		}
	}
}
