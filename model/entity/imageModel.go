package models

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	ProductId uint   `json:"productId"`
	Link      string `json:"imageLink"`
}
