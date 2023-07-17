package adminController

import (
	"kwick/helper/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetOrders(ctx *gin.Context) {
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
	orders := database.GetOrders(int(pn), int(lt))
	numberOfOrders := database.GetTotalNumberOfOrders()
	if numberOfOrders > int(pn)*int(lt) && pn == 1 {
		ctx.JSON(http.StatusOK, gin.H{"Orders": orders, "Current Page": pn, "Previous Page": -1, "Next Page": pn + 1})
		return
	}
	if numberOfOrders <= int(pn)*int(lt) && pn == 1 {
		ctx.JSON(http.StatusOK, gin.H{"Orders": orders, "Current Page": pn, "Previous Page": -1, "Next Page": -1})
		return
	}
	if numberOfOrders <= int(pn)*int(lt) && pn > 1 {
		ctx.JSON(http.StatusOK, gin.H{"Orders": orders, "Current Page": pn, "Previous Page": pn - 1, "Next Page": -1})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Orders": orders, "Current Page": pn, "Previous Page": pn - 1, "Next Page": pn + 1})
}
func GetDeliveredOrders(ctx *gin.Context) {
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
	orders := database.GetDeliveredOrders(int(pn), int(lt))
	numberOfOrders := database.GetTotalNumberOfDeliveredOrders()
	if numberOfOrders > int(pn)*int(lt) && pn == 1 {
		ctx.JSON(http.StatusOK, gin.H{"Orders": orders, "Current Page": pn, "Previous Page": -1, "Next Page": pn + 1})
		return
	}
	if numberOfOrders <= int(pn)*int(lt) && pn == 1 {
		ctx.JSON(http.StatusOK, gin.H{"Orders": orders, "Current Page": pn, "Previous Page": -1, "Next Page": -1})
		return
	}
	if numberOfOrders <= int(pn)*int(lt) && pn > 1 {
		ctx.JSON(http.StatusOK, gin.H{"Orders": orders, "Current Page": pn, "Previous Page": pn - 1, "Next Page": -1})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Orders": orders, "Current Page": pn, "Previous Page": pn - 1, "Next Page": pn + 1})
}
func GetReturnOrders(ctx *gin.Context) {
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
	orders := database.GetReturnOrders(int(pn), int(lt))
	numberOfOrders := database.GetTotalNumberOfReturnOrders()
	if numberOfOrders > int(pn)*int(lt) && pn == 1 {
		ctx.JSON(http.StatusOK, gin.H{"Orders": orders, "Current Page": pn, "Previous Page": -1, "Next Page": pn + 1})
		return
	}
	if numberOfOrders <= int(pn)*int(lt) && pn == 1 {
		ctx.JSON(http.StatusOK, gin.H{"Orders": orders, "Current Page": pn, "Previous Page": -1, "Next Page": -1})
		return
	}
	if numberOfOrders <= int(pn)*int(lt) && pn > 1 {
		ctx.JSON(http.StatusOK, gin.H{"Orders": orders, "Current Page": pn, "Previous Page": pn - 1, "Next Page": -1})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Orders": orders, "Current Page": pn, "Previous Page": pn - 1, "Next Page": pn + 1})
}
func GetPendingOrders(ctx *gin.Context) {
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
	orders := database.GetPendingOrders(int(pn), int(lt))
	numberOfOrders := database.GetTotalNumberOfPendingOrders()
	if numberOfOrders > int(pn)*int(lt) && pn == 1 {
		ctx.JSON(http.StatusOK, gin.H{"Orders": orders, "Current Page": pn, "Previous Page": -1, "Next Page": pn + 1})
		return
	}
	if numberOfOrders <= int(pn)*int(lt) && pn == 1 {
		ctx.JSON(http.StatusOK, gin.H{"Orders": orders, "Current Page": pn, "Previous Page": -1, "Next Page": -1})
		return
	}
	if numberOfOrders <= int(pn)*int(lt) && pn > 1 {
		ctx.JSON(http.StatusOK, gin.H{"Orders": orders, "Current Page": pn, "Previous Page": pn - 1, "Next Page": -1})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Orders": orders, "Current Page": pn, "Previous Page": pn - 1, "Next Page": pn + 1})
}
func GetOrder(ctx *gin.Context) {
	orderId := ctx.Query("id")
	if len(orderId) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Order id is missing in the request"})
		return
	}
	order, exist := database.GetOrderById(orderId)
	if !exist {
		ctx.JSON(http.StatusNotFound, gin.H{"Message": "Order not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Order": order})
}
func ChangeOrderStatus(ctx *gin.Context) {
	orderId := ctx.Query("id")
	if len(orderId) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Order id is missing in the request"})
		return
	}
	statusId := ctx.Query("status-id")
	if len(orderId) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Status id is missing in the request"})
		return
	}
	edited := database.EditOrder(orderId, statusId)
	if !edited {
		ctx.JSON(http.StatusNotFound, gin.H{"Message": "Order status could not be changed"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "Order Status changed"})
}
