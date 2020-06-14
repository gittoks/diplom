package database

import (
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
)

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

func GetDistributorsOffsetLimit(offset, limit uint) ([]Distributor, error) {
	distributor := []Distributor{}
	err := gormDB.Select("*").Offset(offset).Limit(limit).Find(&distributor).Error
	return distributor, err
}

func GetDistributorsCount() uint {
	count := 0
	gormDB.Model(&Distributor{}).Count(&count)
	return uint(count)
}

func DeleteDistributors(id uint) error {
	a := 0
	if gormDB.Model(&Product{}).Where("distributor_id = ?", id).Count(&a); a == 0{
		return gormDB.Where("id = ?", id).Delete(&Distributor{}).Error
	}
	return nil
}

func CreateDistributor(r *http.Request) error {
	return gormDB.Create(&Distributor{
		Name: r.PostFormValue("name"),
		City: r.PostFormValue("city"),
	}).Error
}

func GetDistributorByID(id uint) (Distributor, error) {
	d := Distributor{}
	err := gormDB.Where("id = ?", id).First(&d).Error
	return d, err
}

func UpdateDistributor(r *http.Request) error {
	id, _ := strconv.Atoi(r.PostFormValue("id"))
	d := &Distributor{
		Name: r.PostFormValue("name"),
		City: r.PostFormValue("city"),
	}
	d.ID = uint(id)
	return gormDB.Save(d).Error
}
