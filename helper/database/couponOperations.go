package database

import (
	entity "kwick/model/entity"
)

func GetCoupons(p, l int) []entity.Coupon {
	var coupons []entity.Coupon
	offset := (int(p) - 1) * int(l)
	result := DB.Offset(offset).Limit(l).Find(&coupons)
	if result.Error != nil {
		panic("failed to query users: " + result.Error.Error())
	}
	return coupons
}
func GetTotalNumberOfCoupons() int {
	var coupons int64
	if err := DB.Model(entity.Coupon{}).Count(&coupons).Error; err != nil {
		return 0
	}
	return int(coupons)
}
func GetCouponById(i string) (entity.Coupon, bool) {
	var coupon entity.Coupon
	DB.Where("id = ?", i).Find(&coupon)
	if coupon.Id == 0 {
		return coupon, false
	}
	return coupon, true
}
func CreateCoupon(c entity.Coupon) bool {
	var exist entity.Coupon
	DB.Where("name = ?", c.Name).Find(&exist)
	if exist.Id != 0 {
		return false
	}
	var coupon entity.Coupon
	coupon.Name = c.Name
	coupon.Code = c.Code
	coupon.Limit = c.Limit
	coupon.AvailableFrom = c.AvailableFrom
	coupon.AvailableTill = c.AvailableTill
	coupon.Percentage = c.Percentage
	coupon.MinPurchase = c.MinPurchase
	coupon.MaxDiscount = c.MaxDiscount
	DB.Create(&coupon)
	return true
}
func EditCoupon(c entity.Coupon, id string) bool {
	var coupon entity.Coupon
	var flag int
	DB.Where("id = ?", id).Find(&coupon)
	if coupon.Id == 0 {
		return false
	}
	if coupon.Name == c.Name && coupon.Code == c.Code && coupon.Limit == c.Limit && coupon.AvailableFrom.Format("02 January 2006") == c.AvailableFrom.Format("02 January 2006") && coupon.AvailableTill.Format("02 January 2006") == c.AvailableTill.Format("02 January 2006") && coupon.Percentage == c.Percentage && coupon.MinPurchase == c.MinPurchase && coupon.MaxDiscount == c.MaxDiscount {
		flag++
	}
	if flag != 0 {
		return false
	}
	coupon.Name = c.Name
	coupon.Code = c.Code
	coupon.Limit = c.Limit
	coupon.AvailableFrom = c.AvailableFrom
	coupon.AvailableTill = c.AvailableTill
	coupon.Percentage = c.Percentage
	coupon.MinPurchase = c.MinPurchase
	coupon.MaxDiscount = c.MaxDiscount
	DB.Save(&coupon)
	return true
}
