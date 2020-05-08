package routes

import (
	"net/http"

	db "github.com/gittoks/diplom/server/database"
)

// AccountHandlerGET Handler
// handler GET method for /account
func AccountHandlerGET(w http.ResponseWriter, r *http.Request) {
	cookie := GetCookie(w, r)
	buyer, err := db.GetBuyerByID(cookie.ID)
	orders, _ := db.GetOrders(cookie.ID)
	ordersFull := make([]interface{}, len(orders))
	for i, v := range orders {
		ordersFull[i], _, _ = GeneratePurchases(v)
	}
	if err == nil {
		Answer(w, GetNavBar(cookie), AccountPage{ordersFull, buyer}, "account.html", "", "", 3)
	} else {
		Answer(w, GetNavBar(cookie), nil, "login.html", "вы не авторизованы", "danger", 0)
	}
}

// AccountHandlerPOST Handler
// handler POST method for /account
func AccountHandlerPOST(w http.ResponseWriter, r *http.Request) {
	cookie := GetCookie(w, r)
	buyer, err := db.GetBuyerByID(cookie.ID)
	if err == nil {
		str := r.PostFormValue("fname")
		if str != "" {
			buyer.FirstName = str
		}
		str = r.PostFormValue("sname")
		if str != "" {
			buyer.SecondName = str
		}
		str = r.PostFormValue("phone")
		if str != "" {
			buyer.PhoneNumber = str
		}
		str = r.PostFormValue("password")
		if str != "" {
			buyer.Password = str
		}
		db.UpdateBuyer(buyer)
		AccountHandlerGET(w, r)
	} else {
		Answer(w, GetNavBar(cookie), nil, "login.html", "вы не авторизованы", "danger", 0)
	}
}
