package controllers

import (
	"refine-portal/models"
	"refine-portal/services"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type PropertyAPIController struct {
	web.Controller
}

// GET /api/properties
func (c *PropertyAPIController) GetList() {
	category := strings.TrimSpace(c.GetString("category"))
	location := strings.TrimSpace(c.GetString("location"))

	if category == "" {
		c.CustomAbort(400, "category is required")
		return
	}

	if location == "" {
		c.CustomAbort(400, "location is required")
		return
	}

	order, err := c.GetInt("order")
	if err != nil {
		order = 1
	}

	page, err := c.GetInt("page")
	if err != nil {
		page = 1
	}

	limit, err := c.GetInt("limit")
	if err != nil {
		limit = 192
	}

	items, err := c.GetInt("items")
	if err != nil {
		items = 1
	}

	request := models.PropertyListRequest{
		Category:  category,
		Locations: location,
		Order:     order,
		Page:      page,
		Limit:     limit,
		Items:     items,
		Device:    c.GetString("device", "desktop"),
	}

	properties, err := services.GetProperties(request)
	if err != nil {
		c.CustomAbort(500, "Failed to fetch properties")
		return
	}

	c.Data["json"] = properties
	c.ServeJSON()
}

func (c *PropertyAPIController) GetDetails() {

}
