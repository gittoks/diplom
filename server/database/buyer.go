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
