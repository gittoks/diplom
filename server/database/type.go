package database

import "github.com/jinzhu/gorm"

// Type structure
type Type struct {
	gorm.Model
	Name string
}

func GetTypes() ([]Type, error) {
	types := []Type{}
	err := gormDB.Select("*").Find(&types).Error
	return types, err
}
