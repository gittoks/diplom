package routes

import (
	"net/http"

	db "github.com/gittoks/diplom/server/database"
)

// RegistrationHandlerGET Handler
// handler GET method for /registration
func RegistrationHandlerGET(w http.ResponseWriter, r *http.Request) {
	buyerCookie := GetCookie(w, r)
	if !CheckLoginByCookie(buyerCookie) {
		Answer(w, GetNavBar(GetCookie(w, r)), nil, "registration.html", "", "", 2)
	} else {
		Answer(w, GetNavBar(buyerCookie), nil, "info.html", "вы уже авторизованы", "warning", 0)
	}
}

// RegistrationHandlerPOST Handler
// handler POST method for /registration
func RegistrationHandlerPOST(w http.ResponseWriter, r *http.Request) {
	buyerCookie := GetCookie(w, r)
	if !CheckLoginByCookie(buyerCookie) {
		if db.CheckBuyer(r) {
			buyer, _ := db.CreateBuyer(r)
			cookie := BuyerCookieByBuyer(buyer)
			SetCookie(w, cookie)
			Answer(w, GetNavBar(cookie), nil, "info.html", "успешная регистриция", "success", 0)
		} else {
			Answer(w, GetNavBar(buyerCookie), nil, "registration.html", "неверный логин или пароль", "danger", 2)
		}
	} else {
		Answer(w, GetNavBar(buyerCookie), nil, "info.html", "вы уже авторизованы", "warning", 0)
	}
}
