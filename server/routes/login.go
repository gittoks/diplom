package routes

import (
	"net/http"
    "fmt"

    "github.com/gittoks/diplom/server/database"
)

func loginPageHandler(w http.ResponseWriter, r *http.Request) {
    user := CheckCookie(w, r)
    navs := GenerateNavigationBar(user)
    if !CheckLogin(w, r, user) {
        return
    }

    r.ParseMultipartForm(MAX)
    if (r.Method == "POST") {
        loginPageHandlerPOST(w, r, navs, user)
    }

    loginPageHandlerGET(w, r, navs, user)
}

func loginPageHandlerGET(w http.ResponseWriter, r *http.Request, navs []Nav, user database.User) {

    navs[3].IsActive = "active"
    data := Data{
        Navs: navs,
        Content: []interface{}{},
    }

    tmpl.ExecuteTemplate(w, "login.html", data)
}

func loginPageHandlerPOST(w http.ResponseWriter, r *http.Request, navs []Nav, user database.User) {
    fmt.Println("POST", r.PostForm)
}
