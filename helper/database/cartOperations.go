package database

import (
	"fmt"
	entity "kwick/model/entity"
	"strconv"
)

func RemoveFromCart(ci, pi string) bool {
	var Exist entity.Cart
	DB.Where("id = ? AND product_id = ?", ci, pi).Find(&Exist)
	if Exist.ID == 0 {
		return false
	}
	var Cart entity.Cart
	DB.Where("id = ? AND product_id = ?", ci, pi).Delete(&Cart)
	return true
}
func ReduceCount(ci, pi, ui string) bool {
	var cart entity.Cart
	DB.Where("id = ? AND product_id = ? AND user_id = ?", ci, pi, ui).Find(&cart)
	if cart.Count == 1 {
		res := RemoveFromCart(ci, pi)
		return res
	}
	if cart.ID == 0 {
		return false
	}
	cart.Count--
	DB.Save(&cart)
	return true
}
func IncreaseCount(ci, pi, ui string) bool {
	var cart entity.Cart
	DB.Where("id = ? AND product_id = ? AND user_id = ?", ci, pi, ui).Find(&cart)
	if cart.Count == 0 {
		return false
	}
	cart.Count++
	DB.Save(&cart)
	return true
}
func GetUserCart(i string) []entity.CartResponse {
	var CartResponse []entity.CartResponse
	var cart []entity.Cart
	var item entity.CartResponse
	DB.Where("user_id = ?", i).Find(&cart)
	for _, v := range cart {
		pid := fmt.Sprintf("%d", v.ProductId)
		product, _ := GetProductById(pid)
		p, _ := strconv.ParseUint(pid, 10, 64)
		images := GetImagesByProductId(uint(p))
		productResponse := entity.ProductResponse{
			Product: product,
			Images:  images,
		}
		item.Product = productResponse
		item.CartId = v.ID
		item.Count = uint(v.Count)
		CartResponse = append(CartResponse, item)
	}
	return CartResponse
}
func AddProductToCart(pi, ui string) bool {
	var Exist entity.Cart
	DB.Where("user_id = ? AND product_id = ?", ui, pi).Find(&Exist)
	if Exist.ID != 0 {
		Exist.Count++
		DB.Save(&Exist)
		return true
	}
	var cart entity.Cart
	uId, _ := strconv.ParseUint(ui, 10, 64)
	pId, _ := strconv.ParseUint(pi, 10, 64)
	cart.ProductId = uint(pId)
	cart.UserId = uint(uId)
	cart.Count = 1
	DB.Create(&cart)
	return true
}
func AddProductToWishlist(pi, ui string) bool {
	var Exist entity.Wishlist
	DB.Where("user_id = ? AND product_id = ?", ui, pi).Find(&Exist)
	if Exist.ID != 0 {
		return false
	}
	var cart entity.Wishlist
	uId, _ := strconv.ParseUint(ui, 10, 64)
	pId, _ := strconv.ParseUint(pi, 10, 64)
	cart.ProductId = uint(pId)
	cart.UserId = uint(uId)
	DB.Create(&cart)
	return true
}
func RemoveFromWishlist(ci, pi string) bool {
	var Exist entity.Wishlist
	DB.Where("id = ? AND product_id = ?", ci, pi).Find(&Exist)
	if Exist.ID == 0 {
		return false
	}
	var Cart entity.Wishlist
	DB.Where("id = ? AND product_id = ?", ci, pi).Delete(&Cart)
	return true
}
