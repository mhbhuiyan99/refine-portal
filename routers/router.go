package routers

import (
	"refine-portal/controllers"
	"github.com/beego/beego/v2/server/web"
)

func init() {
    web.Router("/refine", &controllers.RefineController{})

	web.Router("/api/location", &controllers.LocationAPIController{})
	web.Router("/api/properties", &controllers.PropertyAPIController{}, "get:GetList")
	web.Router("/api/property-details", &controllers.PropertyAPIController{}, "get:GetDetails")
}
