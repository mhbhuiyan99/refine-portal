package controllers

import (
	"refine-portal/services"

	"github.com/beego/beego/v2/server/web"
)

type LocationAPIController struct {
	web.Controller
}

func (c *LocationAPIController) GetLocation() {
	keyword := c.GetString("keyword")

	location, err := services.GetLocation(keyword)
	if err != nil {
		c.CustomAbort(500, err.Error())
		return
	}

	c.Data["json"] = location
	c.ServeJSON()
}