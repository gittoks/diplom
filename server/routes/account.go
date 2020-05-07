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
	if err == nil {
		Answer(w, GetNavBar(cookie), []interface{}{buyer}, "account.html", "", "", 3)
	} else {
		Answer(w, GetNavBar(cookie), nil, "login.html", "вы не авторизованы", "danger", 0)
	}
}

// AccountHandlerPOST Handler
// handler POST method for /account
func AccountHandlerPOST(w http.ResponseWriter, r *http.Request) {
	Answer(w, GetNavBar(GetCookie(w, r)), nil, "account.html", "", "", 3)
}
