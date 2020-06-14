package database

import (
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
)

type Box struct {
	gorm.Model
	Name   string
	Weight uint
}

func GetBoxes() ([]Box, error) {
	pkgs := []Box{}
	err := gormDB.Select("*").Find(&pkgs).Error
	return pkgs, err
}

func GetBoxesOffsetLimit(offset, limit uint) ([]Box, error) {
	pkgs := []Box{}
	err := gormDB.Select("*").Offset(offset).Limit(limit).Find(&pkgs).Error
	return pkgs, err
}

func GetBoxesCount() uint {
	count := 0
	gormDB.Model(&Box{}).Count(&count)
	return uint(count)
}

func CreateBox(r *http.Request) error {
	weight, _ := strconv.Atoi(r.PostFormValue("weight"))
	return gormDB.Create(&Box{
		Name:   r.PostFormValue("name"),
		Weight: uint(weight),
	}).Error
}

func GetBoxByID(id uint) (Box, error) {
	pkg := Box{}
	err := gormDB.Where("id = ?", id).First(&pkg).Error
	return pkg, err
}

func DeleteBoxes(id uint) (error) {
	return gormDB.Where("id = ?", id).Delete(&Box{}).Error
}

func UpdateBox(r *http.Request) error {
	id, _ := strconv.Atoi(r.PostFormValue("id"))
	weight, _ := strconv.Atoi(r.PostFormValue("weight"))
	types := Box{}
	types.Name = r.PostFormValue("name")
	types.Weight = uint(weight)
	types.ID = uint(id)
	return gormDB.Save(&types).Error
}
