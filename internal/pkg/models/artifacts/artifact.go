package artifacts

import (
	"time"

	"github.com/atsur/api-server/internal/pkg/models"
	"github.com/atsur/api-server/internal/pkg/models/users"
)

type Artifact struct {
	models.Model
	Title                        string          `gorm:"column:title;not null;" json:"title" form:"title"`
	Description                  string          `gorm:"column:description;not null;" json:"description" form:"description"`
	Type                         string          `gorm:"column:type;not null;" json:"type" form:"type"`
	Edition                      string          `gorm:"column:edition;not null;" json:"edition" form:"edition"`
	Series                       string          `gorm:"column:series;" json:"series" form:"series"`
	YearCreated                  string          `gorm:"column:year_created;not null;" json:"year_created" form:"year_created"`
	Metadata                     string          `gorm:"column:metadata;" json:"metadata" form:"metadata"`
	Frame                        string          `gorm:"column:frame;" json:"frame" form:"frame"`
	Classification               string          `gorm:"column:classification;not null;" json:"classification" form:"classification"`
	CurrentCondition             string          `gorm:"column:current_condition;not null;" json:"current_condition" form:"current_condition"`
	CurrentLocation              string          `gorm:"column:current_location;not null;" json:"current_location" form:"current_location"`
	CertificateOfOwnershipUrl    string          `gorm:"column:coo_url;" json:"coo_url" form:"coo_url"`
	CertificateOfAuthenticityUrl string          `gorm:"column:coa_url;" json:"coa_url" form:"coa_url"`
	Signature                    string          `gorm:"column:signature;" json:"signature" form:"signature"`
	CopyrightUrl                 string          `gorm:"column:copyright_url;" json:"copyright_url" form:"copyright_url"`
	Dimensions                   []string        `gorm:"column:dimensions;not null;" json:"dimensions" form:"dimensions"`
	Materials                    []string        `gorm:"column:materials;not null;" json:"materials" form:"materials"`
	Mediums                      []string        `gorm:"column:mediums;not null;" json:"mediums" form:"mediums"`
	AdditionalImgUrls            []string        `gorm:"column:additional_img_urls;" json:"additional_img_urls" form:"additional_img_urls"`
	ArtistProfileIDs             []uint64        `gorm:"column:artist_profile_ids;not null;" json:"artist_profile_ids" form:"artist_profile_ids"`
	ArtistProfiles               []users.Profile `json:"artistProfiles"`
	Provenance                   []Provenance    `gorm:"foreignKey:artifact_id" json:"provenance"`
	Exhibitions                  []Exhibition    `gorm:"foreignKey:artifact_id" json:"exhibitions"`
	Auctions                     []Auction       `gorm:"foreignKey:artifact_id" json:"auctions"`
}

func (m *Artifact) BeforeCreate() error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

func (m *Artifact) BeforeUpdate() error {
	m.UpdatedAt = time.Now()
	return nil
}
