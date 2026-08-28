package main

import (
	"refine-portal/controllers"
	_ "refine-portal/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}

