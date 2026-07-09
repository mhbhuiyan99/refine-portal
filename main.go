package main

import (
	_ "refine-portal/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

