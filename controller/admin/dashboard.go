package adminController

import (
	"kwick/helper/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDashboard(ctx *gin.Context) {
	users := database.GetTotalNumberOfUsers()
	vendors := database.GetTotalNumberOfVendors()
	products := database.GetTotalNumberOfProducts()
	categories := database.GetTotalNumberOfCategories()
	orders := database.GetTotalNumberOfOrders()
	sales := database.GetTotalNumberOfSales()
	revenue := database.GetTotalAmountofRevenue()
	ctx.JSON(http.StatusOK, gin.H{"Username": "Super Admin", "Number of users": users, "Number of vendors": vendors, "Number of products": products, "Number of categories": categories, "Number of orders": orders, "Number of sales": sales, "Total revenue": revenue})
}
