package adminController

import (
	"kwick/helper/database"
	request "kwick/model/request"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllCategories(ctx *gin.Context) {
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
	categories := database.GetCategories(int(pn), int(lt))
	numberOfCategories := database.GetTotalNumberOfCategories()
	if numberOfCategories > int(pn)*int(lt) && pn == 1 {
		ctx.JSON(http.StatusOK, gin.H{"Categories": categories, "Current Page": pn, "Previous Page": -1, "Next Page": pn + 1})
		return
	}
	if numberOfCategories <= int(pn)*int(lt) && pn == 1 {
		ctx.JSON(http.StatusOK, gin.H{"Categories": categories, "Current Page": pn, "Previous Page": -1, "Next Page": -1})
		return
	}
	if numberOfCategories <= int(pn)*int(lt) && pn > 1 {
		ctx.JSON(http.StatusOK, gin.H{"Categories": categories, "Current Page": pn, "Previous Page": pn - 1, "Next Page": -1})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Categories": categories, "Current Page": pn, "Previous Page": pn - 1, "Next Page": pn + 1})
}
func GetCategory(ctx *gin.Context) {
	categoryId := ctx.Query("id")
	if len(categoryId) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Category id is missing in the request"})
		return
	}
	category, exist := database.GetCategoryById(categoryId)
	if !exist {
		ctx.JSON(http.StatusNotFound, gin.H{"Message": "Category not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Category": category})
}
func AddCategory(ctx *gin.Context) {
	var data request.RequestBody
	ctx.ShouldBind(&data)
	if len(data.CategoryName) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Category name is missing"})
		return
	}
	added := database.AddCategory(data.CategoryName)
	if !added {
		ctx.JSON(http.StatusConflict, gin.H{"Message": "Category already exists"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "Category added successfully"})
}
func BlockCategory(ctx *gin.Context) {
	categoryId := ctx.Query("id")
	if len(categoryId) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Category id is missing in the request"})
		return
	}
	blocked, res := database.BlockCategory(categoryId)
	if !blocked {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": res})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": res})
}
func UnblockCategory(ctx *gin.Context) {
	categoryId := ctx.Query("id")
	if len(categoryId) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Category id is missing in the request"})
		return
	}
	blocked, res := database.UnblockCategory(categoryId)
	if !blocked {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": res})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": res})
}
func EditCategory(ctx *gin.Context) {
	var data request.RequestBody
	ctx.ShouldBind(&data)
	categoryId := ctx.Query("id")
	edited := database.EditCategory(categoryId, data.CategoryName)
	if !edited {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Could not edit category"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "Edited category"})
}
