package models

import (
	"time"
)

type Asset struct {
	Model
	Title                string    `gorm:"column:title;not null;" json:"title" form:"title"`
	Type                 string    `gorm:"column:type;not null;" json:"type" form:"type"`
	Description          string    `gorm:"column:description;not null;" json:"description" form:"description"`
	TwoDimensionPublic   string    `gorm:"column:2d_public;not null;" json:"2d_public" form:"2d_public"`
	TwoDimensionFront    string    `gorm:"column:2d_front;not null;" json:"2d_front" form:"2d_front"`
	TwoDimensionSide     string    `gorm:"column:2d_side;not null;" json:"2d_side" form:"2d_side"`
	ThreeDimensionPublic string    `gorm:"column:3d_public;not null;" json:"3d_public" form:"3d_public"`
	RecordNo             uint64    `gorm:"column:record_no;not null;" json:"record_no" form:"record_no"`
	AmountSold           float64   `gorm:"column:amount_sold;" json:"amount_sold" form:"amount_sold"`
	DateStarted          time.Time `gorm:"column:date_started;type:datetime;not null;" json:"date_started"`
	DateEnded            time.Time `gorm:"column:date_ended;type:datetime;not null;" json:"date_ended"`
	Record               Record    `json:"record"`
}

func (m *Asset) BeforeCreate() error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

func (m *Asset) BeforeUpdate() error {
	m.UpdatedAt = time.Now()
	return nil
}
