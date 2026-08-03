package controllers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"refine-portal/models"
	"refine-portal/services"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/stretchr/testify/assert"
)

func TestCategoryController_Get_Success(t *testing.T) {

	patches := gomonkey.NewPatches()
	defer patches.Reset()

	patches.ApplyFunc(
		services.GetLocation,
		func(keyword string) (*models.LocationResponse, error) {

			assert.Equal(t, "bangladesh", keyword)

			return &models.LocationResponse{
				GeoInfo: models.LocationGeoInfo{
					CountryCode: "BD",
				},
			}, nil
		},
	)

	patches.ApplyFunc(
		services.GetCategory,
		func(slug string, countryCode string) (*models.CategoryResponse, error) {

			assert.Equal(t, "bangladesh:dhaka", slug)
			assert.Equal(t, "BD", countryCode)

			return &models.CategoryResponse{
				GeoInfo: models.GeoInfo{
					Name: "Dhaka",
				},
			}, nil
		},
	)

	req := httptest.NewRequest(
		http.MethodGet,
		"/all/bangladesh/dhaka",
		nil,
	)

	rec := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(rec, req)

	controller := &CategoryController{}
	controller.Init(ctx, "", "", nil)

	controller.Get()

	assert.Equal(t, "Dhaka", controller.Data["Title"])
	assert.NotNil(t, controller.Data["Category"])
	assert.Equal(t, "category.tpl", controller.TplName)
}

func TestCategoryController_Get_LocationError(t *testing.T) {

	patches := gomonkey.ApplyFunc(
		services.GetLocation,
		func(keyword string) (*models.LocationResponse, error) {
			return nil, errors.New("location failed")
		},
	)
	defer patches.Reset()

	req := httptest.NewRequest(
		http.MethodGet,
		"/all/bangladesh/dhaka",
		nil,
	)

	rec := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(rec, req)

	controller := &CategoryController{}
	controller.Init(ctx, "", "", nil)

	assert.Panics(t, func() {
		controller.Get()
	})

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Contains(t, rec.Body.String(), "Internal Server Error")
}

func TestCategoryController_Get_CategoryError(t *testing.T) {

	patches := gomonkey.NewPatches()
	defer patches.Reset()

	patches.ApplyFunc(
		services.GetLocation,
		func(keyword string) (*models.LocationResponse, error) {

			return &models.LocationResponse{
				GeoInfo: models.LocationGeoInfo{
					CountryCode: "BD",
				},
			}, nil
		},
	)

	patches.ApplyFunc(
		services.GetCategory,
		func(slug string, countryCode string) (*models.CategoryResponse, error) {
			return nil, errors.New("category failed")
		},
	)

	req := httptest.NewRequest(
		http.MethodGet,
		"/all/bangladesh/dhaka",
		nil,
	)

	rec := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(rec, req)

	controller := &CategoryController{}
	controller.Init(ctx, "", "", nil)

	assert.Panics(t, func() {
		controller.Get()
	})

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Contains(t, rec.Body.String(), "Internal Server Error")
}