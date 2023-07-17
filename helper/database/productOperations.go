package database

import (
	"fmt"
	entity "kwick/model/entity"
	request "kwick/model/request"
	"strconv"
	"strings"
)

func GetTotalNumberOfProducts() int {
	var products int64
	if err := DB.Model(entity.Product{}).Count(&products).Error; err != nil {
		return 0
	}
	return int(products)
}
func GetProducts(p, l int) []entity.Product {
	var products []entity.Product
	offset := (int(p) - 1) * int(l)
	result := DB.Offset(offset).Limit(l).Find(&products)
	if result.Error != nil {
		panic("failed to query users: " + result.Error.Error())
	}
	return products
}
func GetProductById(i string) (entity.Product, bool) {
	var product entity.Product
	DB.Where("id = ?", i).Find(&product)
	if product.ID == 0 {
		return product, false
	}
	return product, true
}
func AddProduct(pd request.RequestBody) bool {
	var exist entity.Product
	DB.Where("name = ?", pd.ProductName).Find(&exist)
	if exist.ID != 0 {
		return false
	}
	price := strconv.FormatFloat(pd.ProductPrice, 'f', -1, 64)
	pq := strconv.FormatUint(uint64(pd.ProductQuantity), 10)

	var product entity.Product
	product.Category = pd.CategoryId
	product.Name = pd.ProductName
	product.Price = price
	product.Description = pd.ProductDescription
	product.Availability = true
	product.Quantity = pq

	result := DB.Create(&product)
	return result.Error == nil
}
func GetProductByName(n string) uint {
	var product entity.Product
	DB.Where("name = ?", n).Find(&product)
	if product.ID == 0 {
		return 0
	}
	return product.ID
}
func EditProduct(i string, d request.RequestBody) bool {
	var product entity.Product
	DB.Where("id = ?", i).Find(&product)
	product.Name = d.ProductName
	product.Description = d.ProductDescription
	price := d.ProductPrice
	product.Price = strconv.FormatFloat(price, 'f', 2, 64)
	product.Category = d.CategoryId
	product.Quantity = fmt.Sprint(d.ProductQuantity)
	DB.Save(&product)
	return true
}
func BlockProduct(i string) bool {
	var product entity.Product
	DB.Where("id = ?", i).Find(&product)
	if !product.Availability {
		return false
	}
	product.Availability = false
	DB.Save(&product)
	return true
}
func UnblockProduct(i string) bool {
	var product entity.Product
	DB.Where("id = ?", i).Find(&product)
	if product.Availability {
		return false
	}
	product.Availability = true
	DB.Save(&product)
	return true
}
func GetProductBySearch(c, k string) []entity.Product {
	var products []entity.Product
	DB.Where("category = ? AND LOWER(name) LIKE ?", c, "%"+strings.ToLower(k)+"%").Find(&products)
	return products
}
func GetStatusByCode(i string) string {
	var status entity.Status
	DB.Where("id = ? ", i).Find(&status)
	return status.Status
}
func GetProductId(s string) uint {
	var product entity.Product
	DB.Where("name = ?", s).First(&product)
	return product.ID
}
