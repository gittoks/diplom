package routes

import (
	"net/http"
)

func registerPageHandler(w http.ResponseWriter, r *http.Request) {
    user := CheckCookie(w, r)
    navs := GenerateNavigationBar(user)
    if !CheckLogin(w, r, user) {
        return
    }

    navs[2].IsActive = "active"
    data := Data{
        Navs: navs,
        Content: []interface{}{},
    }

    tmpl.ExecuteTemplate(w, "register.html", data)
}
