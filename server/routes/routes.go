package routes

import (
	"net/http"
    "html/template"
)

func Start() {

    tmpl, _ = template.ParseGlob("../web/templates/*")

    http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("../web"))))

    http.HandleFunc("/", infoPageHandler)

    http.HandleFunc("/products", productsPageHandler)

    http.HandleFunc("/register", registerPageHandler)

    http.HandleFunc("/login", loginPageHandler)

	http.ListenAndServe(":" + PORT, nil)
}
