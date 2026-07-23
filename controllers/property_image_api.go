package controllers

import (
	"refine-portal/services"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type PropertyImageController struct {
	web.Controller
}

// Get handles Property Images API requests.
//
// Responsibilities:
//   - Read the property ID from the request.
//   - Call the Property Images service.
//   - Return the image list as JSON.
func (c *PropertyImageController) Get() {

	propertyID := c.GetString("propertyId")

	if strings.TrimSpace(propertyID) == "" {
		c.CustomAbort(400, "propertyId is required")
		return
	}

	result, err := services.GetPropertyImages(propertyID)
	if err != nil {
		c.CustomAbort(500, err.Error())
		return
	}

	c.Data["json"] = result
	c.ServeJSON()
}