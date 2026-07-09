package controllers

import "github.com/beego/beego/v2/server/web"

type RefineController struct {	
	web.Controller
}

func (c *RefineController) Get() {
	c.TplName = "refine.tpl"
}