package models

import (
	"time"
)

type VerificationType string

const (
	InternalVerification VerificationType = "internal"
	ExternalVerification VerificationType = "external"
)

type VerificationStatus string

const (
	PendingVerification  VerificationStatus = "pending"
	VerificationComplete VerificationStatus = "complete"
)

type CertificateOfAuthenticity struct {
	Model
	Title           string            `gorm:"column:title;not null;" json:"title" form:"title"`
	ArtistFullName  string            `gorm:"column:artistName;not null;" json:"artist_name" form:"artist_name"`
	YearCreated     string            `gorm:"column:year_created;not null;" json:"year_created" form:"year_created"`
	Medium          string            `gorm:"column:medium;not null;" json:"medium" form:"medium"`
	Size            string            `gorm:"column:size;not null;" json:"size" form:"size"`
	Status          CertificateStatus `gorm:"column:status;not null;" json:"status" form:"status"`
	RecordNo        uint64            `gorm:"column:record_no;not null;" json:"record_no" form:"record_no"`
	ArtistProfileID uint64            `gorm:"column:artist_profile_id;" json:"artist_profile_id" form:"artist_profile_id"`
	Record          Record            `json:"artifact"`
	ArtistProfile   Profile           `json:"artist_profile"`
}

type CertificateStatus string

const (
	DraftCertificate  CertificateStatus = "draft"
	ReadyCertificate  CertificateStatus = "ready"
	MintedCertificate CertificateStatus = "minted"
)

func (m *CertificateOfAuthenticity) BeforeCreate() error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

func (m *CertificateOfAuthenticity) BeforeUpdate() error {
	m.UpdatedAt = time.Now()
	return nil
}
