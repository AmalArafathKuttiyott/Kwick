package models

import "gorm.io/gorm"

type Vendors struct {
	gorm.Model
	UserId uint
}

type Business struct {
	Id             uint
	Name           string
	BuildingName   string
	BuildingNumber string
	City           string
	Country        string
	PostOffice     string
	Pincode        string
	Gst            string
	Pan            string
	Verified       bool
	Rejected       bool
}
