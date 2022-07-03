package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/orm"
	_ "github.com/gomodule/redigo/redis"
	"iHome/models"
	"time"
)

type AreaController struct {
	beego.Controller
}

func (this *AreaController) RetData(resp map[string]interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

/*func (c *AreaController) GetArea() {
	beego.Info("getarea success")
	resp := make(map[string]interface{})
	defer c.RetData(resp)
	var areas []models.Area
	o := orm.NewOrm()
	num, err := o.QueryTable("area").All(&areas)
	if err != nil {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(resp["errno"].(string))
		return
	}
	if num == 0 {
		resp["errno"] = models.RECODE_NODATA
		resp["errmsg"] = models.RecodeText(resp["errno"].(string))
		return
	}
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(resp["errno"].(string))
	resp["data"] = areas
	beego.Info("query data succee ,resp = ", resp, "num = ", num)
}
*/

func (c *AreaController) GetArea() {
	beego.Info("connect success")
	resp := make(map[string]interface{})
	defer c.RetData(resp)

	// 从redis获取数据
	cacheConn, err := cache.NewCache("redis", `{"key":"lovehome","conn":"6379","dbNum":"0"}`)
	if err != nil {
		beego.Error("cacheConn err = ", err)
		resp["errno"] = models.RECODE_DATAERR
		resp["errmsg"] = models.RecodeText(resp["errno"].(string))
		return
	}
	areaData := cacheConn.Get("area")
	if areaData != nil {
		beego.Info("get data from cache========")
		// 查询成功返回数据
		resp["errno"] = models.RECODE_OK
		resp["errmsg"] = models.RecodeText(resp["errno"].(string))

		// 从redis中取来的数据必须先解码才能在前台显示
		var areasInfo interface{}
		json.Unmarshal(areaData.([]byte), &areasInfo)
		resp["data"] = areasInfo
		return
	}

	// 第一步，先从数据库中拿到数据
	// 生命一个数组用来存从数据库中查询到的所有城区数据
	o := orm.NewOrm()
	var areas []models.Area

	// 查询area表中的所有数据，存到areas缓存数组中，返回的是int64的查询条数以及error
	num, err := o.QueryTable("area").All(&areas)
	if err != nil {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(resp["errno"].(string))
		return
	}
	if num == 0 {
		resp["errno"] = models.RECODE_NODATA
		resp["errmsg"] = models.RecodeText(resp["errno"].(string))
		return
	}

	// 把取到的数据转换成json格式存入缓存
	jsonStr, err := json.Marshal(areas)
	if err != nil {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(resp["errno"].(string))
		return
	}
	jsonErr := cacheConn.Put("area", jsonStr, time.Second*3600)
	if jsonErr != nil {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(resp["errno"].(string))
		return
	}

	// 查询成功返回数据
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)

	beego.Info("query data success,resp = ", resp, "num = ", num)
}
