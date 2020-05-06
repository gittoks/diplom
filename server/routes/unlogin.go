package routes

import (
	"net/http"
)

// UnloginHandler Handler
// handler for /unlogin
func UnloginHandler(w http.ResponseWriter, r *http.Request) {
	cookie := BuyerCookie{ID: 0, Role: 0}
	SetCookie(w, cookie)
	Answer(w, GetNavBar(cookie), nil, "info.html", "", "")
}
