package controllers

import (
	"refine-portal/services"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

type CategoryController struct {
	web.Controller
}

// Get handles category page requests.
//
// Responsibilities:
//   - Parse the requested category slug.
//   - Resolve the country code from the Location service.
//   - Retrieve category data from the Category service.
//   - Pass data to the template for rendering.
func (c *CategoryController) Get() {

	slug := strings.TrimPrefix(
			c.Ctx.Request.URL.Path,
			"/all/",
	)
		
	logs.Info(
		"[CategoryController] URL slug=%s",
		slug,
	)

	// Convert URL format to API format
	categorySlug := strings.ReplaceAll(slug, "/", ":")
	
	logs.Info(
		"[CategoryController] Category API slug=%s",
		categorySlug,
	)

	// Resolve the country code required by the Category API.
	countrySlug := strings.Split(slug, "/")[0] // take the country name only
	location, err := services.GetLocation(countrySlug)

	if err != nil {
		logs.Error(
			"[CategoryController] GetLocation failed | keyword=%s | err=%v",
			countrySlug,
			err,
		)

		c.CustomAbort(500, "Internal Server Error")
		return
	}

	countryCode := location.GeoInfo.CountryCode

	categories, err := services.GetCategory(categorySlug, countryCode)

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
