package userController

import (
	bcrypt "kwick/helper/bcrypt"
	"kwick/helper/database"
	"kwick/helper/jwt"
	"kwick/helper/validation"
	request "kwick/model/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserSignin(ctx *gin.Context) {
	var requestBody request.RequestBody
	if err := ctx.ShouldBind(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Trouble getting signin details. Try again !"})
		return
	}
	if status, check := validation.SigninFormCheck(requestBody); !check {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": status})
		return
	}
	if status, check := validation.ValidateSigninDetails(requestBody); !check {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": status})
		return
	}
	user, exist := database.GetUserFromDb(requestBody.UserSigninData, requestBody.UserSigninData)
	if !exist {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "User does not exist, try signing in"})
		return
	}
	if valid := bcrypt.ComparePassword(user.Password, requestBody.UserPassword); !valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"Message": "Wrong password, try again"})
		return
	}
	token, status := jwt.GeneratJwt(user.ID)
	if !status {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not generate authentication token"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "Sign-in successful", "Token": token, "User": user.FirstName + " " + user.MiddleName + " " + user.LastName})
}
