package database

import (
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
)

type Package struct {
	gorm.Model
	Name string
}

func GetPackages() ([]Package, error) {
	pkgs := []Package{}
	err := gormDB.Select("*").Find(&pkgs).Error
	return pkgs, err
}

func GetPackagesOffsetLimit(offset, limit uint) ([]Package, error) {
	pkgs := []Package{}
	err := gormDB.Select("*").Offset(offset).Limit(limit).Find(&pkgs).Error
	return pkgs, err
}

func GetPackagesCount() uint {
	count := 0
	gormDB.Model(&Package{}).Count(&count)
	return uint(count)
}

func CreatePackage(r *http.Request) error {
	return gormDB.Create(&Package{
		Name: r.PostFormValue("name"),
	}).Error
}

func GetPackageByID(id uint) (Package, error) {
	pkg := Package{}
	err := gormDB.Where("id = ?", id).First(&pkg).Error
	return pkg, err
}

func DeletePackages(id uint) (error) {
	a := 0
	if gormDB.Model(&Product{}).Where("package_id = ?", id).Count(&a); a == 0 {
		return gormDB.Where("id = ?", id).Delete(&Package{}).Error
	}
	return nil
}

func UpdatePackage(r *http.Request) error {
	id, _ := strconv.Atoi(r.PostFormValue("id"))
	types := Package{}
	types.Name = r.PostFormValue("name")
	types.ID = uint(id)
	return gormDB.Save(&types).Error
}
