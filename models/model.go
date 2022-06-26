package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/lovehome?charset=utf8", 30)

	// register model
	orm.RegisterModel(new(User), new(House), new(Area), new(Facility), new(HouseImage), new(OrderHouse))

	// create table
	orm.RunSyncdb("default", false, true)
}

type User struct {
	Id           int           `json:"user_id"`
	Name         string        `orm:"size(32);unique" json:"name"`
	PasswordHash string        `orm:"size(128)"json:"password"`
	Mobile       string        `orm:"size(11)" json:"mobile"`
	RealName     string        `orm:"size(32)" json:"real_Name"`
	IdCard       string        `orm:"size(20)" json:"id_card"`
	avatarUrl    string        `orm:"size(256)" json:"avatar_url"`
	Houses       []*House      `orm:"reverse(many) " json:"houses"`
	Orders       []*OrderHouse `orm:"reverse(many)" json:"orders"`
}

type House struct {
	Id            int           `json:"house_id"`
	User          *User         `orm:"rel(fk)" json:"user_id"`
	Area          *Area         `orm:"rel(fk)" json:"area_id"`
	Title         string        `orm:"size(64)"json:"title"`
	Price         int           `orm:"default(0)" json:"price"`
	Address       string        `orm:"size(512)"orm:"default("")"json:"address"`
	RoomCount     int           `orm:"default(1)"json:"room_count"`
	Acreage       int           `orm:"default(0)"json:"acreage"`
	Unit          string        `orm:"size(32)"orm:"default("")"json:"unit"`
	Capacity      int           `orm:"default(1)"json:"capacity"`
	Beds          string        `orm:"size(64)"orm:"default("")"json:"beds"`
	Deposit       int           `orm:"default(0)" json:"deposit"`
	MinDays       int           `orm:"default(1）"json:"min_days"`
	MaxDays       int           `orm:"default(0)"json:"max_days"`
	OrderCount    int           `orm:"default(0)"json:"order_count"`
	IndexImageUrl string        `orm:"size(256)"orm:"default("")" json:"index_image_url"`
	Facilities    []*Facility   `orm:"reverse(many)"json:"facilities"`
	Images        []*HouseImage `orm:"reverse(many)"json:"img_urls""`
	Orders        []*OrderHouse `orm:"reverse(many)"json:"orders"`
	Ctime         time.Time     `orm:"auto_now_add;type(datetime)"json:"ctime"`
}

// HOME_PAGE_MAX_HOUSES 首页最高展示的房屋数量
var HOME_PAGE_MAX_HOUSES int = 5

// HOUSE_LIST_PAGE_CAPACITY 房屋列表页面每页显示条目数
var HOUSE_LIST_PAGE_CAPACITY int = 2

// Area  区域信息 table_name = area
type Area struct {
	Id     int      `json:"aid"`                        //区域编号
	Name   string   `orm:"size(32)" json:"aname"`       //区域名字
	Houses []*House `orm:"reverse(many)" json:"houses"` //区域所有的房屋
}

// Facility 设施信息 table_name = "facility"
type Facility struct {
	Id     int      `json:"fid"`     //设施编号
	Name   string   `orm:"size(32)"` //设施名字
	Houses []*House `orm:"rel(m2m)"` //都有哪些房屋有此设施
}

// HouseImage 房屋图片 table_name = "house_image"
type HouseImage struct {
	Id    int    `json:"house_image_id"`         //图片id
	Url   string `orm:"size(256)" json:"url"`    //图片url
	House *House `orm:"rel(fk)" json:"house_id"` //图片所属房屋编号
}

const (
	ORDER_STATUS_WAIT_ACCEPT  = "WAIT_ACCEPT"  //待接单
	ORDER_STATUS_WAIT_PAYMENT = "WAIT_PAYMENT" //待支付
	ORDER_STATUS_PAID         = "PAID"         //已支付
	ORDER_STATUS_WAIT_COMMENT = "COMMENT"      //待评价
	ORDER_STATUS_COMPLETE     = "COMPLETE"     //已完成
	ORDER_STATUS_CANCELED     = "CONCELED"     //已取消
	ORDER_STATUS_REJECTED     = "REJECTED"     //已拒单
)

// OrderHouse 订单 table_name = order
type OrderHouse struct {
	Id          int       `json:"order_id"`               //订单编号
	User        *User     `orm:"rel(fk)" json:"user_id"`  //下单的用户编号
	House       *House    `orm:"rel(fk)" json:"house_id"` //预定的房间编号
	Begin_date  time.Time `orm:"type(datetime)"`          //预定的起始时间
	End_date    time.Time `orm:"type(datetime)"`          //预定的结束时间
	Days        int       //预定总天数
	House_price int       //房屋的单价
	Amount      int       //订单总金额
	Status      string    `orm:"default(WAIT_ACCEPT)"` //订单状态
	Comment     string    `orm:"size(512)"`            //订单评论
	Ctime       time.Time `orm:"auto_now_add;type(datetime)" json:"ctime"`
}

/*house := House{Id:1}
houseImage:= HouseImage{Url:"group1..", &house}*/
