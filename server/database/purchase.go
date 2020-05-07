package database

import (
	"github.com/jinzhu/gorm"
)

type Purchase struct {
	gorm.Model
	BuyerID   uint
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
// Select all purchases of buyer
func GetPurchases(id uint) ([]MorePurchase, error) {
	purchases := []MorePurchase{}
	err := gormDB.Raw(
		"SELECT * FROM purchases p INNER JOIN products pr ON pr.id=p.product_id WHERE p.buyer_id=? AND p.deleted_at IS NULL AND pr.deleted_at IS NULL",
		id).Scan(&purchases).Error
	return purchases, err
}

// CreatePurchase functon
// Create purchase or increase
func CreatePurchase(buyerID, productID uint) error {
	purchase := Purchase{}
	err := gormDB.Where("buyer_id = ? AND product_id = ?", buyerID, productID).First(&purchase).Error
	if err != nil {
		err = gormDB.Create(&Purchase{
			BuyerID:   buyerID,
			ProductID: productID,
			Count:     1,
		}).Error
	} else {
		purchase.Count++
		err = gormDB.Save(&purchase).Error
	}
	return err
}

// DeletePurchase functon
// Delete purchase
func DeletePurchase(id uint) error {
	return gormDB.Where("id = ?", id).Delete(&Purchase{}).Error
}

// DeletePurchases functon
// Delete purchase by buyer id
func DeletePurchases(id uint) error {
	return gormDB.Where("buyer_id = ?", id).Delete(&Purchase{}).Error
}
