package database

import (
	"github.com/jinzhu/gorm"
)

// Order structure
// Status:
// 1 - basket
// 2 - making
// 3 - delivering
// 4 - complete
type Order struct {
	gorm.Model
	BuyerID  uint
	SellerID uint
	Status   uint
	Data     string
}

func UpdateOrder(order Order) error {
	return gormDB.Save(&order).Error
}

func DeleteOrder(buyerID, status uint) error {
	return gormDB.Where("buyer_id = ? AND status = ?", buyerID, status).Delete(&Order{}).Error
}

func DeleteBasketOrder(buyerID uint) error {
	order, _ := GetBasketOrder(buyerID)
	return DeletePurchases(order.ID)
}

func GetOrder(buyerID, status uint) (Order, error) {
	order := Order{}
	err := gormDB.Where("buyer_id = ? AND status = ?", buyerID, status).First(&order).Error
	return order, err
}

func GetOrders(buyerID uint) ([]Order, error) {
	orders := []Order{}
	err := gormDB.Where("buyer_id = ? AND status != 1", buyerID).Find(&orders).Error
	return orders, err
}

func GetBasketOrder(buyerID uint) (Order, error) {
	order, err := GetOrder(buyerID, 1)
	if err != nil {
		order = Order{
			BuyerID:  buyerID,
			SellerID: GetRandomSeller(),
			Status:   1,
			Data:     "",
		}
		err = gormDB.Create(&order).Error
	}
	return order, err
}

func DecodeOrderStatus(order Order) string {
	switch order.Status {
	case 1:
		return "лежит в корзине"
	case 2:
		return "собирается"
	case 3:
		return "уже в пути"
	case 4:
		return "доставлен"
	}
	return ""
}
