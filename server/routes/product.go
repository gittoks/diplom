package routes

import (
	"net/http"
	"strconv"

	db "github.com/gittoks/diplom/server/database"
)

// ProductHandlerGET Handler
// handler GET for /product
func ProductHandlerGET(w http.ResponseWriter, r *http.Request) {
	typeID, distributorID := ParseURL(r)
	products, mesTxt, mesTyp := GenerateProducts(typeID, distributorID)
	Answer(w, GetNavBar(GetCookie(w, r)), products, "product.html", mesTxt, mesTyp, 1)
}

// ProductHandlerPOST Handler
// handler POST for /product
func ProductHandlerPOST(w http.ResponseWriter, r *http.Request) {

	switch r.PostFormValue("type") {
	case "buy":
		ProductBuyHandlerPOST(w, r)
		break
	case "clear":
		ProductClearHandlerPOST(w, r)
		break
	}

}

func ProductBuyHandlerPOST(w http.ResponseWriter, r *http.Request) {
	cookie := GetCookie(w, r)
	if CheckLoginByCookie(cookie) {
		var products interface{}
		product, _ := strconv.Atoi(r.PostFormValue("id"))
		order, _ := db.GetBasketOrder(cookie.ID)
		err := db.CreatePurchase(order.ID, uint(product))
		mesTxt, mesTyp := GenerateMessage(err, "не удалось добавить продукт в корзину", "")
		if err == nil {
			typeID, distributorID := ParseURL(r)
			products, mesTxt, mesTyp = GenerateProducts(typeID, distributorID)
		}
		Answer(w, GetNavBar(GetCookie(w, r)), products, "product.html", mesTxt, mesTyp, 1)
	} else {
		Answer(w, GetNavBar(GetCookie(w, r)), nil, "login.html", "вы не авторизованы", "warning", 3)
	}
}

func ProductClearHandlerPOST(w http.ResponseWriter, r *http.Request) {
	products, mesTxt, mesTyp := GenerateProducts(0, 0)
	Answer(w, GetNavBar(GetCookie(w, r)), products, "product.html", mesTxt, mesTyp, 1)
}

// GenerateProducts function
func GenerateProducts(typeID, distributorID uint) (interface{}, string, string) {
	temp, err := db.GetProducts(typeID, distributorID)
	mesTxt, mesTyp := GenerateMessage(err, "не удалось получить данные о продуктах", "")
	products := make([]interface{}, len(temp))
	for i, value := range temp {
		products[i] = value
	}
	return ProductPage{
		Products:     products,
		Types:        GenerateTypes(typeID),
		Distributors: GenerateDistributors(distributorID),
	}, mesTxt, mesTyp
}

func GenerateTypes(id uint) []ProductPageMenu {
	typesArr, _ := db.GetTypes()
	types := make([]ProductPageMenu, len(typesArr))
	for i, value := range typesArr {
		types[i].Name = value.Name
		types[i].ID = value.ID
		if value.ID == id {
			types[i].IsActive = "checked"
		} else {
			types[i].IsActive = ""
		}
	}
	return types
}

func GenerateDistributors(id uint) []ProductPageMenu {
	distributors, _ := db.GetDistributors()
	dis := make([]ProductPageMenu, len(distributors))
	for i, value := range distributors {
		dis[i].Name = value.Name
		dis[i].ID = value.ID
		if value.ID == id {
			dis[i].IsActive = "checked"
		} else {
			dis[i].IsActive = ""
		}
	}
	return dis
}

func ParseURL(r *http.Request) (uint, uint) {
	url := r.URL
	var dis, typ int
	val := url.Query()
	s := val["types_id"]
	if s != nil {
		typ, _ = strconv.Atoi(s[0])
	}
	s = val["distributors_id"]
	if s != nil {
		dis, _ = strconv.Atoi(s[0])
	}
	return uint(typ), uint(dis)
}
