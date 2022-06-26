package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	_ "iHome/models"
	_ "iHome/routers"
	"net/http"
	"strings"
)

func ignoreStaticPath() {
	beego.InsertFilter("/", beego.BeforeRouter, TransparentStatic)
	beego.InsertFilter("/*", beego.BeforeRouter, TransparentStatic)
}

func TransparentStatic(ctx *context.Context) {
	orpath := ctx.Request.URL.Path
	beego.Debug("request url: ", orpath)
	if strings.Index(orpath, "api") >= 0 {
		return
	}
	http.ServeFile(ctx.ResponseWriter, ctx.Request, "static/html/"+ctx.Request.URL.Path)
}
func main() {
	ignoreStaticPath()
	beego.Run()
}
