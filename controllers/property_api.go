package controllers

import (
	"refine-portal/models"
	"refine-portal/services"
	"strings"

	"github.com/beego/beego/v2/core/logs"
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
		logs.Warn("[PropertyAPIController] Missing required query parameter: category")
		c.CustomAbort(400, "category is required")
		return
	}

	if location == "" {
		logs.Warn("[PropertyAPIController] Missing required query parameter: location")
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
		logs.Error(
			"[PropertyAPIController] GetProperties failed | category=%s | location=%s | err=%v",
			category,
			location,
			err,
		)
		c.CustomAbort(500, "Failed to fetch properties")
		return
	}

	c.Data["json"] = properties
	c.ServeJSON()
}

func (c *PropertyAPIController) GetDetails() {
	ids := strings.TrimSpace(
		c.GetString("propertyIdList"),
	)

	if ids == "" {
		c.CustomAbort(400, "propertyIdList is required")
		return
	}

	request := models.PropertyDetailsRequest{
		PropertyIDList: strings.Split(ids, ","),
	}

	details, err := services.GetPropertyDetails(request)
	if err != nil {
		logs.Error(
			"[PropertyAPIController] GetPropertyDetails failed | propertyIdList=%s | err=%v",
			ids,
			err,
		)
		c.CustomAbort(500, "Internal Server Error")
		return
	}

	imageBaseURL, _ := web.AppConfig.String("image_base_url")

	for i := range details.Items {

		image := details.Items[i].Property.FeatureImage

		if image != "" {
			details.Items[i].Property.FeatureImage =
				imageBaseURL + image
		}

		// Add partner feed
		if i < len(request.PropertyIDList) {

			propertyID := request.PropertyIDList[i]

			if partnerInfo, ok := details.Result.ItemsByID[propertyID]; ok {

				details.Items[i].Feed = partnerInfo.Feed
			}
		}
	}

	c.Data["json"] = details
	c.ServeJSON()
}
