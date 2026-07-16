package controllers

import (
	"refine-portal/services"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/core/logs"
	"strings"
)

type LocationAPIController struct {
	web.Controller
}

func (c *LocationAPIController) Get() {
	keyword := c.GetString("keyword")

	if strings.TrimSpace(keyword) == "" {
		logs.Warn("[LocationAPIController] Missing required query parameter: keyword")
		c.CustomAbort(400, "Keyword is required")
		return
	}

	location, err := services.GetLocation(keyword)
	logs.Info("Location Response: %+v", location)
	if err != nil {
		logs.Error(
			"[LocationAPIController] GetLocation failed | keyword=%s | err=%v",
			keyword,
			err,
		)
		c.CustomAbort(500, "Internal Server Error")
		return
	}

	c.Data["json"] = location
	c.ServeJSON()
}