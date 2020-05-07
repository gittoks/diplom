package routes

import (
	"net/http"

	db "github.com/gittoks/diplom/server/database"
)

// LoginHandlerGET Handler
// handler GET method for /login
func LoginHandlerGET(w http.ResponseWriter, r *http.Request) {
	buyerCookie := GetCookie(w, r)
	if !CheckLoginByCookie(buyerCookie) {
		Answer(w, GetNavBar(GetCookie(w, r)), nil, "login.html", "", "", 3)
	} else {
		Answer(w, GetNavBar(buyerCookie), nil, "info.html", "вы уже авторизованы", "warning", 0)
	}
}

// LoginHandlerPOST Handler
// handler POST method for /login
func LoginHandlerPOST(w http.ResponseWriter, r *http.Request) {
	buyerCookie := GetCookie(w, r)
	if !CheckLoginByCookie(buyerCookie) {
		buyer, err := db.GetBuyer(r)
		mesTxt, mesTyp := GenerateMessage(err, "неверный логин или пароль", "успешная авторизация")
		if err == nil {
			cookie := BuyerCookieByBuyer(buyer)
			SetCookie(w, cookie)
			Answer(w, GetNavBar(cookie), nil, "info.html", mesTxt, mesTyp, 0)
		} else {
			Answer(w, GetNavBar(GetCookie(w, r)), nil, "login.html", mesTxt, mesTyp, 3)
		}
	} else {
		Answer(w, GetNavBar(buyerCookie), nil, "info.html", "вы уже авторизованы", "warning", 0)
	}
}
