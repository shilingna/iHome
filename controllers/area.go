package controllers

import (
	"github.com/astaxie/beego"
	"iHome/models"
)

type AreaController struct {
	beego.Controller
}
type AreaResp struct {
	Errno  string      `json:"errno"`
	Errmsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}

func (this *AreaController) RetData(resp interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (this *AreaController) GetAreas() {
	beego.Debug("get /api/v1.0/areas....")
	resp := AreaResp{
		Errno:  models.RECODE_OK,
		Errmsg: models.RecodeText(models.RECODE_OK),
	}

}
