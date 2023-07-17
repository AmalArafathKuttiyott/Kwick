package database

import (
	entity "kwick/model/entity"
	request "kwick/model/request"
	"strconv"
)

func CreateNewAddress(m request.RequestBody, id string) entity.Address {
	userId, _ := strconv.ParseUint(id, 10, 64)
	newAddress := entity.Address{
		UserId:         uint(userId),
		BuildingName:   m.UserBuildingName,
		BuildingNumber: m.UserBuildingNumber,
		City:           m.UserCity,
		Country:        m.UserCountry,
		State:          m.UserState,
		Street:         m.UserStreet,
		PostalCode:     m.UserPostalCode,
	}
	return newAddress
}
func AddAddressToDb(a entity.Address) bool {
	err := DB.Where("user_id = ? AND building_name = ? AND building_number = ? AND street = ? AND city = ? AND state = ? AND country = ? AND postal_code = ?", a.UserId, a.BuildingName, a.BuildingNumber, a.Street, a.City, a.State, a.Country, a.PostalCode).Find(&entity.Address{})
	if err.RowsAffected == 1 {
		return false
	}
	result := DB.Create(&a)
	return result.Error == nil
}
func DeleteAddressFromDb(id, uid string) bool {
	result := DB.Where("id = ? AND user_id = ?", id, uid).Delete(&entity.Address{})
	if result.Error != nil {
		return false
	}
	return result.RowsAffected > 0
}
func UpdateAddress(ra entity.Address, ai string) bool {
	var address entity.Address
	DB.Where("id = ?", ai).Find(&address)
	address.BuildingName = ra.BuildingName
	address.BuildingNumber = ra.BuildingNumber
	address.Street = ra.Street
	address.City = ra.City
	address.State = ra.State
	address.Country = ra.Country
	address.PostalCode = ra.PostalCode
	DB.Save(&address)
	return true
}
func GetUserAddress(ai string) entity.Address {
	var address entity.Address
	DB.Where("id = ?", ai).First(&address)
	return address
}
