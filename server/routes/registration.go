package routes

import "net/http"

// RegistrationHandlerGET Handler
// handler GET method for /registration
func RegistrationHandlerGET(w http.ResponseWriter, r *http.Request) {
	Answer(w, GetNavBar(GetCookie(w, r)), nil, "registration.html", "", "")
}

// RegistrationHandlerPOST Handler
// handler POST method for /registration
func RegistrationHandlerPOST(w http.ResponseWriter, r *http.Request) {
	Answer(w, GetNavBar(GetCookie(w, r)), nil, "registration.html", "", "")
}
