package routes

import (
	"net/http"
	"sort"
	"strconv"
	"time"

	db "github.com/gittoks/diplom/server/database"
)

// BasketHandlerGET Handler
// handler GET method for /basket
func BasketHandlerGET(w http.ResponseWriter, r *http.Request) {
	cookie := GetCookie(w, r)
	if CheckLoginByCookie(cookie) {
		order, _ := db.GetBasketOrder(cookie.ID)
		purchases, mesTxt, mesTyp := GeneratePurchases(order)
		Answer(w, GetNavBar(cookie), purchases, "basket.html", mesTxt, mesTyp, 2)
	} else {
		Answer(w, GetNavBar(cookie), nil, "login.html", "вы не авторизованы", "danger", 3)
	}
}

// BasketHandlerPOST Handler
// handler POST method for /basket
func BasketHandlerPOST(w http.ResponseWriter, r *http.Request) {
	cookie := GetCookie(w, r)
	if CheckLoginByCookie(cookie) {
		mesTyp, mesTxt := "", ""
		switch r.PostFormValue("type") {
		case "delete":
			id, err := strconv.Atoi(r.PostFormValue("id"))
			db.DeletePurchase(uint(id))
			mesTxt, mesTyp = GenerateMessage(err, "не удалось удалить позицию", "успешно удалено")
			break
		case "clear":
			err := db.DeleteBasketOrder(cookie.ID)
			mesTxt, mesTyp = GenerateMessage(err, "не удалось осободить корзину", "корзина пуста")
			break
		case "submit":
			order, _ := db.GetBasketOrder(cookie.ID)
			count := db.CountPurchases(order.ID)
			if count != 0 {
				order.Status = 2
				order.Data = time.Now().Format("15:04 02.01.2006")
				err := db.UpdateOrder(order)
				mesTxt, mesTyp = GenerateMessage(err, "не удалось создать заказ", "заказ успешно создан")
			}
			break
		}
		order, _ := db.GetBasketOrder(cookie.ID)
		purchases, _, _ := GeneratePurchases(order)
		Answer(w, GetNavBar(GetCookie(w, r)), purchases, "basket.html", mesTxt, mesTyp, 2)
	} else {
		Answer(w, GetNavBar(cookie), nil, "login.html", "вы не авторизованы", "danger", 3)
	}
}

// GeneratePurchases function
func GeneratePurchases(order db.Order) (interface{}, string, string) {
	temp, err := db.GetPurchases(order.ID)
	mesTxt, mesTyp := GenerateMessage(err, "не удалось получить данный корзины", "")
	purchases := make([]interface{}, len(temp))
	mass, cost := uint(0), uint(0)

	for i, value := range temp {
		value.Cost = value.Price * value.Count
		value.Mass = value.Weight * value.Count
		mass += value.Mass
		cost += value.Cost
		purchases[i] = value
	}

	boxx := [][]int{}
	boxes, _ := db.GetBoxes()

	sort.Slice(boxes, func(i, j int) bool {
		return boxes[i].Weight < boxes[j].Weight
	})

	for i, b := range boxes {
		if b.Weight >= mass {
			return Basket{
				SumMass:   mass,
				SumCost:   cost,
				Boxes:     []db.Box{b},
				Count:     []int{1},
				Purchases: purchases,
				Order:     order,
				Status:    db.DecodeOrderStatus(order),
			}, mesTxt, mesTyp
		} else if b.Weight*2 >= mass && i != len(boxes)-1 {
			return Basket{
				SumMass:   mass,
				SumCost:   cost,
				Boxes:     []db.Box{b},
				Count:     []int{2},
				Purchases: purchases,
				Order:     order,
				Status:    db.DecodeOrderStatus(order),
			}, mesTxt, mesTyp
		}
	}

	for i := 0; i < len(boxes); i++ {
		for j := i; j < len(boxes); j++ {
			b := boxes[i].Weight + boxes[j].Weight
			if b >= mass {
				boxx = append(boxx, []int{int(b), i, j})
			}
		}
	}

	sort.Slice(boxx, func(i, j int) bool {
		return boxx[i][0] < boxx[j][0]
	})

	if len(boxx) != 0 {
		if boxx[0][1] == boxx[0][2] {
			return Basket{
				SumMass:   mass,
				SumCost:   cost,
				Boxes:     []db.Box{boxes[boxx[0][1]]},
				Count:     []int{2},
				Purchases: purchases,
				Order:     order,
				Status:    db.DecodeOrderStatus(order),
			}, mesTxt, mesTyp
		}
		return Basket{
			SumMass:   mass,
			SumCost:   cost,
			Boxes:     []db.Box{boxes[boxx[0][1]], boxes[boxx[0][2]]},
			Count:     []int{1, 1},
			Purchases: purchases,
			Order:     order,
			Status:    db.DecodeOrderStatus(order),
		}, mesTxt, mesTyp
	} else {
		return Basket{
			SumMass:   mass,
			SumCost:   cost,
			Boxes:     []db.Box{boxes[len(boxes)-1]},
			Count:     []int{int((mass / boxes[len(boxes)-1].Weight)) + 1},
			Purchases: purchases,
			Order:     order,
			Status:    db.DecodeOrderStatus(order),
		}, mesTxt, mesTyp
	}
}
