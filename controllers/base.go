package controllers

import (
    "github.com/astaxie/beego"
    "strings"
)

type baseController struct {
    beego.Controller
}




//获取Ip地址
func (this *baseController) getClientIp() string {
    s := this.Ctx.Request.Header.Get("x-Real-IP")
    if s == ""{
        s = strings.Split(this.Ctx.Request.RemoteAddr,":")[0]
    }
    return s
}