package routers

import (
	"net/http"
	"pair_server/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/static/img/favicon.ico")
}

func init() {
	http.HandleFunc("/favicon.ico", faviconHandler)
	beego.Router("/", &controllers.MainController{}, "get:GetHome")
	beego.Router("/all", &controllers.MainController{}, "get:ShowAll")
	beego.Router("/search", &controllers.MainController{}, "get:Search")
	beego.Router("/add", &controllers.MainController{}, "get:Add")
	beego.Router("/doAdd", &controllers.MainController{}, "post:DoAdd")
	beego.Router("/del/:id", &controllers.MainController{}, "get:Delete")
}
