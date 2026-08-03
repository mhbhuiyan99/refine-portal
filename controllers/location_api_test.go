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

func TestLocationController_Get_Success(t *testing.T) {

	patches := gomonkey.ApplyFunc(
		services.GetLocation,
		func(keyword string) (*models.LocationResponse, error) {

			assert.Equal(t, "dhaka", keyword)

			return &models.LocationResponse{
				Success: true,
				GeoInfo: models.LocationGeoInfo{
					Name: "Dhaka",
				},
			}, nil
		},
	)
	defer patches.Reset()

	req := httptest.NewRequest(
		http.MethodGet,
		"/location?keyword=dhaka",
		nil,
	)

	rec := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(rec, req)

	controller := &LocationAPIController{}
	controller.Init(ctx, "", "", nil)

	controller.Get()

	assert.Equal(t, http.StatusOK, rec.Code)

	body := rec.Body.String()

	assert.Contains(t, body, `"Success":true`)
	assert.Contains(t, body, `"Name":"Dhaka"`)
}

func TestLocationController_Get_MissingKeyword(t *testing.T) {

	req := httptest.NewRequest(
		http.MethodGet,
		"/location",
		nil,
	)

	rec := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(rec, req)

	controller := &LocationAPIController{}
	controller.Init(ctx, "", "", nil)

	assert.Panics(t, func() {
		controller.Get()
	})

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "Keyword is required")
}

func TestLocationController_Get_ServiceError(t *testing.T) {

	patches := gomonkey.ApplyFunc(
		services.GetLocation,
		func(string) (*models.LocationResponse, error) {
			return nil, errors.New("service failed")
		},
	)
	defer patches.Reset()

	req := httptest.NewRequest(
		http.MethodGet,
		"/location?keyword=dhaka",
		nil,
	)

	rec := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(rec, req)

	controller := &LocationAPIController{}
	controller.Init(ctx, "", "", nil)

	assert.Panics(t, func() {
		controller.Get()
	})

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Contains(t, rec.Body.String(), "Internal Server Error")
}