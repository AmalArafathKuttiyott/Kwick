package initializer

import (
	"fmt"
	"kwick/helper/database"
	entity "kwick/model/entity"
	"os"
)

func MigrateTables() {
	database.ConnectDb()
	err := database.DB.AutoMigrate(entity.User{}, entity.Cart{}, entity.Wishlist{}, entity.Product{}, entity.Image{}, entity.Order{}, entity.Status{}, entity.Address{}, entity.Category{}, entity.Coupon{}, entity.Referral{}, entity.Business{}, entity.Vendors{}, entity.Payment{})
	if err != nil {
		os.Exit(1)
	} else {
		fmt.Println("[ PROGRAM STATUS ] : Migrated Tables To Database Successfully ")
	}
}
