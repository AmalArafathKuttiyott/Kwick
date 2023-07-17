package models

import "time"

type Order struct {
	ID          uint `gorm:"primaryKey"`
	UserID      uint
	Address     uint
	ProductId   uint
	Amount      float64
	OrderStatus uint
	Payment     uint
	Date        time.Time
	Quantity    int
}

type Status struct {
	Id     uint `gorm:"primaryKey"`
	Status string
}

type OrdersResponse struct {
	Order       Order
	Product     ProductResponse
	Address     Address
	OrderStatus string
}

type AllOrdersResponse struct {
	Order       Order
	User        User
	Product     ProductResponse
	Address     Address
	OrderStatus string
}

type Payment struct {
	Id     uint `gorm:"primarykey"`
	Method string
}
