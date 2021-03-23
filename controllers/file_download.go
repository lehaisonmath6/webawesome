package controllers

import (
	"log"

	"github.com/beego/beego/v2/server/web"
	"github.com/lehaisonmath6/webawesome/models"
)

type FileDownloadController struct {
	web.Controller
}

func (this *FileDownloadController) Get() {
	filename := this.Ctx.Input.Param(":id")
	log.Println("download file", filename)
	this.Ctx.Output.Download(models.FILE_FOLDER + "/" + filename)
}
