package routes

import (
	"net/http"
)

func infoPageHandler(w http.ResponseWriter, r *http.Request) {
    user := CheckCookie(w, r)
    navs := GenerateNavigationBar(user)

    navs[0].IsActive = "active"
    data := Data{
        Navs: navs,
        Content: []interface{}{},
    }

    tmpl.ExecuteTemplate(w, "info.html", data)
}
