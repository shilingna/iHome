package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"iHome/models"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) RetData(resp map[string]interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}
func (this *UserController) Reg() {
	resp := make(map[string]interface{})
	defer this.RetData(resp)
	json.Unmarshal(this.Ctx.Input.RequestBody, &resp)
	beego.Info(`resp["mobile"] = `, resp["mobile"])
	beego.Info(`resp["password"] = `, resp["password"])
	beego.Info(`resp["sms_code"] =`, resp["sms_code"])

	o := orm.NewOrm()
	user := models.User{}
	user.PasswordHash = resp["password"].(string)
	user.Mobile = resp["mobile"].(string)
	user.Name = resp["sms_code"].(string)

	// //设置一个session,用来登录后显示用户名
	this.SetSession("name", user.Name)

	id, err := o.Insert(&user)
	if err != nil {
		resp["errno"] = models.RECODE_NODATA
		resp["errmsg"] = models.RecodeText(resp["errno"].(string))
		return
	}
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(resp["errno"].(string))
	beego.Info("reg succee ,id = ", id)
}
