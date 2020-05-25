package database

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
)

// Seller structure
type Seller struct {
	gorm.Model
	Name  string
	Phone string
}

func GetSellersCount() uint {
	var count uint
	gormDB.Model(&Seller{}).Count(&count)
	return count
}

func GetRandomSeller() uint {
	return uint(rand.Intn(int(GetSellersCount())) + 1)
}

func GetSellersOffsetLimit(offset, limit uint) ([]Seller, error) {
	sellers := []Seller{}
	err := gormDB.Select("*").Offset(offset).Limit(limit).Find(&sellers).Error
	return sellers, err
}

func GetSellerByID(id uint) (Seller, error) {
	sellers := Seller{}
	err := gormDB.Where("id = ?", id).First(&sellers).Error
	return sellers, err
}

func DeleteSellers(id uint) error {
	return gormDB.Where("id = ?", id).Delete(&Seller{}).Error
}

func CreateSeller(r *http.Request) error {
	return gormDB.Create(&Seller{
		Name:  r.PostFormValue("name"),
		Phone: r.PostFormValue("phone"),
	}).Error
}

func UpdateSeller(r *http.Request) error {
	id, _ := strconv.Atoi(r.PostFormValue("id"))
	s := &Seller{
		Name:  r.PostFormValue("name"),
		Phone: r.PostFormValue("phone"),
	}
	s.ID = uint(id)
	return gormDB.Save(&s).Error
}
