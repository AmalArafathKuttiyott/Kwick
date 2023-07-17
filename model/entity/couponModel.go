package models

import "time"

type Coupon struct {
	Id            uint      `gorm:"primaryKey;not null"`
	Name          string    `gorm:"not null"`
	Code          string    `gorm:"not null"`
	Limit         uint      `gorm:"not null"`
	AvailableFrom time.Time `gorm:"not null"`
	AvailableTill time.Time `gorm:"not null"`
	Percentage    uint      `gorm:"not null"`
	MaxDiscount   uint      `gorm:"not null"`
	MinPurchase   uint      `gorm:"not null"`
}

type Referral struct {
	Code            string `gorm:"not null"`
	ByUser          uint   `gorm:"not null"`
	ToUser          uint   `gorm:"not null"`
	AmountForByUser uint   `gorm:"not null;defult:0"`
	AmountForToUser uint   `gorm:"not null;defult:0"`
}
