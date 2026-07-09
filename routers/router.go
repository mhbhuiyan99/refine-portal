package routers

import (
	"refine-portal/controllers"
	"github.com/beego/beego/v2/server/web"
)

func init() {
    web.Router("/refine", &controllers.RefineController{})

	web.Router("/api/location", &controllers.APIController{}, "get:GetLocation")
	web.Router("/api/properties", &controllers.APIController{}, "get:GetProperties")
	web.Router("/api/property-details", &controllers.APIController{}, "get:GetPropertyDetails")
}
