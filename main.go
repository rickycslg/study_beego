package main

import (
	_ "mytest/routers"
	"github.com/astaxie/beego"
	_"mytest/models"
)

func main() {
	beego.Run()
}

