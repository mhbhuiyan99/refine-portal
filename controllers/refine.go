package controllers

import "github.com/beego/beego/v2/server/web"

type RefineController struct {	
	web.Controller
}

// Get renders the Refine page.
//
// Responsibilities:
//   - Read the search and sorting parameters.
//   - Populate the template context.
//   - Render the Refine page.
func (c *RefineController) Get() {
	search := c.GetString("search", "dhaka, Bangladesh")
	order := c.GetString("order", "1")

	c.Data["Search"] = search
	c.Data["Order"] = order
	c.Data["Title"] = "Refine"

	c.TplName = "refine.tpl"
}