package models

import "time"

type RequestBody struct {
	// Token details
	Token string `json:"token"`
	// Product details
	ProductId             uint    `json:"productId"`
	ProductName           string  `json:"productName"`
	ProductPrice          float64 `json:"productPrice"`
	ProductDiscountPrice  float64 `json:"productDiscountPrice"`
	ProductDiscountStatus bool    `json:"productDiscountStatus"`
	ProductDescription    string  `json:"productDescription"`
	ProductAvailability   bool    `json:"productAvailability"`
	ProductQuantity       uint    `json:"productQuantiy"`
	// Category details
	CategoryId      uint   `json:"categoryId"`
	CategoryName    string `json:"categoryName"`
	CategoryBlocked bool   `json:"categoryBlocked"`
	// User details
	UserId              uint
	UserFirstName       string `json:"userFirstname"`
	UserMiddleName      string `json:"userMiddlename"`
	UserLastName        string `json:"userLastname"`
	UserEmail           string `json:"userEmail"`
	UserPhone           string `json:"userPhone"`
	UserPassword        string `json:"userPassword"`
	UserConfirmPassword string `json:"userConfirmPassword"`
	UserBlocked         bool   `json:"userBlocked"`
	UserIsAdmin         bool   `json:"userIsAdmin"`
	UserBuildingName    string `json:"userBuildingName"`
	UserBuildingNumber  int    `json:"userBuildingNumber"`
	UserStreet          string `json:"userStreet"`
	UserCity            string `json:"userCity"`
	UserState           string `json:"userState"`
	UserCountry         string `json:"userCountry"`
	UserPostalCode      string `json:"userPostalCode"`
	UserSigninData      string `json:"userSignindata"`
	Otp                 string `json:"otp"`
	UserNewPassword     string `json:"newPassword"`
	ReferralCode        string `json:"ReferralCode"`
	// Image details
	Link string `json:"imageLink"`
	// Coupon details
	CouponName          string    `json:"couponName"`
	CouponCode          string    `json:"couponCode"`
	CouponLimit         uint      `json:"couponLimit"`
	CouponAvailableFrom time.Time `json:"couponAvailableFrom"`
	CouponAvailableTill time.Time `json:"couponAvailableTill"`
	CouponPercentage    uint      `json:"couponPercentage"`
	CouponMinPurchase   uint      `json:"couponMinPurchase"`
	CouponMaxDiscount   uint      `json:"couponMaxDiscount"`
}
