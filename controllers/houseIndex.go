package controllers

import (
	"github.com/astaxie/beego"
	"iHome/models"
)

type HouseIndexController struct {
	beego.Controller
}

func (this *HouseIndexController) RetData(resp map[string]interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}
func (this *HouseIndexController) GetHouseIndex() {
	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_DBERR
	resp["errmsg"] = models.RecodeText(resp["errno"].(string))
	this.RetData(resp)
}
