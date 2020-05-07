package database

import "github.com/jinzhu/gorm"

// Distributor structure
type Distributor struct {
	gorm.Model
	Name string
	City string
}

func GetDistributors() ([]Distributor, error) {
	distributor := []Distributor{}
	err := gormDB.Select("*").Find(&distributor).Error
	return distributor, err
}
