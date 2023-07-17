package userController

import (
	database "kwick/helper/database"
	jwt "kwick/helper/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrders(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	id := jwt.GetUserFromJwt(token)
	orders := database.GetUserOrders(id)
	if orders == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "There are no active orders"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Orders": orders})
}
func CancelOrder(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	id := jwt.GetUserFromJwt(token)
	ordersId := ctx.Query("order-id")
	status, code := database.CancelOrder(ordersId, id)
	if !status {
		if code == 1 {
			ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Could not find order , check if order is valid"})
			return
		}
		if code == 2 {
			ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Order has been already shipped, you can't cancel the order right now"})
			return
		}

	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "Cancelled order"})
}
func ReturnOrder(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	id := jwt.GetUserFromJwt(authHeader)
	ordersId := ctx.Query("order-id")
	returned := database.ReturnOrder(ordersId, id)
	if !returned {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Can't return the order at the moment, check if you have recieved the order or whether the order has been cancelled"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "Return request have been recieved, wait for the admins to review it"})
}
