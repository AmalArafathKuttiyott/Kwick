package userController

import (
	"kwick/helper/database"
	entity "kwick/model/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCategory(ctx *gin.Context) {
	categoryId := ctx.Query("id")
	if len(categoryId) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Category id is missing in the request"})
		return
	}
	pageNumber := ctx.Query("page-number")
	if len(pageNumber) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Page number is missing in the request"})
		return
	}
	limit := ctx.Query("limit")
	if len(limit) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Limit number is missing in the request"})
		return
	}
	pn, _ := strconv.ParseInt(pageNumber, 10, 64)
	lt, _ := strconv.ParseInt(limit, 10, 64)
	products := database.GetProductsByCategory(int(pn), int(lt), categoryId)
	var PageResponse []entity.ProductResponse
	var response entity.ProductResponse
	for _, v := range products {
		response.Product = v
		response.Images = database.GetImagesByProductId(v.ID)
		PageResponse = append(PageResponse, response)
	}
	numberOfProducts := database.GetTotalCountOfProductsByCategory(categoryId)
	if numberOfProducts > int(pn)*int(lt) && pn == 1 {
		ctx.JSON(http.StatusOK, gin.H{"Products": PageResponse, "Current Page": pn, "Previous Page": -1, "Next Page": pn + 1})
		return
	}
	if numberOfProducts <= int(pn)*int(lt) && pn == 1 {
		ctx.JSON(http.StatusOK, gin.H{"Products": PageResponse, "Current Page": pn, "Previous Page": -1, "Next Page": -1})
		return
	}
	if numberOfProducts <= int(pn)*int(lt) && pn > 1 {
		ctx.JSON(http.StatusOK, gin.H{"Products": PageResponse, "Current Page": pn, "Previous Page": pn - 1, "Next Page": -1})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Products": PageResponse, "Current Page": pn, "Previous Page": pn - 1, "Next Page": pn + 1})
}
