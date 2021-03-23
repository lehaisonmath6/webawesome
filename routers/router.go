package routers

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/lehaisonmath6/webawesome/controllers"
)

func init() {
	web.Router("/", &controllers.MainController{})
	web.Router("/fileupload", &controllers.FileUploadController{})
	web.Router("/filedownload/:id", &controllers.FileDownloadController{})
	web.SetStaticPath("/static", "react/static")
	web.SetStaticPath("/", "react")

	ns := web.NewNamespace("/api/v1", web.NSNamespace("/user", web.NSInclude(&controllers.UserController{})))
	web.AddNamespace(ns)
}
