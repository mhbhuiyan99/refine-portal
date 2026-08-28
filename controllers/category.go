package controllers

import (
	"errors"
	"net/http"
	"net/url"
	"refine-portal/requests"
	"refine-portal/services"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

type CategoryController struct {
	web.Controller
}

// Get handles normal category page requests.
//
// Responsibilities:
//   - Parse the requested location URL.
//   - Resolve the country code.
//   - Retrieve category data.
//   - Prepare template data.
//   - Render the category page.
func (c *CategoryController) Get() {

	slug := strings.Trim(
		strings.TrimPrefix(
			c.Ctx.Request.URL.Path,
			"/all/",
		),
		"/",
	)

	if slug == "" {
		c.CustomAbort(
			http.StatusNotFound,
			"Not Found",
		)
		return
	}

	logs.Info(
		"[CategoryController] Category URL slug=%s",
		slug,
	)

	// Convert URL format to Category API format.
	categorySlug := strings.ReplaceAll(
		slug,
		"/",
		":",
	)

	logs.Info(
		"[CategoryController] Category API slug=%s",
		categorySlug,
	)

	// Resolve the country code required by the Category API.
	countrySlug := strings.Split(slug, "/")[0]

	location, err := services.GetLocation(countrySlug)

	if err != nil {
		logs.Error(
			"[CategoryController] GetLocation failed | keyword=%s | err=%v",
			countrySlug,
			err,
		)

		c.CustomAbort(
			http.StatusInternalServerError,
			"Internal Server Error",
		)
		return
	}

	countryCode := location.GeoInfo.CountryCode

	// Normal category pages do not have additional
	// sub-category filters.
	params := url.Values{}

	categories, err := services.GetCategory(
		categorySlug,
		countryCode,
		params,
	)

	if err != nil {
		logs.Error(
			"[CategoryController] GetCategory failed | slug=%s | err=%v",
			slug,
			err,
		)

		var httpErr *requests.HTTPError

		if errors.As(err, &httpErr) {
			c.CustomAbort(
				httpErr.StatusCode,
				http.StatusText(httpErr.StatusCode),
			)
			return
		}

		c.CustomAbort(
			http.StatusInternalServerError,
			"Internal Server Error",
		)
		return
	}

	logs.Info(
		"[CategoryController] Location Response: %+v",
		location.GeoInfo,
	)

	c.Data["Title"] = categories.GeoInfo.Name
	c.Data["Category"] = categories

	c.Data["RefineURL"] = buildRefineURL(
		slug,
		params,
	)

	c.TplName = "category.tpl"
}

// buildRefineURL creates the Refine page URL while preserving
// the filters applied to the current category page.
//
// Example:
//
// location = bangladesh/dhaka
// params   = empty
//
// Result:
//
// /refine?search=bangladesh%2Fdhaka
func buildRefineURL(
	location string,
	params url.Values,
) string {

	query := url.Values{}

	query.Set("search", location)

	for key, values := range params {
		for _, value := range values {
			query.Add(key, value)
		}
	}

	return "/refine?" + query.Encode()
}