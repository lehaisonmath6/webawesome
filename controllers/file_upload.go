package controllers

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/lehaisonmath6/webawesome/models"
)

type FileUploadController struct {
	web.Controller
}

func (this *FileUploadController) Post() {
	defer this.ServeJSON()
	_, header, err := this.GetFile("myFile")
	if err != nil {
		this.Data["json"] = map[string]interface{}{
			"code":    404,
			"message": err.Error(),
		}
		return
	}

	filepath := models.FILE_FOLDER + "/" + header.Filename
	err = this.SaveToFile("myFile", filepath)
	if err != nil {
		this.Data["json"] = map[string]interface{}{
			"code":    443,
			"message": err.Error(),
		}
	}
	this.Data["json"] = map[string]interface{}{
		"code":    200,
		"message": "Success",
		"url":     models.DOWNLOAD_URL + header.Filename,
	}
}
