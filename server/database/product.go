package database

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name          string
	Description   string
	Image         string
	distributorID int
	typeID        int
}

func GetProducts() ([]Product, error) {
	products := []Product{}
	err := gormDB.Select("*").Find(&products).Error
	return products, err
}

func UpdateProducts(product Product) error {
	return gormDB.Save(&product).Error
}

func AddProducts(product Product) error {
	return gormDB.Create(&product).Error
}
