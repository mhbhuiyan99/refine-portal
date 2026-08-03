package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
	"github.com/stretchr/testify/assert"
)

func TestRefineController_Get_WithQuery(t *testing.T) {

	req := httptest.NewRequest(
		http.MethodGet,
		"/refine?search=coxsbazar&order=2",
		nil,
	)

	rec := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(rec, req)

	controller := &RefineController{}
	controller.Init(ctx, "", "", nil)

	controller.Get()

	assert.Equal(t, "coxsbazar", controller.Data["Search"])
	assert.Equal(t, "2", controller.Data["Order"])
	assert.Equal(t, "Refine", controller.Data["Title"])
	assert.Equal(t, "refine.tpl", controller.TplName)
}

func TestRefineController_Get_DefaultValues(t *testing.T) {

	req := httptest.NewRequest(
		http.MethodGet,
		"/refine",
		nil,
	)

	rec := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(rec, req)

	controller := &RefineController{}
	controller.Init(ctx, "", "", nil)

	controller.Get()

	assert.Equal(t, "dhaka, Bangladesh", controller.Data["Search"])
	assert.Equal(t, "1", controller.Data["Order"])
	assert.Equal(t, "Refine", controller.Data["Title"])
	assert.Equal(t, "refine.tpl", controller.TplName)
}