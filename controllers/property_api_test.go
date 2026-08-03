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

func TestPropertyAPIController_GetList_Success(t *testing.T) {
	patches := gomonkey.ApplyFunc(
		services.GetProperties,
		func(req models.PropertyListRequest) (*models.PropertyListResponse, error) {

			assert.Equal(t, "hotel", req.Category)
			assert.Equal(t, "dhaka", req.Locations)
			assert.Equal(t, 2, req.Order)
			assert.Equal(t, 3, req.Page)
			assert.Equal(t, 20, req.Limit)
			assert.Equal(t, 5, req.Items)
			assert.Equal(t, "mobile", req.Device)

			return &models.PropertyListResponse{
				Success: true,
			}, nil
		},
	)
	defer patches.Reset()

	req := httptest.NewRequest(
		http.MethodGet,
		"/property?category=hotel&location=dhaka&order=2&page=3&limit=20&items=5&device=mobile",
		nil,
	)

	rec := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(rec, req)

	controller := &PropertyAPIController{}
	controller.Init(ctx, "", "", nil)

	controller.GetList()

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), `"Success":true`)
}

func TestPropertyAPIController_GetList_DefaultValues(t *testing.T) {
	patches := gomonkey.ApplyFunc(
		services.GetProperties,
		func(req models.PropertyListRequest) (*models.PropertyListResponse, error) {

			assert.Equal(t, "hotel", req.Category)
			assert.Equal(t, "dhaka", req.Locations)
			assert.Equal(t, 1, req.Order)
			assert.Equal(t, 1, req.Page)
			assert.Equal(t, 192, req.Limit)
			assert.Equal(t, 1, req.Items)
			assert.Equal(t, "desktop", req.Device)

			return &models.PropertyListResponse{
				Success: true,
			}, nil
		},
	)
	defer patches.Reset()

	req := httptest.NewRequest(
		http.MethodGet,
		"/property?category=hotel&location=dhaka",
		nil,
	)

	rec := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(rec, req)

	controller := &PropertyAPIController{}
	controller.Init(ctx, "", "", nil)

	controller.GetList()

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestPropertyAPIController_GetList_MissingCategory(t *testing.T) {

	req := httptest.NewRequest(
		http.MethodGet,
		"/property?location=dhaka",
		nil,
	)

	rec := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(rec, req)

	controller := &PropertyAPIController{}
	controller.Init(ctx, "", "", nil)

	assert.Panics(t, func() {
		controller.GetList()
	})

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "category is required")
}

func TestPropertyAPIController_GetList_MissingLocation(t *testing.T) {

	req := httptest.NewRequest(
		http.MethodGet,
		"/property?category=hotel",
		nil,
	)

	rec := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(rec, req)

	controller := &PropertyAPIController{}
	controller.Init(ctx, "", "", nil)

	assert.Panics(t, func() {
		controller.GetList()
	})

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "location is required")
}

func TestPropertyAPIController_GetList_ServiceError(t *testing.T) {

	patches := gomonkey.ApplyFunc(
		services.GetProperties,
		func(req models.PropertyListRequest) (*models.PropertyListResponse, error) {
			return nil, errors.New("service failed")
		},
	)
	defer patches.Reset()

	req := httptest.NewRequest(
		http.MethodGet,
		"/property?category=hotel&location=dhaka",
		nil,
	)

	rec := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(rec, req)

	controller := &PropertyAPIController{}
	controller.Init(ctx, "", "", nil)

	assert.Panics(t, func() {
		controller.GetList()
	})

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Contains(t, rec.Body.String(), "Failed to fetch properties")
}

func TestPropertyAPIController_GetDetails_Success(t *testing.T) {

	patches := gomonkey.ApplyFunc(
		services.GetPropertyDetails,
		func(req models.PropertyDetailsRequest) (*models.PropertyDetailsResponse, error) {

			assert.Equal(
				t,
				[]string{"1", "2", "3"},
				req.PropertyIDList,
			)

			return &models.PropertyDetailsResponse{
				Success: true,
			}, nil
		},
	)
	defer patches.Reset()

	req := httptest.NewRequest(
		http.MethodGet,
		"/details?propertyIdList=1,2,3",
		nil,
	)

	rec := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(rec, req)

	controller := &PropertyAPIController{}
	controller.Init(ctx, "", "", nil)

	controller.GetDetails()

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), `"Success":true`)
}

func TestPropertyAPIController_GetDetails_MissingPropertyIDList(t *testing.T) {

	req := httptest.NewRequest(
		http.MethodGet,
		"/details",
		nil,
	)

	rec := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(rec, req)

	controller := &PropertyAPIController{}
	controller.Init(ctx, "", "", nil)

	assert.Panics(t, func() {
		controller.GetDetails()
	})

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "propertyIdList is required")
}

func TestPropertyAPIController_GetDetails_ServiceError(t *testing.T) {

	patches := gomonkey.ApplyFunc(
		services.GetPropertyDetails,
		func(req models.PropertyDetailsRequest) (*models.PropertyDetailsResponse, error) {
			return nil, errors.New("service failed")
		},
	)
	defer patches.Reset()

	req := httptest.NewRequest(
		http.MethodGet,
		"/details?propertyIdList=1,2,3",
		nil,
	)

	rec := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(rec, req)

	controller := &PropertyAPIController{}
	controller.Init(ctx, "", "", nil)

	assert.Panics(t, func() {
		controller.GetDetails()
	})

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Contains(t, rec.Body.String(), "Internal Server Error")
}