package routes

import (
	"html/template"
	"net/http"
)

// Answer function
// Response to server
func Answer(w http.ResponseWriter, navs []Nav, data []interface{}, name, mesTxt, mesTyp string) {
	tmpl.ExecuteTemplate(w, name, Data{
		Navs:        navs,
		Content:     data,
		Message:     mesTxt,
		MessageType: mesTyp,
	})
}

// GetNavBar function
// Return copy of nav bars
// Login or Guess
func GetNavBar(buyerCookie BuyerCookie) []Nav {
	var navs []Nav
	if buyerCookie.ID == 0 {
		navs = make([]Nav, len(navsGuess))
		copy(navs, navsGuess)
	} else {
		navs = make([]Nav, len(navsLogin))
		copy(navs, navsLogin)
	}
	return navs
}

// GenerateMessage function
// return error.toString() and "danger"
func GenerateMessage(err error, errText, sucText string) (string, string) {
	if err != nil {
		return errText, "danger"
	}
	return sucText, "success"
}

// SwitchHandler adapter
// Switch between GET and POST
func SwitchHandler(get, post func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(MEMORY)
		switch r.Method {
		case "GET":
			get(w, r)
			break
		case "POST":
			post(w, r)
			break
		}
	}
}

// Start function
// init templates and run server
func Start() {

	tmpl, _ = template.ParseGlob("../web/templates/*")

	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("../web"))))

	http.HandleFunc("/", InfoHandler)
	http.HandleFunc("/product", ProductHandler)
	http.HandleFunc("/unlogin", UnloginHandler)
	http.HandleFunc("/login", SwitchHandler(LoginHandlerGET, LoginHandlerPOST))
	http.HandleFunc("/registration", SwitchHandler(RegistrationHandlerGET, RegistrationHandlerPOST))

	http.ListenAndServe(":"+PORT, nil)
}
