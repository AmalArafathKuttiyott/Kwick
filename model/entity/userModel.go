package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName    string `gorm:"not null" json:"firstname"`
	MiddleName   string `json:"middlename"`
	LastName     string `gorm:"not null" json:"lastname"`
	Email        string `gorm:"unique;not null" json:"email"`
	Phone        string `gorm:"unique;not null" json:"phone"`
	Password     string `gorm:"not null" json:"password"`
	Blocked      bool   `gorm:"default:false"`
	IsAdmin      bool   `gorm:"default:false"`
	IsVendor     bool   `gorm:"default:false"`
	Wallet       uint   `gorm:"default:0"`
	Purchases    uint   `gorm:"default:0"`
	ReferralCode string
}

type Address struct {
	Id             uint
	UserId         uint
	BuildingName   string
	BuildingNumber int
	Street         string
	City           string
	State          string
	Country        string
	PostalCode     string
}

type Users struct {
	Users []User
}

type UserHomePageResponse struct {
	User       User
	Products   []Product
	Categories []Category
}
type UserProfile struct {
	FirstName    string
	MiddleName   string
	LastName     string
	Email        string
	Phone        string
	ReferralCode string
	Wallet       uint
	Address      []Address
}
