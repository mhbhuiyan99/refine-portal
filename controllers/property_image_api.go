package controllers

import (
	"refine-portal/services"

	"github.com/beego/beego/v2/server/web"
)

type PropertyImageController struct {
	web.Controller
}

func (c *PropertyImageController) Get() {

	propertyID := c.GetString("propertyId")

	result, err := services.GetPropertyImages(propertyID)
	if err != nil {
		c.CustomAbort(500, err.Error())
		return
	}

	c.Data["json"] = result
	c.ServeJSON()
}