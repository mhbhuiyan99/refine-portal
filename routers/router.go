package routers

import (
	"refine-portal/controllers"
	"refine-portal/services"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/refine", &controllers.RefineController{})

	web.Router("/api/location", &controllers.LocationAPIController{})
	web.Router("/api/properties", &controllers.PropertyAPIController{}, "get:GetList")
	web.Router("/api/property-details", &controllers.PropertyAPIController{}, "get:GetDetails")
	web.Router("/api/property/images/v1", &controllers.PropertyImageController{})

	// Register predefined sub-category routes first.
	//
	// Example:
	// /all/*/hotels
	// /all/*/pet-friendly
	// /all/*/pools
	//
	// The wildcard represents the location prefix.
	for _, slug := range services.SubCategorySlugs() {
		web.Router(
			"/all/*/"+slug,
			&controllers.SubCategoryController{},
		)
	}

	// General category route.
	//
	// This acts as the fallback for normal location pages such as:
	// /all/bangladesh
	// /all/bangladesh/dhaka-division
	// /all/bangladesh/dhaka-division/dhaka
	web.Router("/all/*", &controllers.CategoryController{})
}