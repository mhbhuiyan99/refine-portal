package controllers

import (
	"net/http"
	"net/http/httptest"
	"refine-portal/models"
	"refine-portal/services"
	"testing"
	"errors"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/stretchr/testify/assert"
)


func TestPropertyImageController_Get_Success(t *testing.T) {

	patches := gomonkey.ApplyFunc(
		services.GetPropertyImages,
		func(propertyID string) (*models.PropertyImagesResponse, error) {

			assert.Equal(t, "123", propertyID)

			return &models.PropertyImagesResponse{
				Success: true,
			}, nil
		},
	)
	defer patches.Reset()

	req := httptest.NewRequest(
		http.MethodGet,
		"/property-images?propertyId=123",
		nil,
	)

	rec := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(rec, req)

	controller := &PropertyImageController{}
	controller.Init(ctx, "", "", nil)

	controller.Get()

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), `"Success":true`)
}

func TestPropertyImageController_Get_MissingPropertyID(t *testing.T) {

	req := httptest.NewRequest(
		http.MethodGet,
		"/property-images",
		nil,
	)

	rec := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(rec, req)

	controller := &PropertyImageController{}
	controller.Init(ctx, "", "", nil)

	assert.Panics(t, func() {
		controller.Get()
	})

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "propertyId is required")
}

func TestPropertyImageController_Get_ServiceError(t *testing.T) {

	patches := gomonkey.ApplyFunc(
		services.GetPropertyImages,
		func(propertyID string) (*models.PropertyImagesResponse, error) {
			return nil, errors.New("service failed")
		},
	)
	defer patches.Reset()

	req := httptest.NewRequest(
		http.MethodGet,
		"/property-images?propertyId=123",
		nil,
	)

	rec := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(rec, req)

	controller := &PropertyImageController{}
	controller.Init(ctx, "", "", nil)

	assert.Panics(t, func() {
		controller.Get()
	})

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Contains(t, rec.Body.String(), "service failed")
}