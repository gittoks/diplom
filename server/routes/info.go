package routes

import "net/http"

// InfoHandler Handler
// handler for /
func InfoHandler(w http.ResponseWriter, r *http.Request) {
	Answer(w, GetNavBar(GetCookie(w, r)), nil, "info.html", "", "")
}
