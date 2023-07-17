package userController

import (
	"kwick/helper/database"
	jwt "kwick/helper/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHomepage(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	id := jwt.GetUserFromJwt(token)
	if len(id) == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"Message": "User id not found in token"})
		return
	}
	user, exist := database.GetUserById(id)
	if !exist {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "User does not exist"})
		return
	}
	if user.Blocked {
		ctx.JSON(http.StatusUnauthorized, gin.H{"Message": "Your account have been suspended"})
		return
	}
	products := database.GetProducts(1, 20)
	categories := database.GetCategories(1, 10)
	ctx.JSON(http.StatusOK, gin.H{"User": user.FirstName, "Products": products, "Categories": categories})
}
