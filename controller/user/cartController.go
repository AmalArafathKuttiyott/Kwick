package userController

import (
	"fmt"
	"kwick/helper/database"
	jwt "kwick/helper/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReduceProductCount(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	id := jwt.GetUserFromJwt(token)
	ci := ctx.Query("cart-id")
	pi := ctx.Query("product-id")
	res := database.ReduceCount(ci, pi, id)
	if !res {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not reduce product count"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "Reduced product count"})
}
func IncreaseProductcount(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	id := jwt.GetUserFromJwt(token)
	ci := ctx.Query("cart-id")
	pi := ctx.Query("product-id")
	res := database.IncreaseCount(ci, pi, id)
	if !res {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not increase product count"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "Increased product count"})
}
func RemoveFromCart(ctx *gin.Context) {
	ci := ctx.Query("cart-id")
	pi := ctx.Query("product-id")
	res := database.RemoveFromCart(ci, pi)
	if !res {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not increase product count"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "Increased product count"})
}
func GetUserCart(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	id := jwt.GetUserFromJwt(token)
	cart := database.GetUserCart(id)
	if len(cart) == 0 {
		ctx.JSON(http.StatusOK, gin.H{"Message": "User does not have anything in their cart"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Cart": cart})
}
func AddToCart(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	id := jwt.GetUserFromJwt(authHeader)
	pi := ctx.Query("id")
	added := database.AddProductToCart(pi, id)
	if !added {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not add product to cart"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "Added to cart"})
}
func AddToWishlist(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	id := jwt.GetUserFromJwt(authHeader)
	pi := ctx.Query("product-id")
	added := database.AddProductToWishlist(pi, id)
	if !added {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not add to wishlist"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "Added to wishlist"})
}
func RemoveFromWishlist(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	id := jwt.GetUserFromJwt(authHeader)
	pi := ctx.Query("product-id")
	added := database.RemoveFromWishlist(pi, id)
	if !added {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not remove from wishlist"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "Removed from wishlist"})
}
func Checkout(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	id := jwt.GetUserFromJwt(token)
	ai := ctx.Query("address_Id")
	pm := ctx.Query("payment_Id")
	resp, ordered := database.AddOrder(id, ai, pm)
	if !ordered {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Could not checkout"})
		return
	}
	cart := database.GetUserCart(id)
	for _, v := range cart {
		pid := fmt.Sprintf("%d", v.Product.Product.ID)
		cid := fmt.Sprintf("%d", v.CartId)
		database.RemoveFromCart(cid, pid)
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "Order has been placed", "Order Details": resp})
}
