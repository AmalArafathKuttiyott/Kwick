package adminController

import (
	"kwick/helper/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllVendors(ctx *gin.Context) {
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
	vendors := database.GetVendors(int(pn), int(lt))
	numberOfVendors := database.GetTotalNumberOfVendors()
	if numberOfVendors > int(pn)*int(lt) && pn == 1 {
		ctx.JSON(http.StatusOK, gin.H{"Vendors": vendors, "Current Page": pn, "Previous Page": -1, "Next Page": pn + 1})
		return
	}
	if numberOfVendors <= int(pn)*int(lt) && pn == 1 {
		ctx.JSON(http.StatusOK, gin.H{"Vendors": vendors, "Current Page": pn, "Previous Page": -1, "Next Page": -1})
		return
	}
	if numberOfVendors <= int(pn)*int(lt) && pn > 1 {
		ctx.JSON(http.StatusOK, gin.H{"Vendors": vendors, "Current Page": pn, "Previous Page": pn - 1, "Next Page": -1})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Vendors": vendors, "Current Page": pn, "Previous Page": pn - 1, "Next Page": pn + 1})
}
func GetVendor(ctx *gin.Context) {
	vendorId := ctx.Query("id")
	if len(vendorId) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Vendor id is missing in the request"})
		return
	}
	vendor, exist := database.GetVendorById(vendorId)
	if !exist {
		ctx.JSON(http.StatusNotFound, gin.H{"Message": "Vendor not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Vendor": vendor})
}
func BlockVendor(ctx *gin.Context) {
	vendorId := ctx.Query("id")
	if len(vendorId) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Vendor id is missing in the request"})
		return
	}
	blocked, res := database.BlockVendor(vendorId)
	if !blocked {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": res})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": res})
}
func UnblockVendor(ctx *gin.Context) {
	vendorId := ctx.Query("id")
	if len(vendorId) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Vendor id is missing in the request"})
		return
	}
	blocked, res := database.UnblockVendor(vendorId)
	if !blocked {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": res})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": res})
}
func GetVendorRequests(ctx *gin.Context) {
	pageNumber := ctx.Query("page-number")
	if len(pageNumber) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Page number is missing in the request"})
		return
	}
	limit := ctx.Query("limit")
	if len(limit) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "limit number is missing in the request"})
		return
	}
	pn, _ := strconv.ParseInt(pageNumber, 10, 64)
	lt, _ := strconv.ParseInt(limit, 10, 64)
	vendors := database.GetVendorRequests(int(pn), int(lt))
	numberOfRequests := database.GetTotalNumberOfVendors()
	if numberOfRequests > int(pn)*int(lt) && pn == 1 {
		ctx.JSON(http.StatusOK, gin.H{"Users": vendors, "Current Page": pn, "Previous Page": -1, "Next Page": pn + 1})
		return
	}
	if numberOfRequests <= int(pn)*int(lt) && pn == 1 {
		ctx.JSON(http.StatusOK, gin.H{"Users": vendors, "Current Page": pn, "Previous Page": -1, "Next Page": -1})
		return
	}
	if numberOfRequests <= int(pn)*int(lt) && pn > 1 {
		ctx.JSON(http.StatusOK, gin.H{"Users": vendors, "Current Page": pn, "Previous Page": pn - 1, "Next Page": -1})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Users": vendors, "Current Page": pn, "Previous Page": pn - 1, "Next Page": pn + 1})
}
func AcceptVendorship(ctx *gin.Context) {
	requestId := ctx.Query("req-id")
	if len(requestId) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Vendor id is missing in the request"})
		return
	}
	accepted := database.AcceptVedorship(requestId)
	if !accepted {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not accept vendorship request"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "Vendorship request have been accepted"})
}
func RejectVendorship(ctx *gin.Context) {
	requestId := ctx.Query("req-id")
	if len(requestId) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Vendor id is missing in the request"})
		return
	}
	rejected := database.RejectVedorship(requestId)
	if !rejected {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not reject vendorship request"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "Vendorship request have been rejected"})
}
