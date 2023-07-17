package adminController

import (
	"kwick/helper/database"
	images "kwick/helper/image"
	entity "kwick/model/entity"
	request "kwick/model/request"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllProducts(ctx *gin.Context) {
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
	products := database.GetProducts(int(pn), int(lt))
	var PageResponse []entity.ProductResponse
	var response entity.ProductResponse
	for _, v := range products {
		response.Product = v
		response.Images = database.GetImagesByProductId(v.ID)
		PageResponse = append(PageResponse, response)
	}
	numberOfProducts := database.GetTotalNumberOfProducts()
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
func GetProduct(ctx *gin.Context) {
	productId := ctx.Query("id")
	if len(productId) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Product id is missing in the request"})
		return
	}
	product, exist := database.GetProductById(productId)
	images := database.GetImagesByProductId(product.ID)

	if !exist {
		ctx.JSON(http.StatusNotFound, gin.H{"Message": "Product not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Images": images, "Product": product})
}
func AddProduct(ctx *gin.Context) {
	var data request.RequestBody
	cid, _ := strconv.ParseUint(ctx.PostForm("categoryId"), 10, 64)
	pq, _ := strconv.ParseUint(ctx.PostForm("productQuantiy"), 10, 64)
	data.CategoryId = uint(cid)
	data.ProductName = ctx.PostForm("productName")
	data.ProductDescription = ctx.PostForm("productDescription")
	data.ProductPrice, _ = strconv.ParseFloat(ctx.PostForm("productPrice"), 64)
	data.ProductQuantity = uint(pq)
	if data.CategoryId == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Category id should be valid"})
		return
	}
	if len(data.ProductName) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Product name is missing"})
		return
	}
	if len(data.ProductDescription) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Product description is missing"})
		return
	}
	if data.ProductPrice == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Product price should not be 0"})
		return
	}
	if data.ProductQuantity == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Can't add product without a single quantity"})
		return
	}
	added := database.AddProduct(data)
	if !added {
		ctx.JSON(http.StatusConflict, gin.H{"Message": "Product already exists"})
		return
	}
	productId := database.GetProductByName(data.ProductName)
	uploaded := images.ImageUploader(ctx, productId)
	if !uploaded {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not add product due to internal server issues"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "Product added successfully"})
}
func EditProduct(ctx *gin.Context) {
	id := ctx.Query("id")
	if len(id) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Product id is not available in request"})
		return
	}
	var data request.RequestBody
	cid, _ := strconv.ParseUint(ctx.PostForm("categoryId"), 10, 64)
	pq, _ := strconv.ParseUint(ctx.PostForm("productQuantiy"), 10, 64)
	data.CategoryId = uint(cid)
	data.ProductName = ctx.PostForm("productName")
	data.ProductDescription = ctx.PostForm("productDescription")
	data.ProductPrice, _ = strconv.ParseFloat(ctx.PostForm("productPrice"), 64)
	data.ProductQuantity = uint(pq)
	edited := database.EditProduct(id, data)
	if !edited {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not edit the product"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "Product edited"})
}
func BlockProduct(ctx *gin.Context) {
	id := ctx.Query("id")
	if len(id) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Product id is not available in request"})
		return
	}
	blocked := database.BlockProduct(id)
	if !blocked {
		ctx.JSON(http.StatusConflict, gin.H{"Message": "Product is already unavailable"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "Product set to unavailable"})
}
func UnblockProduct(ctx *gin.Context) {
	id := ctx.Query("id")
	if len(id) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Product id is not available in request"})
		return
	}
	unblocked := database.UnblockProduct(id)
	if !unblocked {
		ctx.JSON(http.StatusConflict, gin.H{"Message": "Product is already available"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "Product set to available"})
}
