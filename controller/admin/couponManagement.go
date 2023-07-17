package adminController

import (
	"kwick/helper/database"
	entity "kwick/model/entity"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetCoupons(ctx *gin.Context) {
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
	coupons := database.GetCoupons(int(pn), int(lt))
	numberOfCoupons := database.GetTotalNumberOfCoupons()
	if numberOfCoupons > int(pn)*int(lt) && pn == 1 {
		ctx.JSON(http.StatusOK, gin.H{"Coupons": coupons, "Current Page": pn, "Previous Page": -1, "Next Page": pn + 1})
		return
	}
	if numberOfCoupons <= int(pn)*int(lt) && pn == 1 {
		ctx.JSON(http.StatusOK, gin.H{"Coupons": coupons, "Current Page": pn, "Previous Page": -1, "Next Page": -1})
		return
	}
	if numberOfCoupons <= int(pn)*int(lt) && pn > 1 {
		ctx.JSON(http.StatusOK, gin.H{"Coupons": coupons, "Current Page": pn, "Previous Page": pn - 1, "Next Page": -1})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Coupons": coupons, "Current Page": pn, "Previous Page": pn - 1, "Next Page": pn + 1})
}
func GetCoupon(ctx *gin.Context) {
	couponId := ctx.Query("id")
	if len(couponId) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Coupon id is missing in the request"})
		return
	}
	coupon, exist := database.GetCouponById(couponId)
	if !exist {
		ctx.JSON(http.StatusNotFound, gin.H{"Message": "Coupon not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Coupon": coupon})
}
func CreateCoupon(ctx *gin.Context) {
	var coupon entity.Coupon
	cl, _ := strconv.ParseUint(ctx.PostForm("couponLimit"), 10, 64)
	cp, _ := strconv.ParseUint(ctx.PostForm("couponPercentage"), 10, 64)
	cmp, _ := strconv.ParseUint(ctx.PostForm("couponMinPurchase"), 10, 64)
	cmd, _ := strconv.ParseUint(ctx.PostForm("couponMaxDiscount"), 10, 64)
	af, _ := time.Parse("2006-01-02", ctx.PostForm("couponAvailableFrom"))
	at, _ := time.Parse("2006-01-02", ctx.PostForm("couponAvailableTill"))
	coupon.Name = ctx.PostForm("CouponName")
	coupon.Code = ctx.PostForm("CouponCode")
	coupon.Limit = uint(cl)
	coupon.AvailableFrom = af
	coupon.AvailableTill = at
	coupon.Percentage = uint(cp)
	coupon.MinPurchase = uint(cmp)
	coupon.MaxDiscount = uint(cmd)
	created := database.CreateCoupon(coupon)
	if !created {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not create coupon"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "Created coupon"})
}
func EditCoupon(ctx *gin.Context) {
	couponId := ctx.Query("id")
	if len(couponId) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Coupon id is missing in the request"})
		return
	}
	var coupon entity.Coupon
	cl, _ := strconv.ParseUint(ctx.PostForm("couponLimit"), 10, 64)
	cp, _ := strconv.ParseUint(ctx.PostForm("couponPercentage"), 10, 64)
	cmp, _ := strconv.ParseUint(ctx.PostForm("couponMinPurchase"), 10, 64)
	cmd, _ := strconv.ParseUint(ctx.PostForm("couponMaxDiscount"), 10, 64)
	af, _ := time.Parse("2006-01-02", ctx.PostForm("couponAvailableFrom"))
	at, _ := time.Parse("2006-01-02", ctx.PostForm("couponAvailableTill"))
	coupon.Name = ctx.PostForm("CouponName")
	coupon.Code = ctx.PostForm("CouponCode")
	coupon.Limit = uint(cl)
	coupon.AvailableFrom = af
	coupon.AvailableTill = at
	coupon.Percentage = uint(cp)
	coupon.MinPurchase = uint(cmp)
	coupon.MaxDiscount = uint(cmd)
	edited := database.EditCoupon(coupon, couponId)
	if !edited {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not edit coupon"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "Edited coupon"})

}
