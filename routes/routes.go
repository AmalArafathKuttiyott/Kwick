package routes

import (
	adminController "kwick/controller/admin"
	staticController "kwick/controller/static"
	userController "kwick/controller/user"
	middleware "kwick/middleware"

	gin "github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {

	r.GET("/", adminController.GithubResponse)

	// Admin routes without authorization
	r.POST("/admin/sign-in", adminController.Signin)
	// Group of admin routes with authorization
	admin := r.Group("/admin")
	admin.Use(middleware.AdminAuthorization)
	{
		// Dashboard
		admin.GET("/dashboard", adminController.GetDashboard)
		// User Management
		admin.GET("/users", adminController.GetAllUsers)
		admin.GET("/user", adminController.GetUser)
		admin.PATCH("/block-user", adminController.BlockUser)
		admin.PATCH("/unblock-user", adminController.UnblockUser)
		// Vendor Management
		admin.GET("/vendors", adminController.GetAllVendors)
		admin.GET("/vendor", adminController.GetVendor)
		admin.GET("/vendorship-requests", adminController.GetVendorRequests)
		admin.PATCH("/accept-vendorship", adminController.AcceptVendorship)
		admin.PATCH("/reject-vendorship", adminController.RejectVendorship)
		admin.PATCH("/block-vendor", adminController.BlockVendor)
		admin.PATCH("/unblock-vendor", adminController.UnblockVendor)
		// Category Management
		admin.GET("/categories", adminController.GetAllCategories)
		admin.GET("/category", adminController.GetCategory)
		admin.POST("/add-category", adminController.AddCategory)
		admin.PUT("/edit-category", adminController.EditCategory)
		admin.PATCH("/block-category", adminController.BlockCategory)
		admin.PATCH("/unblock-category", adminController.UnblockCategory)
		// Product Management
		admin.GET("/products", adminController.GetAllProducts)
		admin.GET("/product", adminController.GetProduct)
		admin.POST("/add-product", adminController.AddProduct)
		admin.PUT("/edit-product", adminController.EditProduct)
		admin.PATCH("/block-product", adminController.BlockProduct)
		admin.PATCH("/unblock-product", adminController.UnblockProduct)
		// Coupon Management
		admin.GET("/coupons", adminController.GetCoupons)
		admin.GET("/coupon", adminController.GetCoupon)
		admin.POST("/create-coupon", adminController.CreateCoupon)
		admin.PUT("/edit-coupon", adminController.EditCoupon)
		// Offer Management
		admin.GET("/offers", adminController.GetOffers)
		admin.GET("/offer", adminController.GetOffer)
		admin.POST("/create-offer", adminController.AddOffer)
		admin.POST("/create-category-offer", adminController.AddCategoryOffer)
		admin.PATCH("/remove-offer", adminController.RemoveOffer)
		admin.PATCH("/remove-category-offer", adminController.RemoveCategoryOffer)
		admin.PATCH("/edit-offer", adminController.EditOffer)
		// Order Management
		admin.GET("/orders", adminController.GetOrders)
		admin.GET("/delivered-orders", adminController.GetDeliveredOrders)
		admin.GET("/return-orders", adminController.GetReturnOrders)
		admin.GET("/pending-orders", adminController.GetPendingOrders)
		admin.GET("/order", adminController.GetOrder)
		admin.PATCH("/edit-order-status", adminController.ChangeOrderStatus)
		// Sales Management
		admin.GET("/download-sales-pdf", adminController.DownloadSalesPdf)
	}

	// User routes without authorization
	r.POST("/sign-up", userController.UserSignup)
	r.POST("/sign-up/otp-verification", userController.SignupOtpVerification)
	r.POST("/sign-in", userController.UserSignin)

	user := r.Group("/user")
	user.Use(middleware.UserAuthorization)
	{
		user.GET("/homepage", userController.GetHomepage)
		user.GET("/products", adminController.GetAllProducts)
		user.GET("/product", adminController.GetProduct)
		user.GET("/categories", adminController.GetAllCategories)
		user.GET("/category", userController.GetCategory)
		user.GET("/profile", userController.GetProfile)
		user.PUT("/update-profile", userController.EditProfile)
		user.POST("/add-address", userController.AddAddress)
		user.DELETE("/delete-address", userController.DeleteAddress)
		user.PUT("/edit-address", userController.EditAddress)
		user.POST("/forgot-password", userController.ForgotPassword)
		user.POST("/forgot-password-otp-verification", userController.ForgotPasswordOtpVerification)
		user.PATCH("/change-password", userController.ChangePassword)
		user.GET("/cart", userController.GetUserCart)
		user.POST("/add-to-cart", userController.AddToCart)
		user.POST("/remove-from-cart", userController.RemoveFromCart)
		user.PATCH("/increase-cart-product-count", userController.IncreaseProductcount)
		user.PATCH("/decrease-cart-product-count", userController.ReduceProductCount)
		user.POST("/add-to-wishlist", userController.AddToWishlist)
		user.PATCH("/remove-from-wishlist", userController.RemoveFromWishlist)
		user.GET("/orders", userController.GetOrders)
		user.POST("/cancel-order", userController.CancelOrder)
		user.POST("/return-order", userController.ReturnOrder)
		user.POST("/checkout", userController.Checkout)
		user.POST("/search", userController.GetSearchResult)
	}

	// Vendor routes without authorization
	r.POST("/vendor/register")
	r.POST("/vendor/register/otp-verification")
	r.POST("/vendor/sign-in")

	vendor := r.Group("/vendor")
	vendor.Use(middleware.VendorAuthorization)
	{
		vendor.GET("/dashboard")
	}

	r.GET("/static/:filename", staticController.HandleStaticFiles)
}
