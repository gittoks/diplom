package database

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name          string
	Description   string
	Image         string
	DistributorID uint
	Weight        uint
	Price         uint
	TypeID        uint
}

func GetProducts(typeID, distributorID uint) ([]Product, error) {
	products := []Product{}
	var err error
	if typeID == 0 && distributorID != 0 {
		err = gormDB.Where("distributor_id = ?", distributorID).Find(&products).Error
	} else if distributorID == 0 && typeID != 0 {
		err = gormDB.Where("type_id = ?", typeID).Find(&products).Error
	} else if distributorID == 0 && typeID == 0 {
		err = gormDB.Select("*").Find(&products).Error
	} else {
		err = gormDB.Where("type_id = ? AND distributor_id = ?", typeID, distributorID).Find(&products).Error
	}
	return products, err
}
