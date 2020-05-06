package routes

import (
	"net/http"

	db "github.com/gittoks/diplom/server/database"
)

// ProductHandler Handler
// handler for /product
func ProductHandler(w http.ResponseWriter, r *http.Request) {
	var products []interface{}
	temp, err := db.GetProducts()
	mesTxt, mesTyp := GenerateMessage(err, "Неудалось получить данные о прдуктах", "")
	for i, value := range temp {
		products[i] = value
	}
	Answer(w, GetNavBar(GetCookie(w, r)), products, "product.html", mesTxt, mesTyp)
}
