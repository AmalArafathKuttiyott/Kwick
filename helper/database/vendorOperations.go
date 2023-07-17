package database

import (
	entity "kwick/model/entity"
)

func GetTotalNumberOfVendors() int {
	var vendors int64
	if err := DB.Model(entity.User{}).Where("is_vendor = ?", 1).Count(&vendors).Error; err != nil {
		return 0
	}
	return int(vendors)
}
func GetVendors(p, l int) []entity.User {
	var user []entity.User
	offset := (int(p) - 1) * int(l)
	result := DB.Where("is_vendor = ?", true).Offset(offset).Limit(l).Find(&user)
	if result.Error != nil {
		panic("failed to query users: " + result.Error.Error())
	}
	return user
}
func GetVendorById(i string) (entity.User, bool) {
	var user entity.User
	DB.Where("Id = ? AND is_vendor = ?", i, true).First(&user)
	return user, user.ID != 0
}
func BlockVendor(i string) (bool, string) {
	var user entity.User
	DB.Where("id = ? AND is_vendor = ?", i, true).First(&user)
	if user.ID == 0 {
		return false, "Vendor does not exist"
	}
	if user.IsAdmin {
		return false, "User is admin"
	}
	if user.Blocked {
		return false, "Vendor is already blocked"
	}
	user.Blocked = true
	DB.Save(&user)
	return true, "Vendor blocked"
}
func UnblockVendor(i string) (bool, string) {
	var user entity.User
	DB.Where("id = ? AND is_vendor = ?", i, true).First(&user)
	if user.ID == 0 {
		return false, "Vendor does not exist"
	}
	if user.IsAdmin {
		return false, "User is admin"
	}
	if !user.Blocked {
		return false, "Vendor is already unblocked"
	}
	user.Blocked = false
	DB.Save(&user)
	return true, "Vendor unblocked"
}
func GetVendorRequests(p, l int) []entity.Business {
	var vendors []entity.Business
	offset := (int(p) - 1) * int(l)
	result := DB.Where("verified = ? AND rejected = ?", false, false).Offset(offset).Limit(l).Find(&vendors)
	if result.Error != nil {
		panic("failed to query users: " + result.Error.Error())
	}
	return vendors
}
func GetTotalNumberOfVendorRequests() int {
	var vendors int64
	if err := DB.Model(entity.Business{}).Where("verified = ?", false).Count(&vendors).Error; err != nil {
		return 0
	}
	return int(vendors)
}
func AcceptVedorship(i string) bool {
	var request entity.Business
	DB.Where("id = ?", i).First(&request)
	if request.Id == 0 {
		return false
	}
	request.Verified = true
	DB.Save(&request)
	return true
}
func RejectVedorship(i string) bool {
	var request entity.Business
	DB.Where("id = ?", i).First(&request)
	if request.Id == 0 {
		return false
	}
	request.Verified = false
	request.Rejected = true
	DB.Save(&request)
	return true
}
