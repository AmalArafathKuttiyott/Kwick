package database

import entity "kwick/model/entity"

func GetCategories(p, l int) []entity.Category {
	var categories []entity.Category
	offset := (int(p) - 1) * int(l)
	result := DB.Offset(offset).Limit(l).Find(&categories)
	if result.Error != nil {
		panic("failed to query users: " + result.Error.Error())
	}
	return categories
}
func GetTotalNumberOfCategories() int {
	var categories int64
	if err := DB.Model(entity.Category{}).Count(&categories).Error; err != nil {
		return 0
	}
	return int(categories)
}
func GetCategoryById(i string) (entity.Category, bool) {
	var category entity.Category
	DB.Where("id = ?", i).Find(&category)
	if category.ID == 0 {
		return category, false
	}
	return category, true
}
func AddCategory(cn string) bool {
	var exist entity.Category
	DB.Where("name = ?", cn).First(&exist)
	if exist.ID != 0 {
		return false
	}
	var category entity.Category
	category.Name = cn
	result := DB.Create(&category)
	return result.Error == nil
}
func BlockCategory(i string) (bool, string) {
	var Category entity.Category
	DB.Where("id = ? ", i).First(&Category)
	if Category.ID == 0 {
		return false, "Category does not exist"
	}
	if Category.Blocked {
		return false, "Category is already blocked"
	}
	Category.Blocked = true
	DB.Save(&Category)
	return true, "Category blocked"
}
func UnblockCategory(i string) (bool, string) {
	var category entity.Category
	DB.Where("id = ? ", i).First(&category)
	if category.ID == 0 {
		return false, "Category does not exist"
	}
	if !category.Blocked {
		return false, "Cateory is already unblocked"
	}
	category.Blocked = false
	DB.Save(&category)
	return true, "Category unblocked"
}
func EditCategory(i, c string) bool {
	var category entity.Category
	result := DB.Where("id =?", i).First(&category)
	if result.Error != nil {
		return false
	}
	if category.Name == c {
		return false
	}
	if category.Blocked {
		return false
	}
	category.Name = c
	DB.Save(&category)
	return true
}
func GetProductsByCategory(p, l int, c string) []entity.Product {
	var products []entity.Product
	offset := (int(p) - 1) * int(l)
	result := DB.Where("category = ?", c).Offset(offset).Limit(l).Find(&products)
	if result.Error != nil {
		panic("failed to query users: " + result.Error.Error())
	}
	return products
}
func GetTotalCountOfProductsByCategory(c string) int {
	var categories int64
	if err := DB.Where("category = ?", c).Model(entity.Category{}).Count(&categories).Error; err != nil {
		return 0
	}
	return int(categories)
}
