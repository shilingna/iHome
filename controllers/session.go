package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
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

func (this *SessionController) Login() {
	resp := make(map[string]interface{})
	defer this.RetData(resp)
	json.Unmarshal(this.Ctx.Input.RequestBody, &resp)
	beego.Info("=====name =", resp["mobile"])
	beego.Info("===== password==", resp["password"])

	// 判断合法性
	if resp["mobile"] == nil || resp["password"] == nil {
		resp["errno"] = models.RECODE_DATAERR
		resp["errmsg"] = models.RecodeText(resp["errno"].(string))
	} else if len(resp["mobile"].(string)) != 11 {
		resp["errno"] = models.RECODE_DATAERR
		resp["errmsg"] = models.RecodeText(resp["errno"].(string))
		return
	}

	// 数据库验证,读取
	o := orm.NewOrm()
	user := models.User{
		Name:         resp["mobile"].(string),
		PasswordHash: resp["password"].(string),
	}
	if err := o.Read(&user); err != nil {
		beego.Info("o.Read(&User) err = ", err)
		resp["errno"] = models.RECODE_DATAERR
		resp["errmsg"] = models.RecodeText(resp["errno"].(string))
		return
	}
	if user.PasswordHash != resp["password"] {
		resp["errno"] = models.RECODE_DATAERR
		resp["errmsg"] = models.RecodeText(resp["errno"].(string))
		return
	}

	// 添加session
	this.SetSession("name", resp["mobile"])
	this.SetSession("mobile", resp["mobile"])
	this.SetSession("userId", user.Id)

	// 返回数据给前端
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(resp["errno"].(string))
}
