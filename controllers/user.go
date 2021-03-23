package controllers

import (
	"log"

	"github.com/beego/beego/v2/server/web"
)

type UserController struct {
	web.Controller
}

// @Title Login
// @Description User login
// @Param	username	query	string	true	"username"
// @Param	password	query	string	true	"password"
// @router	/login	[post]
func (this *UserController) Login() {
	var username, password string
	this.Ctx.Input.Bind(&username, "username")
	this.Ctx.Input.Bind(&password, "password")
	log.Println("login username", username, "password", password)

	this.Data["json"] = map[string]interface{}{
		"token": "lehaison",
	}
	this.ServeJSON()
}

// @router	/logout [get]
func (this *UserController) Logout() {
	log.Println("user logout")
}
