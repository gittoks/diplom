package routes

import (
	"net/http"
	"strconv"
	"time"

	db "github.com/gittoks/diplom/server/database"
)

// BasketHandlerGET Handler
// handler GET method for /basket
func BasketHandlerGET(w http.ResponseWriter, r *http.Request) {
	cookie := GetCookie(w, r)
	if CheckLoginByCookie(cookie) {
		order, _ := db.GetBasketOrder(cookie.ID)
		purchases, mesTxt, mesTyp := GeneratePurchases(order)
		Answer(w, GetNavBar(cookie), purchases, "basket.html", mesTxt, mesTyp, 2)
	} else {
		Answer(w, GetNavBar(cookie), nil, "login.html", "вы не авторизованы", "danger", 3)
	}
}

// BasketHandlerPOST Handler
// handler POST method for /basket
func BasketHandlerPOST(w http.ResponseWriter, r *http.Request) {
	cookie := GetCookie(w, r)
	if CheckLoginByCookie(cookie) {
		switch r.PostFormValue("type") {
		case "delete":
			id, _ := strconv.Atoi(r.PostFormValue("id"))
			db.DeletePurchase(uint(id))
			break
		case "clear":
			db.DeleteBasketOrder(cookie.ID)
			break
		case "submit":
			order, _ := db.GetBasketOrder(cookie.ID)
			count := db.CountPurchases(order.ID)
			if count != 0 {
				order.Status = 2
				order.Data = time.Now().Format("15:04 02.01.2006")
				db.UpdateOrder(order)
			}
			break
		}
		order, _ := db.GetBasketOrder(cookie.ID)
		purchases, mesTxt, mesTyp := GeneratePurchases(order)
		Answer(w, GetNavBar(GetCookie(w, r)), purchases, "basket.html", mesTxt, mesTyp, 2)
	} else {
		Answer(w, GetNavBar(cookie), nil, "login.html", "вы не авторизованы", "danger", 3)
	}
}

// GeneratePurchases function
func GeneratePurchases(order db.Order) (interface{}, string, string) {
	temp, err := db.GetPurchases(order.ID)
	mesTxt, mesTyp := GenerateMessage(err, "не удалось получить данный корзины", "")
	purchases := make([]interface{}, len(temp))
	mass, cost := uint(0), uint(0)
	for i, value := range temp {
		value.Cost = value.Price * value.Count
		value.Mass = value.Weight * value.Count
		mass += value.Mass
		cost += value.Cost
		purchases[i] = value
	}
	return Basket{
		SumMass:   mass,
		SumCost:   cost,
		Purchases: purchases,
		Order:     order,
		Status:    db.DecodeOrderStatus(order),
	}, mesTxt, mesTyp
}
