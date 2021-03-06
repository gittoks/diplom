package routes

import (
	"net/http"
	"strconv"
	"strings"

	db "github.com/gittoks/diplom/server/database"
)

// BuyerCookie structure
// Cookie information
type BuyerCookie struct {
	ID   uint
	Role uint
}

// SetCookie to ResponseWriter
// BuyerCookie=id,role
func SetCookie(w http.ResponseWriter, cookie BuyerCookie) {
	cookies := &http.Cookie{
		Name:   "BuyerCookie",
		Value:  strconv.Itoa(int(cookie.ID)) + "," + strconv.Itoa(int(cookie.Role)),
		Path:   "/",
		Domain: "localhost",
	}
	http.SetCookie(w, cookies)
}

// GetCookie from Request
// BuyerCookie=id,role
func GetCookie(w http.ResponseWriter, r *http.Request) BuyerCookie {
	buyerCookie := BuyerCookie{0, 0}
	cookie, err := r.Cookie("BuyerCookie")
	if err != nil {
		SetCookie(w, buyerCookie)
	} else {
		strs := strings.Split(cookie.Value, ",")
		id, errID := strconv.Atoi(strs[0])
		role, errRole := strconv.Atoi(strs[1])
		if errID == nil && errRole == nil {
			buyerCookie = BuyerCookie{
				ID:   uint(id),
				Role: uint(role),
			}
		} else {
			SetCookie(w, buyerCookie)
		}
	}
	return buyerCookie
}

// CheckLoginByCookie function
// return true if login
func CheckLoginByCookie(cookie BuyerCookie) bool {
	return cookie.ID != 0
}

// BuyerCookieByBuyer function
// generate cookie from buyer
func BuyerCookieByBuyer(buyer db.Buyer) BuyerCookie {
	return BuyerCookie{
		ID:   buyer.ID,
		Role: buyer.Role,
	}
}
