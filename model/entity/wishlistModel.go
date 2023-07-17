package models

type Wishlist struct {
	ID        uint `gorm:"primaryKey"`
	UserId    uint
	ProductId uint
}
