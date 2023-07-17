package userController

import (
	"net/http"

	"kwick/helper/database"
	twilio "kwick/helper/twilio"
	"kwick/helper/validation"
	entity "kwick/model/entity"
	request "kwick/model/request"

	"github.com/gin-gonic/gin"
)

var Details entity.User

func UserSignup(ctx *gin.Context) {
	var requestBody request.RequestBody
	if err := ctx.ShouldBind(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Trouble getting signup details. Try again !"})
		return
	}
	if response, ok := validation.SignupFormCheck(requestBody); !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": response})
		return
	}
	if response, ok := validation.ValidateSignupDetails(requestBody); !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": response})
		return
	}
	if _, exists := database.GetUserBySignUpData(requestBody.UserEmail, requestBody.UserPhone); exists {
		ctx.JSON(http.StatusConflict, gin.H{"Error": "User already exists"})
		return
	}
	if send := twilio.SendOtp(requestBody.UserPhone); !send {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not send otp to the phone number. Try again"})
		return
	}
	Details = entity.User{
		FirstName:    requestBody.UserFirstName,
		MiddleName:   requestBody.UserMiddleName,
		LastName:     requestBody.UserLastName,
		Email:        requestBody.UserEmail,
		Phone:        requestBody.UserPhone,
		Password:     requestBody.UserPassword,
		ReferralCode: requestBody.ReferralCode,
	}
	ctx.JSON(http.StatusOK, gin.H{"Status": "Please verify otp sent to your phone number at :'http://localhost:8080/signup-otpverification'"})
}
func SignupOtpVerification(ctx *gin.Context) {
	var requestBody request.RequestBody
	if err := ctx.ShouldBind(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Trouble getting otp. Try again !"})
		return
	}
	if verified := twilio.VerifyOtp(Details.Phone, requestBody.Otp); !verified {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Otp verification failed"})
		return
	}
	if added := database.AddUserToDb(Details); !added {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "User already exists ! Try signing-in"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "User signed up successfully"})
}
