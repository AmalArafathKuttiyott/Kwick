package userController

import (
	"fmt"
	"kwick/helper/bcrypt"
	"kwick/helper/database"
	"kwick/helper/jwt"
	twilio "kwick/helper/twilio"
	"kwick/helper/validation"
	entity "kwick/model/entity"
	request "kwick/model/request"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProfile(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	id := jwt.GetUserFromJwt(token)
	user, exist := database.GetUserById(id)
	if !exist {
		ctx.JSON(http.StatusNotFound, gin.H{"Message": "User not found"})
		return
	}
	address := database.GetUserAddresses(id)
	profile := entity.UserProfile{
		FirstName:    user.FirstName,
		MiddleName:   user.MiddleName,
		LastName:     user.LastName,
		Email:        user.Email,
		Phone:        user.Phone,
		Wallet:       user.Wallet,
		ReferralCode: user.ReferralCode,
		Address:      address,
	}
	ctx.JSON(http.StatusOK, gin.H{"User Details": profile})
}
func EditProfile(ctx *gin.Context) {
	var data request.RequestBody
	data.UserFirstName = ctx.PostForm("userFirstname")
	data.UserMiddleName = ctx.PostForm("userMiddlename")
	data.UserLastName = ctx.PostForm("userLastname")
	data.UserEmail = ctx.PostForm("userEmail")
	token := ctx.GetHeader("Authorization")
	id := jwt.GetUserFromJwt(token)
	edited := database.EditUser(id, data)
	if !edited {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not edit user profile"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "User profile updated"})
}
func AddAddress(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	id := jwt.GetUserFromJwt(token)
	var data request.RequestBody
	data.UserBuildingName = ctx.PostForm("userBuildingName")
	bn, _ := strconv.ParseInt(ctx.PostForm("userBuildingNumber"), 10, 64)
	data.UserBuildingNumber = int(bn)
	data.UserCity = ctx.PostForm("userCity")
	data.UserStreet = ctx.PostForm("userStreet")
	data.UserState = ctx.PostForm("userState")
	data.UserCountry = ctx.PostForm("userCountry")
	data.UserPostalCode = ctx.PostForm("userPostalCode")
	address := database.CreateNewAddress(data, id)
	res := database.AddAddressToDb(address)
	if !res {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not add address"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "Added new address"})
}
func DeleteAddress(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	id := jwt.GetUserFromJwt(token)
	addressId := ctx.Query("address-id")
	deleted := database.DeleteAddressFromDb(addressId, id)
	if !deleted {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not delete address"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "Deleted address"})
}
func EditAddress(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	id := jwt.GetUserFromJwt(authHeader)
	addressId := ctx.Query("address-id")
	userId, _ := strconv.ParseUint(id, 10, 64)
	adId, _ := strconv.ParseUint(addressId, 10, 64)
	bn, _ := strconv.ParseInt(ctx.PostForm("userBuildingNumber"), 10, 64)
	ra := entity.Address{
		Id:             uint(adId),
		UserId:         uint(userId),
		BuildingName:   ctx.PostForm("userBuildingName"),
		BuildingNumber: int(bn),
		Street:         ctx.PostForm("userStreet"),
		City:           ctx.PostForm("userCity"),
		State:          ctx.PostForm("userState"),
		Country:        ctx.PostForm("userCountry"),
		PostalCode:     ctx.PostForm("userPostalCode"),
	}
	existingAddress := database.GetUserAddress(addressId)
	adid := fmt.Sprintf("%d", existingAddress.UserId)
	if id != adid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"Message": "This is not your address"})
		return
	}
	same := validation.CompareAddresses(ra, existingAddress)
	if !same {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "No changes have been made, change atleast 1 field to update your address"})
		return
	}
	database.UpdateAddress(ra, id)
	ctx.JSON(http.StatusOK, gin.H{"Message": "Updated address"})
}
func ForgotPassword(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	id := jwt.GetUserFromJwt(token)
	user, exist := database.GetUserById(id)
	if !exist {
		ctx.JSON(http.StatusNotFound, gin.H{"Message": "User not found"})
		return
	}
	sent := twilio.SendOtp(user.Phone)
	if !sent {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not send otp"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "Sent otp to user phone number"})
}
func ForgotPasswordOtpVerification(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	id := jwt.GetUserFromJwt(token)
	user, exist := database.GetUserById(id)
	if !exist {
		ctx.JSON(http.StatusNotFound, gin.H{"Message": "User not found"})
		return
	}
	otp := ctx.PostForm("otp")
	if otp == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Otp not found"})
		return
	}
	verified := twilio.VerifyOtp(user.Phone, otp)
	if !verified {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Otp verification failed, enter the otp you got"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "Otp verified, go to /homepage/profile/editprofile/changepassword"})
}
func ChangePassword(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	id := jwt.GetUserFromJwt(token)
	data := request.RequestBody{
		UserNewPassword:     ctx.PostForm("userNewPassword"),
		UserConfirmPassword: ctx.PostForm("userConfirmPassword"),
	}
	if data.UserNewPassword == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "New password should not be blank"})
		return
	}
	if data.UserConfirmPassword == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "New password confirmation should not be blank"})
		return
	}
	if data.UserNewPassword != data.UserConfirmPassword {
		ctx.JSON(http.StatusBadRequest, gin.H{"MEssage": "Password does not match"})
		return
	}
	valid := validation.ValidatePassword(data.UserNewPassword)
	if !valid {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Enter a strong password"})
		return
	}
	password := bcrypt.HashPassword(data.UserNewPassword)
	changed := database.ChangePassword(id, password)
	if !changed {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not change the password"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "Changed Password"})
}
