package routes

import (
	"net/http"

	db "github.com/gittoks/diplom/server/database"
)

// LoginHandlerGET Handler
// handler GET method for /login
func LoginHandlerGET(w http.ResponseWriter, r *http.Request) {
	Answer(w, GetNavBar(GetCookie(w, r)), nil, "login.html", "", "")
}

// LoginHandlerPOST Handler
// handler POST method for /login
func LoginHandlerPOST(w http.ResponseWriter, r *http.Request) {
	buyer, err := db.GetBuyer(r)
	mesTxt, mesTyp := GenerateMessage(err, "Неверный логин или пароль", "Успешная авторизация")
	if err == nil {
		cookie := BuyerCookie{ID: buyer.ID, Role: buyer.Role}
		SetCookie(w, cookie)
		Answer(w, GetNavBar(cookie), nil, "info.html", mesTxt, mesTyp)
	} else {
		Answer(w, GetNavBar(GetCookie(w, r)), nil, "login.html", mesTxt, mesTyp)
	}
}
