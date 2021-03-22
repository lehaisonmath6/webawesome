package routers

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/lehaisonmath6/webawesome/controllers"
)

func init() {
	web.Router("/", &controllers.MainController{})
	web.SetStaticPath("/static", "react/static")
	web.SetStaticPath("/", "react")
}
