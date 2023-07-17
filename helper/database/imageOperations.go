package database

import (
	entity "kwick/model/entity"
	"strconv"
)

func AddImages(i uint, s string) bool {
	var image entity.Image
	number := uint(i)
	pid := strconv.FormatUint(uint64(number), 10)
	DB.Where("link = ?", pid+s).Find(&image)
	if image.ID == 0 {
		var img entity.Image
		img.ProductId = i
		img.Link = s
		DB.Create(&img)
		return true
	}
	return false
}
func GetImagesByProductId(i uint) []entity.Image {
	var images []entity.Image
	DB.Where("product_id = ?", i).Find(&images)
	return images
}
