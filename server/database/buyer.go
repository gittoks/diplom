package database

import (
	"net/http"

	"github.com/jinzhu/gorm"
)

// Buyer structure
// Basic structure for user
type Buyer struct {
	gorm.Model
	FirstName   string
	SecondName  string
	PhoneNumber string
	Login       string
	Password    string
	Role        uint
}

// GetBuyer function
// return buyer if it in database
func GetBuyer(r *http.Request) (Buyer, error) {
	buyer := Buyer{}
	err := gormDB.Where(
		"login = ? AND password = ?",
		r.PostFormValue("login"),
		r.PostFormValue("password"),
	).First(&buyer).Error
	return buyer, err
}

// CheckBuyer function
// return true if login is free
func CheckBuyer(r *http.Request) bool {
	buyer := Buyer{}
	err := gormDB.Where(
		"login = ?",
		r.PostFormValue("login"),
	).First(&buyer).Error
	return err != nil
}

// CreateBuyer function
// create buyer from registration
func CreateBuyer(r *http.Request) (Buyer, error) {
	buyer := Buyer{
		FirstName:   r.PostFormValue("fname"),
		SecondName:  r.PostFormValue("sname"),
		PhoneNumber: r.PostFormValue("phone"),
		Login:       r.PostFormValue("login"),
		Password:    r.PostFormValue("password"),
	}
	err := gormDB.Create(&buyer).Error
	return buyer, err
}

// GetBuyerByID function
// return buyer by it id
func GetBuyerByID(id uint) (Buyer, error) {
	buyer := Buyer{}
	err := gormDB.Where(
		"id = ?", id,
	).First(&buyer).Error
	return buyer, err
}
