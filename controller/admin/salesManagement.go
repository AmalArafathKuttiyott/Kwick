package adminController

import (
	"fmt"
	"kwick/helper/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
)

func DownloadSalesPdf(ctx *gin.Context) {
	salesAmount := database.GetTotalAmountofRevenue()
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, fmt.Sprintf("Total Sales: $%.2v", salesAmount))
	err := pdf.OutputFileAndClose("public/pdfs/sales_report.pdf")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not download pdf"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "Downloaded pdf"})
}
