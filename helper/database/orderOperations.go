package database

import (
	"database/sql"
	"fmt"
	entity "kwick/model/entity"
	"log"
	"strconv"
	"time"

	"github.com/razorpay/razorpay-go"
)

func GetTotalNumberOfOrders() int {
	var count int64
	if err := DB.Model(entity.Order{}).Count(&count).Error; err != nil {
		return 0
	}
	return int(count)
}
func GetTotalNumberOfDeliveredOrders() int {
	var count int64
	if err := DB.Model(entity.Order{}).Where("discount_status = ?", 6).Count(&count).Error; err != nil {
		return 0
	}
	return int(count)
}
func GetTotalNumberOfReturnOrders() int {
	var count int64
	if err := DB.Model(entity.Order{}).Where("order_status >= ? AND order_status <= ?", 7, 11).Count(&count).Error; err != nil {
		return 0
	}
	return int(count)
}
func GetTotalNumberOfPendingOrders() int {
	var count int64
	if err := DB.Not("order_status >= ? AND order_status <= ?", 6, 11).Model(entity.Order{}).Count(&count).Error; err != nil {
		return 0
	}
	return int(count)
}
func GetTotalAmountofRevenue() int {
	var totalAmount sql.NullFloat64
	if err := DB.Table("orders").Where("order_status = ?", 6).Select("(amount)").Scan(&totalAmount).Error; err != nil {
		return 0
	}
	return int(totalAmount.Float64)
}
func GetTotalNumberOfSales() int {
	var sales int64
	if err := DB.Table("orders").Where("order_status = ? ", 6).Count(&sales).Error; err != nil {
		return 0
	}
	return int(sales)
}
func GetOrders(p, l int) []entity.Order {
	var orders []entity.Order
	DB.Find(&orders)
	return orders
}
func GetDeliveredOrders(p, l int) []entity.Order {
	var orders []entity.Order
	DB.Where("order_status = ?", 6).Find(&orders)
	return orders
}
func GetReturnOrders(p, l int) []entity.Order {
	var orders []entity.Order
	DB.Where("order_status >= ? AND order_status <= ?", 7, 11).Find(&orders)
	return orders
}
func GetPendingOrders(p, l int) []entity.Order {
	var orders []entity.Order
	DB.Not("order_status >= ? AND order_status <= ?", 6, 11).Find(&orders)
	return orders
}
func GetOrderById(i string) (entity.Order, bool) {
	var order entity.Order
	DB.Where("id = ?", i).Find(&order)
	if order.ID == 0 {
		return order, false
	}
	return order, true
}
func EditOrder(i, s string) bool {
	var order entity.Order
	st, _ := strconv.ParseUint(s, 10, 64)
	DB.Model(&order).Where("id = ?", i).Update("order_status", uint(st))
	return true
}
func GetUserOrders(i string) []entity.OrdersResponse {
	var orders []entity.Order
	if err := DB.Where("user_id = ? ", i).Find(&orders).Error; err != nil {
		return []entity.OrdersResponse{}
	}
	var orderlist []entity.OrdersResponse
	for _, v := range orders {
		var order entity.OrdersResponse
		order.Order = v
		sid := fmt.Sprintf("%d", v.OrderStatus)
		pid := fmt.Sprintf("%d", v.ProductId)
		p, _ := strconv.ParseUint(pid, 10, 64)
		product, _ := GetProductById(pid)
		images := GetImagesByProductId(uint(p))
		productResponse := entity.ProductResponse{
			Product: product,
			Images:  images,
		}
		order.Product = productResponse
		order.Address = GetUserAddress(fmt.Sprint(order.Address.Id))
		order.OrderStatus = GetStatusByCode(sid)
		orderlist = append(orderlist, order)
	}

	return orderlist
}
func CancelOrder(i string, uid string) (bool, int) {
	var order entity.Order
	DB.Where("id = ? AND user_id = ?", i, uid).Find(&order)
	if order.ID == 0 {
		return false, 1
	}
	if order.OrderStatus >= 3 && order.OrderStatus <= 7 {
		return false, 2
	}
	order.Date = time.Now()
	order.OrderStatus = 7
	DB.Save(&order)
	var user entity.User
	DB.Where("id = ?", uid).First(&user)
	user.Wallet = user.Wallet + uint(order.Amount)
	DB.Save(&user)
	return true, 0
}
func ReturnOrder(oid string, uid string) bool {
	var order entity.Order
	DB.Where("id = ? AND user_id = ?", oid, uid).Find(&order)
	if order.ID == 0 || order.OrderStatus == 4 || order.OrderStatus == 5 || order.OrderStatus == 7 {
		return false
	}
	order.OrderStatus = 8
	DB.Save(&order)
	var user entity.User
	DB.Where("id = ?", uid).Find(&user)
	user.Wallet += uint(order.Amount)
	return true
}
func AddOrder(id, ai, pm string) (map[string]interface{}, bool) {
	cart := GetUserCart(id)
	var orderamount uint
	for _, v := range cart {
		var order entity.Order
		uid, _ := strconv.ParseUint(id, 10, 64)
		pp, _ := strconv.ParseFloat(v.Product.Product.Price, 64)
		pdp, _ := strconv.ParseFloat(v.Product.Product.DiscountPrice, 64)

		order.ProductId = uint(v.Product.Product.ID)
		order.UserID = uint(uid)
		order.OrderStatus = 1
		order.Date = time.Now()
		order.Quantity = int(v.Count)
		if !v.Product.Product.DiscountStatus {
			order.Amount = float64(v.Count) * pdp
		}
		order.Amount = float64(v.Count) * pp
		orderamount = orderamount + uint(pp)
		DB.Create(&order)
	}
	if pm == "2" {
		client := razorpay.NewClient("rzp_test_0J3sQ8Oa3roCtI", "XKLK7IXr1bqXxuGnplscXY7J")
		fmt.Println(orderamount)
		params := map[string]interface{}{
			"amount":   orderamount,
			"currency": "INR",
			"receipt":  "some_receipt_id",
		}
		resp, err := client.Order.Create(params, nil)
		if err != nil {
			log.Fatal(err)
			return resp, false

		}
		return resp, true
	}
	return map[string]interface{}{}, true
}
