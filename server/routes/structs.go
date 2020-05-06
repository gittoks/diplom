package routes

import (
    "html/template"
)

const (
	PORT = "3456"
    MAX = 4096
)

var (
    tmpl *template.Template
    navsLogin = []Nav{
        Nav{IsActive: "", Href: "/", Name: "Информация"},
        Nav{IsActive: "", Href: "/products", Name: "Товары"},
        Nav{IsActive: "", Href: "/basket", Name: "Корзина"},
        Nav{IsActive: "", Href: "/profile", Name: "Личный кабинет"},
        Nav{IsActive: "", Href: "/forum", Name: "Форум"},
        Nav{IsActive: "", Href: "/unlogin", Name: "Выйти"},
    }
    navsGuess = []Nav{
        Nav{IsActive: "", Href: "/", Name: "Информация"},
        Nav{IsActive: "", Href: "/products", Name: "Товары"},
        Nav{IsActive: "", Href: "/register", Name: "Регистрация"},
        Nav{IsActive: "", Href: "/login", Name: "Войти"},
    }
)

type Nav struct {
    IsActive string
    Href string
    Name string
}

type Data struct {
    Navs []Nav
    Content []interface{}
}
