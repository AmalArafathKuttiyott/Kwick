package models

type Product struct {
	ID             uint   `gorm:"primaryKey; increment"`
	Category       uint   `json:"CategoryId"`
	Name           string `json:"productName"`
	Price          string `json:"productPrice"`
	DiscountPrice  string `json:"discountPrice"`
	DiscountStatus bool   `json:"discountStatus"`
	Description    string `json:"description"`
	Availability   bool   `json:"availability"`
	Quantity       string `json:"quantiy"`
	Blocked        bool   `gorm:"default:false" json:"productBlocked"`
}
type Products struct {
	Products []Product
}
type Category struct {
	ID      uint   `gorm:"primaryKey; increment"`
	Name    string `json:"categoryName"`
	Blocked bool   `gorm:"default:false" json:"categoryBlocked"`
}
type Categories struct {
	Categories []Category
}
type ProductResponse struct {
	Product Product
	Images  []Image
}
