package routes

import (
	"net/http"
	"strconv"

	db "github.com/gittoks/diplom/server/database"
)

// BasketHandlerGET Handler
// handler GET method for /basket
func BasketHandlerGET(w http.ResponseWriter, r *http.Request) {
	cookie := GetCookie(w, r)
	if CheckLoginByCookie(cookie) {
		purchases, mesTxt, mesTyp := GeneratePurchases(cookie.ID)
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
			db.DeletePurchases(cookie.ID)
			break
		case "submit":
			break
		}
		purchases, mesTxt, mesTyp := GeneratePurchases(cookie.ID)
		Answer(w, GetNavBar(GetCookie(w, r)), purchases, "basket.html", mesTxt, mesTyp, 2)
	} else {
		Answer(w, GetNavBar(cookie), nil, "login.html", "вы не авторизованы", "danger", 3)
	}
}

// GeneratePurchases function
func GeneratePurchases(id uint) (interface{}, string, string) {
	temp, err := db.GetPurchases(id)
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
	return Basket{SumMass: mass, SumCost: cost, Purchases: purchases}, mesTxt, mesTyp
}
