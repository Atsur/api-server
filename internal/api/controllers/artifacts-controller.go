package controllers

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	models "github.com/atsur/api-server/internal/pkg/models/artifacts"
	"github.com/atsur/api-server/internal/pkg/persistence"
	// "github.com/atsur/api-server/pkg/crypto"
	"github.com/atsur/api-server/pkg/http_err"
	"github.com/gin-gonic/gin"
)

type ArtifactInput struct {
	Artifactname string `json:"Artifactname" binding:"required"`
	Lastname     string `json:"lastname"`
	Firstname    string `json:"firstname"`
	Password     string `json:"password" binding:"required"`
	Role         string `json:"role"`
}

// GetArtifactById godoc
// @Summary Retrieves Artifact based on given ID
// @Description get Artifact by ID
// @Produce json
// @Param id path integer true "Artifact ID"
// @Success 200 {object} Artifacts.Artifact
// @Router /api/Artifacts/{id} [get]
// @Security Authorization Token
func GetArtifactById(c *gin.Context) {
	s := persistence.GetArtifactRepository()
	id := c.Param("id")
	if Artifact, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("Artifact not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, Artifact)
	}
}

// GetArtifactById godoc
// @Summary Retrieves Artifact based on given ID
// @Description get Artifact by ID
// @Produce json
// @Param id path integer true "Artifact ID"
// @Success 200 {object} Artifacts.Artifact
// @Router /api/Artifacts/{id} [get]
// @Security Authorization Token
func GetProvenanceByArtifactId(c *gin.Context) {
	s := persistence.GetArtifactRepository()
	artifact_id := c.Param("id")
	if provenance, err := s.GetProvenanceByArtifact(artifact_id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New(fmt.Sprintf("Provenance not found for artifact id %s", artifact_id)))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, provenance)
	}
}

// GetArtifactById godoc
// @Summary Retrieves Artifact based on given ID
// @Description get Artifact by ID
// @Produce json
// @Param id path integer true "Artifact ID"
// @Success 200 {object} Artifacts.Artifact
// @Router /api/Artifacts/{id} [get]
// @Security Authorization Token
func GetExhibitionByArtifactId(c *gin.Context) {
	s := persistence.GetArtifactRepository()
	artifact_id := c.Param("id")
	if exhibitions, err := s.GetExhibitionByArtifact(artifact_id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New(fmt.Sprintf("Exhibitions not found for artifact id %s", artifact_id)))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, exhibitions)
	}
}

// GetArtifactById godoc
// @Summary Retrieves Artifact based on given ID
// @Description get Artifact by ID
// @Produce json
// @Param id path integer true "Artifact ID"
// @Success 200 {object} Artifacts.Artifact
// @Router /api/Artifacts/{id} [get]
// @Security Authorization Token
func GetAuctionsByArtifactId(c *gin.Context) {
	s := persistence.GetArtifactRepository()
	artifact_id := c.Param("id")
	if auctions, err := s.GetAuctionArtifact(artifact_id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New(fmt.Sprintf("Auctions not found for artifact id %s", artifact_id)))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, auctions)
	}
}

// GetArtifacts godoc
// @Summary Retrieves Artifacts based on query
// @Description Get Artifacts
// @Produce json
// @Param Artifactname query string false "Artifactname"
// @Param firstname query string false "Firstname"
// @Param lastname query string false "Lastname"
// @Success 200 {array} []Artifacts.Artifact
// @Router /api/Artifacts [get]
// @Security Authorization Token
func GetArtifacts(c *gin.Context) {
	s := persistence.GetArtifactRepository()
	var q models.Artifact
	_ = c.Bind(&q)
	if Artifacts, err := s.Query(&q); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("Artifacts not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, Artifacts)
	}
}

func CreateArtifact(c *gin.Context) {
	s := persistence.GetArtifactRepository()
	var ArtifactInput ArtifactInput
	_ = c.BindJSON(&ArtifactInput)
	Artifact := models.Artifact{
		// Artifactname:  ArtifactInput.Artifactname,
		// Firstname: ArtifactInput.Firstname,
		// Lastname:  ArtifactInput.Lastname,
		// Hash:      crypto.HashAndSalt([]byte(ArtifactInput.Password)),
		// Role:      models.ArtifactRole{RoleName: ArtifactInput.Role},
	}
	if err := s.Add(&Artifact); err != nil {
		http_err.NewError(c, http.StatusBadRequest, err)
		log.Println(err)
	} else {
		c.JSON(http.StatusCreated, Artifact)
	}
}

func UpdateArtifact(c *gin.Context) {
	s := persistence.GetArtifactRepository()
	id := c.Params.ByName("id")
	var ArtifactInput ArtifactInput
	_ = c.BindJSON(&ArtifactInput)
	if Artifact, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("Artifact not found"))
		log.Println(err)
	} else {
		// Artifact.Artifactname = ArtifactInput.Artifactname
		// Artifact.Lastname = ArtifactInput.Lastname
		// Artifact.Firstname = ArtifactInput.Firstname
		// Artifact.Hash = crypto.HashAndSalt([]byte(ArtifactInput.Password))
		// Artifact.Role = models.ArtifactRole{RoleName: ArtifactInput.Role}
		if err := s.Update(Artifact); err != nil {
			http_err.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, Artifact)
		}
	}
}

func DeleteArtifact(c *gin.Context) {
	s := persistence.GetArtifactRepository()
	id := c.Params.ByName("id")
	var ArtifactInput ArtifactInput
	_ = c.BindJSON(&ArtifactInput)
	if Artifact, err := s.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("Artifact not found"))
		log.Println(err)
	} else {
		if err := s.Delete(Artifact); err != nil {
			http_err.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusNoContent, "")
		}
	}
}
