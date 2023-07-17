package models

import (
	"time"
)

type Record struct {
	Model
	RecordNo          string         `gorm:"column:record_number;not null;unique_index:record_number" json:"record_number" form:"record_number"`
	Title             string         `gorm:"column:title;not null;" json:"title" form:"title"`
	Description       string         `gorm:"column:description;not null;" json:"description" form:"description"`
	RecordType        RecordType     `gorm:"column:record_type;not null;" json:"record_type" form:"record_type"`
	Edition           string         `gorm:"column:edition;not null;" json:"edition" form:"edition"`
	Series            string         `gorm:"column:series;" json:"series" form:"series"`
	YearCreated       uint16         `gorm:"column:year_created;not null;" json:"year_created" form:"year_created"`
	Metadata          string         `gorm:"column:metadata;" json:"metadata" form:"metadata"`
	Frame             string         `gorm:"column:frame;" json:"frame" form:"frame"`
	Classification    Classification `gorm:"column:classification;not null;" json:"classification" form:"classification"`
	COO_ID            string         `gorm:"column:coo_id;" json:"coo_id" form:"coo_id"`
	COA_ID            string         `gorm:"column:coa_id;" json:"coa_id" form:"coa_id"`
	Signature         string         `gorm:"column:signature;" json:"signature" form:"signature"`
	CopyrightID       string         `gorm:"column:copyright_id;" json:"copyright_id" form:"copyright_id"`
	Dimensions        []string       `gorm:"column:dimensions;not null;" json:"dimensions" form:"dimensions"`
	Materials         []string       `gorm:"column:materials;not null;" json:"materials" form:"materials"`
	Mediums           []string       `gorm:"column:mediums;not null;" json:"mediums" form:"mediums"`
	AdditionalImgUrls []string       `gorm:"column:additional_img_urls;" json:"additional_img_urls" form:"additional_img_urls"`
	ArtistProfileIDs  []uint64       `gorm:"column:artist_profile_ids;not null;" json:"artist_profile_ids" form:"artist_profile_ids"`
	ArtistProfiles    []Profile      `json:"artistProfiles"`
	Provenance        []Provenance   `gorm:"foreignKey:artifact_tag_no" json:"provenance"`
	Exhibitions       []Exhibition   `gorm:"foreignKey:artifact_tag_no" json:"exhibitions"`
	Auctions          []Auction      `gorm:"foreignKey:artifact_tag_no" json:"auctions"`
}

type RecordType string

const (
	ArtifactRecord RecordType = "artifact"
	ArtPieceRecord RecordType = "art-piece"
)

type AppraisalType string

const (
	InternalAppraisal AppraisalType = "internal"
	ExternalAppraisal AppraisalType = "external"
)

type AppraisalStatus string

const (
	PendingAppraisal  AppraisalStatus = "pending"
	AppraisalComplete AppraisalStatus = "complete"
)

func (m *Record) BeforeCreate() error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

func (m *Record) BeforeUpdate() error {
	m.UpdatedAt = time.Now()
	return nil
}
