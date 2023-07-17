package database

import (
	"fmt"
	"kwick/helper/validation"
	entity "kwick/model/entity"
	request "kwick/model/request"

	"golang.org/x/crypto/bcrypt"
)

func GetTotalNumberOfUsers() int {
	var users int64
	if err := DB.Model(entity.User{}).Where("is_admin = ?", 0).Count(&users).Error; err != nil {
		return 0
	}
	return int(users)
}
func GetUserBySignUpData(e string, p string) (entity.User, bool) {
	var user entity.User
	DB.Where("email = ? OR phone = ?", e, p).Find(&user)
	if user.ID != 0 {
		return user, true
	}
	return entity.User{}, false
}
func GetUserById(i string) (entity.User, bool) {
	var user entity.User
	DB.Where("Id = ?", i).First(&user)
	return user, user.ID != 0
}
func GetUsers(p, l int) []entity.User {
	var user []entity.User
	offset := (int(p) - 1) * int(l)
	result := DB.Offset(offset).Limit(l).Find(&user)
	if result.Error != nil {
		panic("failed to query users: " + result.Error.Error())
	}
	return user
}
func BlockUser(i string) (bool, string) {
	var user entity.User
	DB.First(&user, i)
	if user.ID == 0 {
		return false, "User does not exist"
	}
	if user.IsAdmin {
		return false, "User is admin"
	}
	if user.Blocked {
		return false, "User is already blocked"
	}
	user.Blocked = true
	DB.Save(&user)
	return true, "User blocked"
}
func UnblockUser(i string) (bool, string) {
	var user entity.User
	DB.First(&user, i)
	if user.ID == 0 {
		return false, "User does not exist"
	}
	if user.IsAdmin {
		return false, "User is admin"
	}
	if !user.Blocked {
		return false, "User is already unblocked"
	}
	user.Blocked = false
	DB.Save(&user)
	return true, "User unblocked"
}
func AddUserToDb(u entity.User) bool {
	var user entity.User
	DB.Where("email = ? OR phone = ?", u.Email, u.Phone).First(&user)
	if user.ID != 0 {
		return false
	}
	hp, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	u.Password = string(hp)
	DB.Create(&u)
	return true
}
func GetUserFromDb(e string, p string) (entity.User, bool) {
	var user entity.User
	DB.Where("email = ? OR phone = ?", e, p).Find(&user)
	if user.ID != 0 {
		return user, true
	}
	return entity.User{}, false
}
func GetUserAddresses(i string) []entity.Address {
	var address []entity.Address
	DB.Where("user_id = ?", i).Find(&address)
	return address
}
func EditUser(i string, d request.RequestBody) bool {
	var user entity.User
	DB.Where("id = ?", i).Find(&user)
	if len(d.UserFirstName) != 0 || user.FirstName != d.UserFirstName {
		user.FirstName = d.UserFirstName
	}
	if len(d.UserMiddleName) != 0 || user.MiddleName != d.UserMiddleName {
		user.MiddleName = d.UserMiddleName
	}
	if len(d.UserLastName) != 0 || user.LastName != d.UserLastName {
		user.LastName = d.UserLastName
	}
	if len(d.UserEmail) != 0 || user.Email != d.UserEmail {
		res, valid := validation.ValidateEmail(d.UserEmail)
		fmt.Println(res)
		if !valid {
			return false
		}
		user.Email = d.UserEmail
	}
	DB.Save(&user)
	return true
}
func ChangePassword(id string, p string) bool {
	var user entity.User
	DB.Where("Id = ?", id).First(&user)
	if user.ID == 0 {
		return false
	}
	user.Password = p
	DB.Save(&user)
	return true
}
