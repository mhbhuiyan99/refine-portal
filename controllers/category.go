package controllers

import (
	"errors"
	"net/http"
	"net/url"
	"refine-portal/services"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

type CategoryController struct {
	web.Controller
}

// Get handles category and sub-category page requests.
//
// Responsibilities:
//   - Parse the requested category URL.
//   - Resolve an optional sub-category.
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
		c.CustomAbort(http.StatusNotFound, "Not Found")
		return
	}

	logs.Info(
		"[CategoryController] URL slug=%s",
		slug,
	)

	locationSlug, subCategory, err := resolveSubCategory(slug)

	if err != nil {
		status := http.StatusInternalServerError

		var routeErr *SubCategoryRouteError

		if errors.As(err, &routeErr) {
			status = routeErr.StatusCode
		}

		logs.Warn(
			"[CategoryController] Invalid sub-category route | slug=%s | err=%v",
			slug,
			err,
		)

		c.CustomAbort(status, err.Error())
		return
	}

	// Convert URL format to Category API format.
	categorySlug := strings.ReplaceAll(
		locationSlug,
		"/",
		":",
	)

	logs.Info(
		"[CategoryController] Category API slug=%s",
		categorySlug,
	)

	// Resolve the country code required by the Category API.
	countrySlug := strings.Split(locationSlug, "/")[0]

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

	// Normal category pages don't have additional filters.
	params := url.Values{}

	if subCategory != nil {
		params = subCategory.Params
	}

	categories, err := services.GetCategory(
		categorySlug,
		countryCode,
		params,
	)

	if err != nil {
		logs.Error(
			"[CategoryController] GetCategory failed | slug=%s | err=%v",
			locationSlug,
			err,
		)

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
		locationSlug,
		params,
	)

	if subCategory != nil {
		c.Data["SubCategoryKey"] = subCategory.CanonicalKey
	}

	c.TplName = "category.tpl"
}

// SubCategoryRouteError represents a client-facing sub-category
// routing problem.
type SubCategoryRouteError struct {
	StatusCode int
	Message    string
}

func (e *SubCategoryRouteError) Error() string {
	return e.Message
}

func resolveSubCategory(
	slug string,
) (string, *services.SubCategory, error) {

	segments := strings.Split(
		strings.Trim(slug, "/"),
		"/",
	)

	lastSegment := strings.ToLower(
		strings.TrimSpace(
			segments[len(segments)-1],
		),
	)

	subCategory, mapped, known :=
		services.LookupSubCategory(lastSegment)

	logs.Info(
		"[CategoryController] sub-category lookup | slug=%s | mapped=%v | known=%v | key=%s | params=%v",
		lastSegment,
		mapped,
		known,
		subCategory.CanonicalKey,
		subCategory.Params,
	)

	// The route is explicitly supported by the specification,
	// but its API mapping has not been provided.
	if known && !mapped {
		return "", nil, &SubCategoryRouteError{
			StatusCode: http.StatusNotImplemented,
			Message: "Sub-category mapping is not available yet",
		}
	}

	// Not a sub-category.
	//
	// Example:
	// /all/usa/texas
	//
	// "texas" remains part of the location.
	if !mapped {
		return slug, nil, nil
	}

	// A sub-category must have a location.
	//
	// Example:
	// /all/hotels
	if len(segments) < 2 {
		return "", nil, &SubCategoryRouteError{
			StatusCode: http.StatusBadRequest,
			Message:    "Location is required for sub-category pages",
		}
	}

	locationSlug := strings.Join(
		segments[:len(segments)-1],
		"/",
	)

	if locationSlug == "" {
		return "", nil, &SubCategoryRouteError{
			StatusCode: http.StatusBadRequest,
			Message:    "Location is required for sub-category pages",
		}
	}

	return locationSlug, &subCategory, nil
}

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