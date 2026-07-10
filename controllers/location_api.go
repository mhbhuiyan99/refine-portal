package controllers

import (
	"refine-portal/services"
	"github.com/beego/beego/v2/server/web"

	"strings"
)

type LocationAPIController struct {
	web.Controller
}

func (c *LocationAPIController) Get() {
	keyword := c.GetString("keyword")

	if strings.TrimSpace(keyword) == "" {
		c.CustomAbort(400, "Keyword is required")
		return
	}

	location, err := services.GetLocation(keyword)
	if err != nil {
		c.CustomAbort(500, "Internal Server Error")
		return
	}

	c.Data["json"] = location
	c.ServeJSON()
}