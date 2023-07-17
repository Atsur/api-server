package models

import (
	"time"
)

type Copyright struct {
	Model
	Title            string    `gorm:"column:title;not null;" json:"title" form:"title"`
	Type             string    `gorm:"column:type;not null;" json:"type" form:"type"`
	Description      string    `gorm:"column:description;not null;" json:"description" form:"description"`
	RecordNo         uint64    `gorm:"column:record_no;not null;" json:"record_no" form:"record_no"`
	AuctionProfileID uint64    `gorm:"column:auction_profile_id;" json:"auction_profile_id" form:"auction_profile_id"`
	AmountSold       float64   `gorm:"column:amount_sold;" json:"amount_sold" form:"amount_sold"`
	DateStarted      time.Time `gorm:"column:date_started;type:datetime;not null;" json:"date_started"`
	DateEnded        time.Time `gorm:"column:date_ended;type:datetime;not null;" json:"date_ended"`
	Record           Record    `json:"artifact"`
	AuctionProfile   Profile   `json:"auction_profile"`
}

func (m *Copyright) BeforeCreate() error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

func (m *Copyright) BeforeUpdate() error {
	m.UpdatedAt = time.Now()
	return nil
}
