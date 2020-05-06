package routes

import (
	"net/http"

    "github.com/gittoks/diplom/server/database"
)

func productsPageHandler(w http.ResponseWriter, r *http.Request) {
    user := CheckCookie(w, r)
    navs := GenerateNavigationBar(user)

    navs[1].IsActive = "active"
    data := Data{
        Navs: navs,
        Content: []interface{}{
            database.Product{
                Name: "Яйцо",
                Description: "Только что отняли у куриц",
            },
            database.Product{
                Name: "Корова",
                Description: "2 года как родилась у Буренки",
            },
            database.Product{
                Name: "Огурец",
                Description: "Свежий, хороший, покупай давай",
            },
            database.Product{
                Name: "Сергей",
                Description: "Особенный овощ - пасхалка",
            },
            database.Product{
                Name: "Курица",
                Description: "Нет ничего натуральнее чем курица",
            },
        },
    }

    tmpl.ExecuteTemplate(w, "products.html", data)
}
