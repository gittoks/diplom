package database

import (
	"math/rand"

	"github.com/jinzhu/gorm"
)

// Seller structure
type Seller struct {
	gorm.Model
	Name  string
	Phone string
}

func GetRandomSeller() uint {
	count := 0
	gormDB.Model(&Seller{}).Count(&count)
	return uint(rand.Intn(count) + 1)
}
