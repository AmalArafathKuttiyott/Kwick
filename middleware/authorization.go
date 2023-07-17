package middleware

import (
	"fmt"
	"kwick/helper/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminAuthorization(ctx *gin.Context) {
	auth := ctx.GetHeader("authorization")
	if auth == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Authentication failed, token missing."})
		return
	}
	token := jwt.ExtractBearerToken(auth)
	exist := jwt.AdminParseJwt(token)
	if !exist {
		ctx.JSON(http.StatusNotFound, gin.H{"Message": "User mentioned in token not found in the database or user might not be an admin."})
		return
	} else {
		ctx.Next()
	}
}
func UserAuthorization(ctx *gin.Context) {
	auth := ctx.GetHeader("authorization")
	if auth == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Authentication failed, token missing."})
		return
	}
	token := jwt.ExtractBearerToken(auth)
	exist := jwt.ParseJwt(token)
	if !exist {
		ctx.JSON(http.StatusNotFound, gin.H{"Message": "User mentioned in token not found in the database."})
		return
	} else {
		ctx.Next()
	}
}
func VendorAuthorization(ctx *gin.Context) {
	fmt.Println("Entered Vendor Authorization")
	ctx.Next()
}
