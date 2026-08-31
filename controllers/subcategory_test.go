package controllers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"refine-portal/models"
	"refine-portal/services"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/stretchr/testify/assert"
)

// newSubCategoryController creates a controller with
// a test HTTP request and response context.
func newSubCategoryController(
	path string,
) *SubCategoryController {

	request := httptest.NewRequest(
		http.MethodGet,
		path,
		nil,
	)

	response := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(response, request)

	controller := &SubCategoryController{}
	controller.Init(ctx, "", "", nil)

	return controller
}

func subCategoryStatusCode(controller *SubCategoryController) int {
	if controller.Ctx.ResponseWriter == nil {
		return http.StatusOK
	}
	return controller.Ctx.ResponseWriter.Status
}

func TestSubCategoryController_Get_Success(t *testing.T) {

	controller := newSubCategoryController(
		"/all/bangladesh/rentals-with-pools",
	)

	expectedLocation := &models.LocationResponse{
		GeoInfo: models.LocationGeoInfo{
			Name:        "Bangladesh",
			CountryCode: "BD",
		},
	}

	expectedCategory := &models.CategoryResponse{
		GeoInfo: models.GeoInfo{
			Name: "Bangladesh",
		},
	}

	patches := gomonkey.NewPatches()
	defer patches.Reset()

	patches.ApplyFunc(
		services.GetLocation,
		func(country string) (*models.LocationResponse, error) {

			assert.Equal(
				t,
				"bangladesh",
				country,
			)

			return expectedLocation, nil
		},
	)

	patches.ApplyFunc(
		services.GetCategory,
		func(
			slug string,
			countryCode string,
			params map[string][]string,
		) (*models.CategoryResponse, error) {

			assert.Equal(
				t,
				"bangladesh",
				slug,
			)

			assert.Equal(
				t,
				"BD",
				countryCode,
			)

			assert.Equal(
				t,
				"12",
				params["amenities"][0],
			)

			return expectedCategory, nil
		},
	)

	controller.Get()

	assert.Equal(
		t,
		"Bangladesh",
		controller.Data["Title"],
	)

	assert.Equal(
		t,
		expectedCategory,
		controller.Data["Category"],
	)

	assert.Equal(
		t,
		"pools",
		controller.Data["SubCategoryKey"],
	)

	assert.Equal(
		t,
		"/refine?amenities=12&search=bangladesh",
		controller.Data["RefineURL"],
	)

	assert.Equal(
		t,
		"category.tpl",
		controller.TplName,
	)
}

func TestSubCategoryController_Get_EmptySlug(t *testing.T) {

	controller := newSubCategoryController("/all/")

	assert.Panics(t, func() {
		controller.Get()
	})

	assert.Equal(
		t,
		http.StatusNotFound,
		subCategoryStatusCode(controller),
	)
}

func TestSubCategoryController_Get_MissingLocation(t *testing.T) {

	controller := newSubCategoryController(
		"/all/pools",
	)

	assert.Panics(t, func() {
		controller.Get()
	})

	assert.Equal(
		t,
		http.StatusBadRequest,
		subCategoryStatusCode(controller),
	)
}

func TestSubCategoryController_Get_UnknownSubCategory(t *testing.T) {

	controller := newSubCategoryController(
		"/all/bangladesh/unknown-category",
	)

	assert.Panics(t, func() {
		controller.Get()
	})

	assert.Equal(
		t,
		http.StatusNotFound,
		subCategoryStatusCode(controller),
	)
}

func TestSubCategoryController_Get_KnownButUnmappedSubCategory(t *testing.T) {

	controller := newSubCategoryController(
		"/all/bangladesh/condos",
	)

	assert.Panics(t, func() {
		controller.Get()
	})

	assert.Equal(
		t,
		http.StatusNotFound,
		subCategoryStatusCode(controller),
	)
}

func TestSubCategoryController_Get_LocationError(t *testing.T) {

	controller := newSubCategoryController(
		"/all/bangladesh/rentals-with-pools",
	)

	expectedErr := errors.New("location request failed")

	patches := gomonkey.ApplyFunc(
		services.GetLocation,
		func(string) (*models.LocationResponse, error) {
			return nil, expectedErr
		},
	)
	defer patches.Reset()

	assert.Panics(t, func() {
		controller.Get()
	})

	assert.Equal(
		t,
		http.StatusInternalServerError,
		subCategoryStatusCode(controller),
	)
}

func TestSubCategoryController_Get_CategoryError(t *testing.T) {

	controller := newSubCategoryController(
		"/all/bangladesh/rentals-with-pools",
	)

	location := &models.LocationResponse{
		GeoInfo: models.LocationGeoInfo{
			Name:        "Bangladesh",
			CountryCode: "BD",
		},
	}

	expectedErr := errors.New("category request failed")

	patches := gomonkey.NewPatches()
	defer patches.Reset()

	patches.ApplyFunc(
		services.GetLocation,
		func(string) (*models.LocationResponse, error) {
			return location, nil
		},
	)

	patches.ApplyFunc(
		services.GetCategory,
		func(
			string,
			string,
			map[string][]string,
		) (*models.CategoryResponse, error) {
			return nil, expectedErr
		},
	)

	assert.Panics(t, func() {
		controller.Get()
	})

	assert.Equal(
		t,
		http.StatusInternalServerError,
		subCategoryStatusCode(controller),
	)
}