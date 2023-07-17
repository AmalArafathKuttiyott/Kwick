package models

type Cart struct {
	ID        uint `gorm:"primaryKey"`
	UserId    uint
	ProductId uint
	Count     int
}

type CartResponse struct {
	CartId  uint
	Count   uint
	Product ProductResponse
}
