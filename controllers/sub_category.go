package controllers

import (
	"net/http"
	"refine-portal/services"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

type SubCategoryController struct {
	web.Controller
}

// Get handles sub-category page requests.
//
// Responsibilities:
//   - Parse the requested sub-category URL.
//   - Extract the last URL segment as the sub-category slug.
//   - Resolve the sub-category from the predefined registry.
//   - Extract the location part from the URL.
//   - Resolve the country code.
//   - Retrieve category data using the sub-category parameters.
//   - Prepare template data.
//   - Render the category template.
func (c *SubCategoryController) Get() {

	slug := strings.Trim(
		strings.TrimPrefix(
			c.Ctx.Request.URL.Path,
			"/all/",
		),
		"/",
	)

	if slug == "" {
		c.CustomAbort(http.StatusNotFound, "Not Found")
		return
	}

	logs.Info(
		"[SubCategoryController] URL slug=%s",
		slug,
	)

	// Split the URL into individual segments.
	segments := strings.Split(slug, "/")

	if len(segments) < 2 {
		c.CustomAbort(
			http.StatusBadRequest,
			"Location is required for sub-category pages",
		)
		return
	}

	// The last segment is always the sub-category slug.
	subCategorySlug := strings.ToLower(
		strings.TrimSpace(
			segments[len(segments)-1],
		),
	)

	// Resolve the sub-category from the predefined registry.
	subCategory, mapped, known :=
		services.LookupSubCategory(subCategorySlug)

	logs.Info(
		"[SubCategoryController] sub-category lookup | slug=%s | mapped=%v | known=%v | key=%s | params=%v",
		subCategorySlug,
		mapped,
		known,
		subCategory.CanonicalKey,
		subCategory.Params,
	)

	// The URL is not registered as a supported sub-category.
	if !known {
		c.CustomAbort(
			http.StatusNotFound,
			"Sub-category not found",
		)
		return
	}

	// The sub-category is known but does not have
	// an API mapping yet.
	if !mapped {
		c.CustomAbort(
			http.StatusNotFound,
			"Not Found",
		)
		return
	}

	// Remove the last URL segment to get the location path.
	locationSlug := strings.Join(
		segments[:len(segments)-1],
		"/",
	)

	if locationSlug == "" {
		c.CustomAbort(
			http.StatusBadRequest,
			"Location is required for sub-category pages",
		)
		return
	}

	// Convert URL format to Category API format.
	categorySlug := strings.ReplaceAll(
		locationSlug,
		"/",
		":",
	)

	logs.Info(
		"[SubCategoryController] Category API slug=%s",
		categorySlug,
	)

	// The first location segment represents the country.
	countrySlug := strings.Split(locationSlug, "/")[0]

	location, err := services.GetLocation(countrySlug)

	if err != nil {
		logs.Error(
			"[SubCategoryController] GetLocation failed | keyword=%s | err=%v",
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

	// Use the parameters defined by the sub-category registry.
	categories, err := services.GetCategory(
		categorySlug,
		countryCode,
		subCategory.Params,
	)

	if err != nil {
		logs.Error(
			"[SubCategoryController] GetCategory failed | location=%s | sub-category=%s | err=%v",
			locationSlug,
			subCategorySlug,
			err,
		)

		c.CustomAbort(
			http.StatusInternalServerError,
			"Internal Server Error",
		)
		return
	}

	logs.Info(
		"[SubCategoryController] Location Response: %+v",
		location.GeoInfo,
	)

	// Prepare data required by the category template.
	c.Data["Title"] = categories.GeoInfo.Name
	c.Data["Category"] = categories

	c.Data["SubCategoryKey"] = subCategory.CanonicalKey

	c.Data["RefineURL"] = buildRefineURL(
		locationSlug,
		subCategory.Params,
	)

	c.TplName = "category.tpl"
}