package adminController

import (
	"kwick/helper/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetOffers(ctx *gin.Context) {
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
	offers := database.GetOffers(int(pn), int(lt))
	numberOfOffers := database.GetTotalNumberOfOffers()
	if numberOfOffers > int(pn)*int(lt) && pn == 1 {
		ctx.JSON(http.StatusOK, gin.H{"Offers": offers, "Current Page": pn, "Previous Page": -1, "Next Page": pn + 1})
		return
	}
	if numberOfOffers <= int(pn)*int(lt) && pn == 1 {
		ctx.JSON(http.StatusOK, gin.H{"Offers": offers, "Current Page": pn, "Previous Page": -1, "Next Page": -1})
		return
	}
	if numberOfOffers <= int(pn)*int(lt) && pn > 1 {
		ctx.JSON(http.StatusOK, gin.H{"Offers": offers, "Current Page": pn, "Previous Page": pn - 1, "Next Page": -1})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Offers": offers, "Current Page": pn, "Previous Page": pn - 1, "Next Page": pn + 1})
}
func GetOffer(ctx *gin.Context) {
	productId := ctx.Query("id")
	if len(productId) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Product id is missing in the request"})
		return
	}
	product, exist := database.GetProductById(productId)
	if !exist {
		ctx.JSON(http.StatusNotFound, gin.H{"Message": "Product not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Product": product})
}
func AddOffer(ctx *gin.Context) {
	productId := ctx.Query("id")
	if len(productId) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Product id is missing in the request"})
		return
	}
	discountPrice := ctx.PostForm("discountPrice")
	addOffer := database.AddProductOfferById(productId, discountPrice)
	if !addOffer {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Could not add discount to the product"})
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Added discount to the product"})
}
func AddCategoryOffer(ctx *gin.Context) {
	categoryId := ctx.Query("id")
	if len(categoryId) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Category id is missing in the request"})
		return
	}
	discountPrice := ctx.PostForm("discountPrice")
	addOffer := database.AddProductOfferById(categoryId, discountPrice)
	if !addOffer {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Could not add discount to the category"})
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Added discount to the category"})
}
func RemoveOffer(ctx *gin.Context) {
	productId := ctx.Query("id")
	if len(productId) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Product id is missing in the request"})
		return
	}
	removedOffer := database.RemoveProductOfferById(productId)
	if !removedOffer {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Could not remove discount of the product"})
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Removed discount of the product"})
}
func RemoveCategoryOffer(ctx *gin.Context) {
	categoryId := ctx.Query("id")
	if len(categoryId) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Category id is missing in the request"})
		return
	}
	addOffer := database.RemoveCategoryOfferById(categoryId)
	if !addOffer {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Could not remove discount of the category"})
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Removed discount of the category"})
}
func EditOffer(ctx *gin.Context) {
	productId := ctx.Query("id")
	if len(productId) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Category id is missing in the request"})
		return
	}
	discountPrice := ctx.PostForm("discountPrice")
	edited := database.EditProductOfferById(productId, discountPrice)
	if !edited {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Could not edit discount of the product"})
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Edited discount to the product"})
}
