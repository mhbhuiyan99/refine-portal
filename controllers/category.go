package controllers

import (
	"fmt"
	"refine-portal/services"
	"strings"

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

	/* Updating {{.Location}}
	In API: 
	"Title": "Luxury Places to Stay in {{.Location}}",
    "SubTitle": "Luxury Places to Stay in or Near {{.Location}}",
	*/

	location := categories.GeoInfo.ShortName

	for i := range categories.Result.Sections {
		categories.Result.Sections[i].Title = strings.ReplaceAll(
			categories.Result.Sections[i].Title,
			"{{.Location}}",
			location,
		)

		categories.Result.Sections[i].SubTitle = strings.ReplaceAll(
			categories.Result.Sections[i].SubTitle,
			"{{.Location}}",
			location,
		)
	}

	c.Data["Title"] = categories.GeoInfo.Name
	c.Data["Category"] = categories

	imageBaseURL, _ := web.AppConfig.String("image_base_url")
	c.Data["ImageBaseURL"] = imageBaseURL
	
	c.TplName = "category.tpl"
}
