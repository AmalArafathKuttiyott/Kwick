package database

import (
	"fmt"
	entity "kwick/model/entity"
)

func GetOffers(p, l int) []entity.Product {
	var products []entity.Product
	DB.Where("discount_status = ?", true).Find(&products)
	return products
}
func GetTotalNumberOfOffers() int {
	var offers int64
	if err := DB.Where("discount_status = ?", true).Model(entity.Product{}).Count(&offers).Error; err != nil {
		return 0
	}
	return int(offers)
}
func AddProductOfferById(i, p string) bool {
	var product entity.Product
	DB.Where("id = ?", i).Find(&product)
	fmt.Println(product)
	if product.ID == 0 {
		return false
	}
	if product.DiscountStatus {
		return false
	}
	product.DiscountStatus = true
	product.DiscountPrice = p
	DB.Save(&product)
	return true
}
func AddCategoryOfferById(i, p string) bool {
	product := &entity.Product{}
	if err := DB.Model(product).Where("category = ?", i).Find(product).Error; err != nil {
		return false
	}
	updates := map[string]interface{}{
		"discount_status": true,
		"discount_price":  p,
	}
	if err := DB.Model(product).Where("category = ?", i).Updates(updates).Error; err != nil {
		return false
	}

	return true
}
func RemoveProductOfferById(i string) bool {
	var product entity.Product
	DB.Where("id = ?", i).Find(&product)
	fmt.Println(product)
	if product.ID == 0 {
		return false
	}
	if !product.DiscountStatus {
		return false
	}
	product.DiscountStatus = false
	product.DiscountPrice = "0"
	DB.Save(&product)
	return true
}
func RemoveCategoryOfferById(i string) bool {
	product := &entity.Product{}
	if err := DB.Model(product).Where("category = ?", i).Find(product).Error; err != nil {
		return false
	}
	updates := map[string]interface{}{
		"discount_status": false,
		"discount_price":  0,
	}
	if err := DB.Model(product).Where("category = ?", i).Updates(updates).Error; err != nil {
		return false
	}

	return true
}
func EditProductOfferById(i, p string) bool {
	var product entity.Product
	DB.Where("id = ?", i).Find(&product)
	fmt.Println(product)
	if product.ID == 0 {
		return false
	}
	if !product.DiscountStatus {
		return false
	}
	product.DiscountPrice = p
	DB.Save(&product)
	return true
}
