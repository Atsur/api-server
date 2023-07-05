package persistence

import (
	"strconv"

	"github.com/atsur/api-server/internal/pkg/db"
	models "github.com/atsur/api-server/internal/pkg/models/artifacts"
)

type ArtifactRepository struct{}

var artifactRepository *ArtifactRepository

func GetArtifactRepository() *ArtifactRepository {
	if artifactRepository == nil {
		artifactRepository = &ArtifactRepository{}
	}
	return artifactRepository
}

func (r *ArtifactRepository) Get(id string) (*models.Artifact, error) {
	var Artifact models.Artifact
	where := models.Artifact{}
	where.ID, _ = strconv.ParseUint(id, 10, 64)
	_, err := First(&where, &Artifact, []string{"Artifact"})
	if err != nil {
		return nil, err
	}
	return &Artifact, err
}

func (r *ArtifactRepository) GetProvenanceByArtifact(artifactId string) (*[]models.Provenance, error) {
	var Provenance []models.Provenance
	// where := models.Provenance{}
	// artifactId, _ := strconv.ParseUint(id, 10, 64)
	err := Find("artifact_id = ?", &Provenance, []string{artifactId}, "id asc")
	if err != nil {
		return nil, err
	}
	return &Provenance, err
}

func (r *ArtifactRepository) GetExhibitionByArtifact(artifactId string) (*[]models.Exhibition, error) {
	var Exhibition []models.Exhibition
	// where := models.Provenance{}
	// artifactId, _ := strconv.ParseUint(id, 10, 64)
	err := Find("artifact_id = ?", &Exhibition, []string{artifactId}, "id asc")
	if err != nil {
		return nil, err
	}
	return &Exhibition, err
}

func (r *ArtifactRepository) GetAuctionArtifact(artifactId string) (*[]models.Auction, error) {
	var Auction []models.Auction
	// where := models.Provenance{}
	// artifactId, _ := strconv.ParseUint(id, 10, 64)
	err := Find("artifact_id = ?", &Auction, []string{artifactId}, "id asc")
	if err != nil {
		return nil, err
	}
	return &Auction, err
}

func (r *ArtifactRepository) All() (*[]models.Artifact, error) {
	var Artifacts []models.Artifact
	err := Find(&models.Artifact{}, &Artifacts, []string{"Artifact"}, "id asc")
	return &Artifacts, err
}

func (r *ArtifactRepository) Query(q *models.Artifact) (*[]models.Artifact, error) {
	var Artifacts []models.Artifact
	err := Find(&q, &Artifacts, []string{"Artifact"}, "id asc")
	return &Artifacts, err
}

func (r *ArtifactRepository) Add(Artifact *models.Artifact) error {
	err := Create(&Artifact)
	err = Save(&Artifact)
	return err
}

func (r *ArtifactRepository) Update(Artifact *models.Artifact) error {
	return db.GetDB().Omit("Artifact").Save(&Artifact).Error
}

func (r *ArtifactRepository) Delete(Artifact *models.Artifact) error {
	return db.GetDB().Unscoped().Delete(&Artifact).Error
}
