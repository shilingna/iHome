package controllers

import (
	"github.com/astaxie/beego"
	"iHome/models"
)

type SessionController struct {
	beego.Controller
}

func (this *SessionController) RetData(resp map[string]interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}
func (this *SessionController) GetSessionData() {
	resp := make(map[string]interface{})
	defer this.RetData(resp)
	user := models.User{}

	// 获取session
	name := this.GetSession("name")
	if name != nil {
		user.Name = name.(string) //需要断言
		resp["errno"] = models.RECODE_OK
		resp["errmsg"] = models.RecodeText(resp["errno"].(string))
		resp["data"] = user

		// 把结构体数据传给data
		resp["data"] = user
	}

}

func (this *SessionController) DeleteSessionData() {
	resp := make(map[string]interface{})
	defer this.RetData(resp)
	this.DelSession("name")
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(resp["errno"].(string))
}
