package controllers

import (
	"fmt"
	"refine-portal/services"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

type CategoryController struct {
	web.Controller
}

func (c *CategoryController) Get() {
	country := c.Ctx.Input.Param(":country")
	state := c.Ctx.Input.Param(":state")

	slug := fmt.Sprintf("%s:%s", country, state)
	logs.Info("[CategoryController] Fetching category for slug: %s", slug)

	categories, err := services.GetCategory(slug)

	if err != nil {
		logs.Error(
			"[CategoryController] GetCategory failed | slug=%s | err=%v",
			slug,
			err,
		)
		c.CustomAbort(500, "Internal Server Error")
		return
	}

	c.Data["Title"] = categories.GeoInfo.Name
	c.Data["Category"] = categories
	c.TplName = "category.tpl"
}
