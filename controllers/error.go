package controllers

import (
	"net/http"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)


func renderError(c *web.Controller, status int, tplName string) {
	c.Ctx.ResponseWriter.WriteHeader(status)

	rendered := func() (ok bool) {
		defer func() {
			if r := recover(); r != nil {
				logs.Error(
					"[renderError] template render failed | tpl=%s | recover=%v",
					tplName,
					r,
				)
				ok = false
			}
		}()

		c.TplName = tplName
		return c.Render() == nil
	}()

	if !rendered {
		c.Ctx.WriteString(http.StatusText(status))
	}

	c.StopRun()
}

func renderNotFound(c *web.Controller) {
	renderError(c, http.StatusNotFound, "errors/404.tpl")
}

func renderBadRequest(c *web.Controller) {
	renderError(c, http.StatusBadRequest, "errors/400.tpl")
}

func renderServerError(c *web.Controller) {
	renderError(c, http.StatusInternalServerError, "errors/500.tpl")
}