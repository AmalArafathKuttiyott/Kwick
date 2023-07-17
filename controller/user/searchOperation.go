package userController

import (
	"kwick/helper/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSearchResult(ctx *gin.Context) {
	category := ctx.Query("category")
	keyword := ctx.Query("keyword")
	products := database.GetProductBySearch(category, keyword)
	ctx.JSON(http.StatusOK, gin.H{"Products": products})
}
