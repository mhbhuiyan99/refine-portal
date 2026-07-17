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

func (c *CategoryController) Get() {

	slug := strings.TrimPrefix(
			c.Ctx.Request.URL.Path,
			"/all/",
	)
		
	logs.Info("URL Slug: %s", slug)

	// Convert URL format to API format
	apiSlug := strings.ReplaceAll(slug, "/", ":")
	
	logs.Info("[CategoryController] Fetching category for slug: %s", apiSlug)

	/* Get Country code
	   for locations=BD or US in api
	*/
	locationKeyword := strings.Split(slug, "/")[0] // take the country name only
	location, err := services.GetLocation(locationKeyword)

	if err != nil {
		logs.Error(
			"[CategoryController] GetLocation failed | keyword=%s | err=%v",
			locationKeyword,
			err,
		)

		c.CustomAbort(500, "Internal Server Error")
		return
	}

	countryCode := location.GeoInfo.CountryCode

	categories, err := services.GetCategory(apiSlug, countryCode)

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

	displayLocation := categories.GeoInfo.ShortName

	for i := range categories.Result.Sections {
		categories.Result.Sections[i].Title = strings.ReplaceAll(
			categories.Result.Sections[i].Title,
			"{{.Location}}",
			displayLocation,
		)

		categories.Result.Sections[i].SubTitle = strings.ReplaceAll(
			categories.Result.Sections[i].SubTitle,
			"{{.Location}}",
			displayLocation,
		)
	}

	/* Build image URL 
	   (base url + file name)*/

	imageBaseURL, _ := web.AppConfig.String("image_base_url")

	for i := range categories.Result.Sections {
		for j := range categories.Result.Sections[i].Items {
			if categories.Result.Sections[i].Items[j].Property.FeatureImage != "" {
				categories.Result.Sections[i].Items[j].Property.FeatureImage =
					imageBaseURL + categories.Result.Sections[i].Items[j].Property.FeatureImage
			}
		}
	}

	c.Data["Title"] = categories.GeoInfo.Name
	c.Data["Category"] = categories

	c.TplName = "category.tpl"
}
