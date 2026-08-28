package controllers

import (
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

// renderNotFound renders the application's Page Not Found UI.
//
// Responsibilities:
//   - Set the HTTP response status to 404.
//   - Render the shared 404 template.
//   - Stop the current controller flow.
func renderNotFound(c *web.Controller) {
	c.Ctx.ResponseWriter.WriteHeader(http.StatusNotFound)

	c.TplName = "404.tpl"

	if err := c.Render(); err != nil {
		return
	}
}