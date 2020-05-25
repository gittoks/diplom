package routes

import (
	"html/template"

	db "github.com/gittoks/diplom/server/database"
)

const (
	// PORT of server
	PORT = "3456"
	// MEMORY for multipart form
	MEMORY = 4096
)

var (
	tmpl      *template.Template
	navsLogin = []Nav{
		Nav{Href: "/", Name: "Информация"},
		Nav{Href: "/product", Name: "Товары"},
		Nav{Href: "/basket", Name: "Корзина"},
		Nav{Href: "/account", Name: "Личный кабинет"},
		Nav{Href: "/forum", Name: "Форум"},
		Nav{Href: "/unlogin", Name: "Выйти"},
	}
	navsGuess = []Nav{
		Nav{Href: "/", Name: "Информация"},
		Nav{Href: "/product", Name: "Товары"},
		Nav{Href: "/registration", Name: "Регистрация"},
		Nav{Href: "/login", Name: "Войти"},
	}
	navsAdmin = []Nav{
		Nav{Href: "/", Name: "Информация"},
		Nav{Href: "/product", Name: "Товары"},
		Nav{Href: "/basket", Name: "Корзина"},
		Nav{Href: "/account", Name: "Личный кабинет"},
		Nav{Href: "/forum", Name: "Форум"},
		Nav{Href: "/unlogin", Name: "Выйти"},
		Nav{Href: "/admin", Name: "Админ"},
	}
)

// Nav struct
// navigation bar prototype
type Nav struct {
	IsActive string
	Href     string
	Name     string
}

// Data struct
// Standart response builder for server
type Data struct {
	Navs        []Nav
	Content     interface{}
	Message     string
	MessageType string
}

// Basket struct
type Basket struct {
	Purchases []interface{}
	Boxes     []db.Box
	Count     []int
	Order     interface{}
	Status    string
	SumMass   uint
	SumCost   uint
}

// ProductPageMenu struct
type ProductPageMenu struct {
	Name     string
	ID       uint
	IsActive string
}

// ProductPage struct
type ProductPage struct {
	Products     []interface{}
	Distributors []ProductPageMenu
	Types        []ProductPageMenu
}

// AccountPage struct
type AccountPage struct {
	Products interface{}
	Buyer    db.Buyer
}

// CommentPage struct
type CommentPage struct {
	Comments interface{}
	Cookie   BuyerCookie
	Topic    db.Topic
}

// AdminPage struct
type AdminPage struct {
	Types           []db.Type
	TypesNav        []uint
	Distributors    []db.Distributor
	DistributorsNav []uint
	Products        []db.MoreProduct
	ProductsNav     []uint
	Orders          []db.MoreOrder
	OrdersNav       []uint
	Sellers         []db.Seller
	SellersNav      []uint
	Packages        []db.Package
	PackagesNav     []uint
	Boxes           []db.Box
	BoxesNav        []uint
	AllPackages     []db.Package
	AllTypes        []db.Type
	AllDistributors []db.Distributor
}

// ProductAdminPage struct
type ProductAdminPage struct {
	Product         db.Product
	AllPackages     []db.Package
	AllTypes        []db.Type
	AllDistributors []db.Distributor
}
