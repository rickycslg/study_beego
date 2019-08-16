package routers

import (
    "github.com/astaxie/beego"
    "mytest/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{},"*:Index")
    beego.Router("/login", &controllers.MainController{}, "*:Login")
}
