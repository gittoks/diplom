package database

import (
	"net/http"
	"strconv"

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

type MoreOrder struct {
	Order
	Buyer
	Seller
}

func UpdateOrder(order Order) error {
	return gormDB.Save(&order).Error
}

func DeleteOrder(id uint) error {
	return gormDB.Where("id = ?", id).Delete(&Order{}).Error
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

func GetOrderByID(id uint) (Order, error) {
	order := Order{}
	err := gormDB.Where("id = ?", id).First(&order).Error
	return order, err
}

func GetOrders(buyerID uint) ([]Order, error) {
	orders := []Order{}
	err := gormDB.Where("buyer_id = ? AND status != 1", buyerID).Find(&orders).Error
	return orders, err
}

func GetOrdersOffsetLimit(offset, limit uint) ([]MoreOrder, error) {
	orders := []MoreOrder{}
	err := gormDB.Raw(
		"SELECT * FROM orders o INNER JOIN buyers b ON o.buyer_id=b.id INNER JOIN sellers s ON o.seller_id=s.id WHERE o.deleted_at IS NULL AND s.deleted_at IS NULL AND b.deleted_at IS NULL AND o.status != 1 LIMIT ? OFFSET ?",
		limit, offset).Scan(&orders).Error
	return orders, err
}

func GetOrdersCount() uint {
	count := 0
	gormDB.Where("status != 1").Model(&Order{}).Count(&count)
	return uint(count)
}

func UpdateOrders(r *http.Request) {
	order := Order{}
	id, _ := strconv.Atoi(r.PostFormValue("id"))
	gormDB.Where("id = ?", id).First(&order)
	if order.Status != 4 && order.Status != 5 {
		order.Status++
	}
	gormDB.Save(&order)
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
	case 5:
		return "отменен"
	}
	return ""
}
