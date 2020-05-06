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
		Nav{Href: "/profile", Name: "Личный кабинет"},
		Nav{Href: "/forum", Name: "Форум"},
		Nav{Href: "/unlogin", Name: "Выйти"},
	}
	navsGuess = []Nav{
		Nav{Href: "/", Name: "Информация"},
		Nav{Href: "/product", Name: "Товары"},
		Nav{Href: "/registration", Name: "Регистрация"},
		Nav{Href: "/login", Name: "Войти"},
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
	Content     []interface{}
	Message     string
	MessageType string
}
