package models

import (
	"time"
)

type Entry struct {
	Model
	Title             string         `gorm:"column:title;not null;" json:"title,omitempty" form:"title"`
	Description       string         `gorm:"column:description;not null;" json:"description,omitempty" form:"description"`
	Edition           string         `gorm:"column:edition;not null;" json:"edition,omitempty" form:"edition"`
	Series            string         `gorm:"column:series;" json:"series,omitempty" form:"series"`
	Metadata          string         `gorm:"column:metadata;" json:"metadata,omitempty" form:"metadata"`
	Frame             string         `gorm:"column:frame;" json:"frame,omitempty" form:"frame"`
	YearCreated       uint16         `gorm:"column:year_created;not null;" json:"year_created,omitempty" form:"year_created"`
	SubmittedBy       uint64         `gorm:"column:submitted_by;not null;" json:"submitted_by,omitempty" form:"submitted_by"`
	SellingPrice      float64        `gorm:"column:selling_price;" json:"selling_price,omitempty" form:"selling_price"`
	Classification    Classification `gorm:"column:classification;not null;" json:"classification,omitempty" form:"classification"`
	RecordType        RecordType     `gorm:"column:record_type;not null;" json:"record_type,omitempty" form:"record_type"`
	EntryType         EntryType      `gorm:"column:submission_type;not null;" json:"submission_type,omitempty" form:"submission_type"`
	Materials         Materials      `gorm:"column:materials;" json:"materials,omitempty" form:"materials"`
	Medium            Medium         `gorm:"column:medium;" json:"medium,omitempty" form:"medium"`
	SubmittedOn       time.Time      `gorm:"column:submitted_on;type:datetime;not null;" json:"submitted_on,omitempty"`
	Dimensions        []uint16       `gorm:"column:dimensions;" json:"dimensions,omitempty" form:"dimensions"`
	AdditionalImgUrls []string       `gorm:"column:additional_img_urls;" json:"additional_img_urls,omitempty" form:"additional_img_urls"`
	ArtistProfileIDs  []uint64       `gorm:"column:artist_profile_ids;not null;" json:"artist_profile_ids,omitempty" form:"artist_profile_ids"`
	ArtistProfiles    []Profile      `json:"artistProfiles"`
	Record            Record         `json:"record"`
	Submitter         User           `json:"submitter"`
}

type EntryType string

const (
	InternalEntry EntryType = "internal"
	ExternalEntry EntryType = "external"
)

type Classification string

const (
	Unique         Classification = "unique"
	LimitedEdition Classification = "limited-edition"
	OpenEdition    Classification = "open-edition"
	UnknownEdition Classification = "unknown-edition"
)

type Materials string

const (
	Paint Materials = "paint"
)

type Medium string

const (
	ArchitectureArt Medium = "architecture"
	PortfolioArt    Medium = "portfolio"
	DecorativeArt   Medium = "decorative"
	PaperWorkArt    Medium = "paperwork"
	MerchandiseArt  Medium = "merchandise"
	WearableArt     Medium = "wearable"
	JewelryArt      Medium = "jewelry"
	InstallationArt Medium = "installation"
	MixedMediaArt   Medium = "mixed-media"
	PaintingArt     Medium = "painting"
	PerformanceArt  Medium = "performance"
	PhotographyArt  Medium = "photography"
	PrintArt        Medium = "prints"
	PosterArt       Medium = "posters"
	ReproductionArt Medium = "reproduction"
	SculptureArt    Medium = "sculpture"
	TextileArt      Medium = "textile"
	FilmArt         Medium = "film"
	NFTArt          Medium = "nft"
	Other           Medium = "other"
)

func (m *Entry) BeforeCreate() error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

func (m *Entry) BeforeUpdate() error {
	m.UpdatedAt = time.Now()
	return nil
}
