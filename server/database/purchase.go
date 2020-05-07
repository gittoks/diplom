package database

import (
	"github.com/jinzhu/gorm"
)

type Purchase struct {
	gorm.Model
	OrderID   uint
	ProductID uint
	Count     uint
	Cost      uint `gorm:"-"`
	Mass      uint `gorm:"-"`
}

type MorePurchase struct {
	Purchase
	Product
}

// GetPurchases functon
// Select all purchases of order
func GetPurchases(id uint) ([]MorePurchase, error) {
	purchases := []MorePurchase{}
	err := gormDB.Raw(
		"SELECT * FROM purchases p INNER JOIN products pr ON pr.id=p.product_id WHERE p.order_id=? AND p.deleted_at IS NULL AND pr.deleted_at IS NULL",
		id).Scan(&purchases).Error
	return purchases, err
}

// CreatePurchase functon
// Create purchase or increase
func CreatePurchase(orderID, productID uint) error {
	purchase := Purchase{}
	err := gormDB.Where("order_id = ? AND product_id = ?", orderID, productID).First(&purchase).Error
	if err != nil {
		err = gormDB.Create(&Purchase{
			OrderID:   orderID,
			ProductID: productID,
			Count:     1,
		}).Error
	} else {
		purchase.Count++
		err = gormDB.Save(&purchase).Error
	}
	return err
}

// CountPurchases functon
func CountPurchases(orderID uint) uint {
	count := 0
	gormDB.Model(&Purchase{}).Where("order_id = ?", orderID).Count(&count)
	return uint(count)
}

// DeletePurchase functon
// Delete purchase
func DeletePurchase(id uint) error {
	return gormDB.Where("id = ?", id).Delete(&Purchase{}).Error
}

// DeletePurchases functon
// Delete purchase by buyer id
func DeletePurchases(id uint) error {
	return gormDB.Where("order_id = ?", id).Delete(&Purchase{}).Error
}
