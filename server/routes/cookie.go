package routes

import (
	"net/http"
    "strconv"
    "strings"
    "time"

    "github.com/gittoks/diplom/server/database"
)

func CheckCookie(w http.ResponseWriter, r *http.Request) database.User {
    user := database.User{}
    cookieId, err1 := r.Cookie("id");
    cookieRl, err2 := r.Cookie("role");
    if err1 != nil || err2 != nil {
        cookieId = &http.Cookie{"id", "0", "/", "localhost", time.Time{}, "", 0, false, false, 0, "", nil}
        cookieRl = &http.Cookie{"role", "0", "/", "localhost", time.Time{}, "", 0, false, false, 0, "", nil}
        http.SetCookie(w, cookieId)
        http.SetCookie(w, cookieRl)
        user.ID = 0
        user.Role = 0
    } else {
        user.ID, _ = strconv.Atoi(strings.Split(cookieId.String(), "=")[1])
        user.Role, _ = strconv.Atoi(strings.Split(cookieRl.String(), "=")[1])
    }
    return user
}

func GenerateNavigationBar(user database.User) []Nav {
    var navs []Nav
    if user.Role == 0 {
        navs = make([]Nav, len(navsGuess))
        copy(navs, navsGuess)
    } else {
        navs = make([]Nav, len(navsLogin))
        copy(navs, navsLogin)
    }
    return navs
}

func CheckLogin(w http.ResponseWriter, r *http.Request, user database.User) bool {
    if (user.ID != 0) {
        infoPageHandler(w, r)
        return false
    }
    return true
}
