package adminController

import (
	"kwick/helper/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(ctx *gin.Context) {
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
	users := database.GetUsers(int(pn), int(lt))
	numberOfUsers := database.GetTotalNumberOfUsers()
	if numberOfUsers > int(pn)*int(lt) && pn == 1 {
		ctx.JSON(http.StatusOK, gin.H{"Users": users, "Current Page": pn, "Previous Page": -1, "Next Page": pn + 1})
		return
	}
	if numberOfUsers <= int(pn)*int(lt) && pn == 1 {
		ctx.JSON(http.StatusOK, gin.H{"Users": users, "Current Page": pn, "Previous Page": -1, "Next Page": -1})
		return
	}
	if numberOfUsers <= int(pn)*int(lt) && pn > 1 {
		ctx.JSON(http.StatusOK, gin.H{"Users": users, "Current Page": pn, "Previous Page": pn - 1, "Next Page": -1})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Users": users, "Current Page": pn, "Previous Page": pn - 1, "Next Page": pn + 1})
}
func GetUser(ctx *gin.Context) {
	userId := ctx.Query("id")
	if len(userId) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "User id is missing in the request"})
		return
	}
	user, exist := database.GetUserById(userId)
	if !exist {
		ctx.JSON(http.StatusNotFound, gin.H{"Message": "User not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"User": user})
}
func BlockUser(ctx *gin.Context) {
	userId := ctx.Query("id")
	if len(userId) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "User id is missing in the request"})
		return
	}
	blocked, res := database.BlockUser(userId)
	if !blocked {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": res})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": res})
}
func UnblockUser(ctx *gin.Context) {
	userId := ctx.Query("id")
	if len(userId) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "User id is missing in the request"})
		return
	}
	blocked, res := database.UnblockUser(userId)
	if !blocked {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": res})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": res})
}
