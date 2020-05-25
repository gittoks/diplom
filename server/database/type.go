package database

import (
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
)

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

func GetTypesOffsetLimit(offset, limit uint) ([]Type, error) {
	types := []Type{}
	err := gormDB.Select("*").Offset(offset).Limit(limit).Find(&types).Error
	return types, err
}

func GetTypesByID(id uint) (Type, error) {
	types := Type{}
	err := gormDB.Where("id = ?", id).First(&types).Error
	return types, err
}

func GetTypesCount() uint {
	count := 0
	gormDB.Model(&Type{}).Count(&count)
	return uint(count)
}

func DeleteTypes(id uint) error {
	return gormDB.Where("id = ?", id).Delete(&Type{}).Error
}

func CreateType(r *http.Request) error {
	return gormDB.Create(&Type{
		Name: r.PostFormValue("name"),
	}).Error
}

func UpdateType(r *http.Request) error {
	id, _ := strconv.Atoi(r.PostFormValue("id"))
	types := Type{}
	types.Name = r.PostFormValue("name")
	types.ID = uint(id)
	return gormDB.Save(&types).Error
}
