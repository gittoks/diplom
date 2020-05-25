package routes

import (
	"fmt"
	"net/http"
	"strconv"

	db "github.com/gittoks/diplom/server/database"
)

const PAGE = 5

func AdminHandlerGET(w http.ResponseWriter, r *http.Request) {
	cookie := GetCookie(w, r)
	if cookie.Role == 1 {

		typesCount := db.GetTypesCount()
		distrbutorsCount := db.GetDistributorsCount()
		productsCount := db.GetProductsCount()
		ordersCount := db.GetOrdersCount()
		sellersCount := db.GetSellersCount()
		packageCount := db.GetPackagesCount()
		boxesCount := db.GetBoxesCount()
		typesID, distrbutorsID, productsID, ordersID, sellersID, packagesID, boxesID, val := 1, 1, 1, 1, 1, 1, 1, r.URL.Query()
		if s := val["types"]; s != nil {
			typesID, _ = strconv.Atoi(s[0])
			if typesCount < uint((typesID-1)*PAGE) {
				typesID = 1
			}
		}
		if s := val["packages"]; s != nil {
			packagesID, _ = strconv.Atoi(s[0])
			if packageCount < uint((packagesID-1)*PAGE) {
				packagesID = 1
			}
		}
		if s := val["distrbutors"]; s != nil {
			distrbutorsID, _ = strconv.Atoi(s[0])
			if distrbutorsCount < uint((distrbutorsID-1)*PAGE) {
				distrbutorsID = 1
			}
		}
		if s := val["products"]; s != nil {
			productsID, _ = strconv.Atoi(s[0])
			if productsCount < uint((productsID-1)*PAGE) {
				productsID = 1
			}
		}
		if s := val["orders"]; s != nil {
			ordersID, _ = strconv.Atoi(s[0])
			if ordersCount < uint((ordersID-1)*PAGE) {
				ordersID = 1
			}
		}
		if s := val["sellers"]; s != nil {
			sellersID, _ = strconv.Atoi(s[0])
			if sellersCount < uint((sellersID-1)*PAGE) {
				sellersID = 1
			}
		}
		if s := val["boxes"]; s != nil {
			boxesID, _ = strconv.Atoi(s[0])
			if boxesCount < uint((boxesID-1)*PAGE) {
				boxesID = 1
			}
		}

		apkg, _ := db.GetPackages()
		atyp, _ := db.GetTypes()
		adis, _ := db.GetDistributors()

		types, _ := db.GetTypesOffsetLimit(uint(typesID-1)*PAGE, PAGE)
		typesNav := make([]uint, GetCount(typesCount))
		if len(typesNav) != 0 {
			typesNav[typesID-1] = 1
		}

		distrbutors, _ := db.GetDistributorsOffsetLimit(uint(distrbutorsID-1)*PAGE, PAGE)
		distrbutorsNav := make([]uint, GetCount(distrbutorsCount))
		if len(distrbutorsNav) != 0 {
			distrbutorsNav[distrbutorsID-1] = 1
		}

		boxes, _ := db.GetBoxesOffsetLimit(uint(boxesID-1)*PAGE, PAGE)
		boxesNav := make([]uint, GetCount(boxesCount))
		if len(boxesNav) != 0 {
			boxesNav[boxesID-1] = 1
		}

		products, _ := db.GetProductsOffsetLimit(uint(productsID-1)*PAGE, PAGE)
		productsNav := make([]uint, GetCount(productsCount))
		if len(productsNav) != 0 {
			productsNav[productsID-1] = 1
		}

		orders, _ := db.GetOrdersOffsetLimit(uint(ordersID-1)*PAGE, PAGE)
		ordersNav := make([]uint, GetCount(ordersCount))
		if len(ordersNav) != 0 {
			ordersNav[ordersID-1] = 1
		}

		sellers, _ := db.GetSellersOffsetLimit(uint(sellersID-1)*PAGE, PAGE)
		sellersNav := make([]uint, GetCount(sellersCount))
		if len(sellersNav) != 0 {
			sellersNav[sellersID-1] = 1
		}

		packages, _ := db.GetPackagesOffsetLimit(uint(packagesID-1)*PAGE, PAGE)
		packagesNav := make([]uint, GetCount(packageCount))
		if len(packagesNav) != 0 {
			packagesNav[packagesID-1] = 1
		}

		Answer(w, GetNavBar(GetCookie(w, r)),
			AdminPage{types, typesNav, distrbutors, distrbutorsNav, products, productsNav, orders, ordersNav, sellers, sellersNav, packages, packagesNav, boxes, boxesNav, apkg, atyp, adis},
			"admin.html", "", "", 6)

	} else {
		Answer(w, GetNavBar(cookie), nil, "info.html", "вы не администратор", "danger", 0)
	}
}

func GetCount(a uint) uint {
	r := a / PAGE
	if a%PAGE != 0 {
		r++
	}
	return r
}

func AdminHandlerPOST(w http.ResponseWriter, r *http.Request) {
	cookie := GetCookie(w, r)
	if cookie.Role == 1 {

		types, distrbutors, products, orders, sellers, packages, boxes, val := 1, 1, 1, 1, 1, 1, 1, r.URL.Query()
		if s := val["types"]; s != nil {
			types, _ = strconv.Atoi(s[0])
		}
		if s := val["distrbutors"]; s != nil {
			distrbutors, _ = strconv.Atoi(s[0])
		}
		if s := val["products"]; s != nil {
			products, _ = strconv.Atoi(s[0])
		}
		if s := val["orders"]; s != nil {
			orders, _ = strconv.Atoi(s[0])
		}
		if s := val["sellers"]; s != nil {
			sellers, _ = strconv.Atoi(s[0])
		}
		if s := val["packages"]; s != nil {
			packages, _ = strconv.Atoi(s[0])
		}
		if s := val["boxes"]; s != nil {
			boxes, _ = strconv.Atoi(s[0])
		}

		switch r.PostFormValue("action") {
		case "action_page_orders":
			orders, _ = strconv.Atoi(r.PostFormValue("id"))
			break
		case "action_page_boxes":
			boxes, _ = strconv.Atoi(r.PostFormValue("id"))
			break
		case "action_page_types":
			types, _ = strconv.Atoi(r.PostFormValue("id"))
			break
		case "action_page_products":
			products, _ = strconv.Atoi(r.PostFormValue("id"))
			break
		case "action_page_distrbutors":
			distrbutors, _ = strconv.Atoi(r.PostFormValue("id"))
			break
		case "action_page_sellers":
			sellers, _ = strconv.Atoi(r.PostFormValue("id"))
			break
		case "action_page_packages":
			packages, _ = strconv.Atoi(r.PostFormValue("id"))
			break
		case "action_delete_orders":
			id, _ := strconv.Atoi(r.PostFormValue("id"))
			order, _ := db.GetOrderByID(uint(id))
			order.Status = 5
			db.UpdateOrder(order)
			break
		case "action_delete_types":
			id, _ := strconv.Atoi(r.PostFormValue("id"))
			db.DeleteTypes(uint(id))
			break
		case "action_delete_products":
			id, _ := strconv.Atoi(r.PostFormValue("id"))
			db.DeleteProducts(uint(id))
			break
		case "action_delete_distributors":
			id, _ := strconv.Atoi(r.PostFormValue("id"))
			db.DeleteDistributors(uint(id))
			break
		case "action_delete_sellers":
			id, _ := strconv.Atoi(r.PostFormValue("id"))
			db.DeleteSellers(uint(id))
			break
		case "action_create_types":
			db.CreateType(r)
			break
		case "action_create_distributors":
			db.CreateDistributor(r)
			break
		case "action_create_sellers":
			db.CreateSeller(r)
			break
		case "action_create_packages":
			db.CreatePackage(r)
			break
		case "action_create_boxes":
			db.CreateBox(r)
			break
		case "action_create_products":
			db.CreateProduct(r)
			break
		case "action_edit_orders":
			db.UpdateOrders(r)
			break
		case "action_edit_types":
			id, _ := strconv.Atoi(r.PostFormValue("id"))
			types, _ := db.GetTypesByID(uint(id))
			Answer(w, GetNavBar(cookie), types, "admin_change_type.html", "", "", 6)
			return
		case "action_edit_packages":
			id, _ := strconv.Atoi(r.PostFormValue("id"))
			pkg, _ := db.GetPackageByID(uint(id))
			Answer(w, GetNavBar(cookie), pkg, "admin_change_package.html", "", "", 6)
			return
		case "action_update_packages":
			db.UpdatePackage(r)
			break
		case "action_edit_products":
			id, _ := strconv.Atoi(r.PostFormValue("id"))
			pkg, _ := db.GetProductByID(uint(id))
			apkg, _ := db.GetPackages()
			atyp, _ := db.GetTypes()
			adis, _ := db.GetDistributors()
			Answer(w, GetNavBar(cookie), ProductAdminPage{pkg, apkg, atyp, adis}, "admin_change_product.html", "", "", 6)
			return
		case "action_update_products":
			db.UpdateProduct(r)
			break
		case "action_update_types":
			db.UpdateType(r)
			break
		case "action_edit_distributors":
			id, _ := strconv.Atoi(r.PostFormValue("id"))
			distrbutor, _ := db.GetDistributorByID(uint(id))
			Answer(w, GetNavBar(cookie), distrbutor, "admin_change_distributor.html", "", "", 6)
			return
		case "action_update_distributors":
			db.UpdateDistributor(r)
			break
		case "action_edit_sellers":
			id, _ := strconv.Atoi(r.PostFormValue("id"))
			seller, _ := db.GetSellerByID(uint(id))
			Answer(w, GetNavBar(cookie), seller, "admin_change_seller.html", "", "", 6)
			return
		case "action_update_sellers":
			db.UpdateSeller(r)
			break
		}

		http.Redirect(w, r,
			fmt.Sprintf("/admin?orders=%d&types=%d&distrbutors=%d&products=%d&sellers=%d&packages=%d&boxes=%d",
				orders, types, distrbutors, products, sellers, packages, boxes,
			), http.StatusSeeOther)
	} else {
		Answer(w, GetNavBar(cookie), nil, "info.html", "вы не администратор", "danger", 0)
	}
}
