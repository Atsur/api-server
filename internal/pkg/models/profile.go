package models

import (
	"time"
)

type Profile struct {
	Model
	FullName          string                 `gorm:"column:full_name;not null;" json:"full_name" form:"full_name"`
	NickName          string                 `gorm:"column:nick_name;not null;" json:"nick_name" form:"nick_name"`
	Phone             string                 `gorm:"column:phone;not null;" json:"phone" form:"phone"`
	Bio               string                 `gorm:"column:bio;" json:"bio" form:"bio"`
	Avatar            string                 `gorm:"column:avatar;" json:"avatar" form:"avatar"`
	Address           string                 `gorm:"column:address;" json:"address" form:"address"`
	BusinessName      string                 `gorm:"column:business_name;" json:"business_name" form:"business_name"`
	BusinessAddress   string                 `gorm:"column:business_address;" json:"business_address" form:"business_address"`
	UserID            uint64                 `gorm:"column:user_id;unique_index:user_id;" json:"user_id" form:"user_id"`
	AdditionalImgUrls []string               `gorm:"column:additional_img_urls;" json:"additional_img_urls" form:"additional_img_urls"`
	Socials           map[string]interface{} `gorm:"column:socials;" json:"socials" form:"socials"`
	Type              ProfileType            `gorm:"column:type;" json:"type" form:"type"`
	User              User                   `json:"user"`
}

type ProfileType string

const (
	CustodianProfile   ProfileType = "custodian"
	GalleryProfile     ProfileType = "gallery"
	ArtistProfile      ProfileType = "artist"
	CollectorProfile   ProfileType = "collector"
	InstitutionProfile ProfileType = "institution"
)

func (m *Profile) BeforeCreate() error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

func (m *Profile) BeforeUpdate() error {
	m.UpdatedAt = time.Now()
	return nil
}
