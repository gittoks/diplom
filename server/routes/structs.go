package routes

import (
	"html/template"
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
