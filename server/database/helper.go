package database

import (
	"github.com/jinzhu/gorm"
	// import for sqlite3
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	gormDB *gorm.DB
)

// Start function
// function init database
func Start() {
	db, err := gorm.Open("sqlite3", "../database/data.db")
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	gormDB = db
	db.AutoMigrate(&Product{}, &Buyer{}, &Purchase{}, &Distributor{}, &Type{}, &Order{}, &Seller{}, &Topic{}, &Comment{}, &Package{}, &Box{})
}
