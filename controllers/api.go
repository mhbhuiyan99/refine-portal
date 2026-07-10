package controllers

import (
	"refine-portal/services"

	"github.com/beego/beego/v2/server/web"
)

type APIController struct {
	web.Controller
}

func (c *APIController) GetLocation() {
	keyword := c.GetString("keyword")

	location, err := services.GetLocation(keyword)
	if err != nil {
		c.CustomAbort(500, err.Error())
		return
	}

	c.Data["json"] = location
	c.ServeJSON()
}

func (c *APIController) GetProperties() {

}

func (c *APIController) GetPropertyDetails() {

}