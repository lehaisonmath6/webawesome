package main

import (
	"log"

	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
	_ "github.com/lehaisonmath6/webawesome/routers"
)

func main() {
	web.BConfig.WebConfig.ViewsPath = "react"
	web.BConfig.MaxUploadSize = 10 * 1024 * 1024 * 1024 // 10  GB upload
	log.Println("maxUploadfile", web.BConfig.MaxUploadSize)
	web.InsertFilter("*", web.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "content-type", "Content-Type", "sessionkey", ""},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Content-Type"},
		AllowCredentials: true,
	}))
	web.Run()
}
