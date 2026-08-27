package services

import (
	"errors"
	"net/url"
	"testing"

	"refine-portal/models"
	"refine-portal/requests"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetCategory_Success(t *testing.T) {

	slug := "dhaka"
	country := "BD"

	expectedParams := url.Values{
		"amenities": {"11"},
	}

	expected := &models.CategoryResponse{
		GeoInfo: models.GeoInfo{
			ShortName: "Dhaka",
		},
		Result: models.CategoryResult{
			Sections: []models.CategorySection{
				{
					Title:    "Hotels in {{.Location}}",
					SubTitle: "Stay at {{.Location}}",
					Items: []models.Item{
						{
							Property: models.Property{
								FeatureImage: "hotel.jpg",
							},
						},
					},
				},
			},
		},
	}

	patches := gomonkey.NewPatches()

	patches.ApplyFunc(
		requests.GetCategoryRequest,
		func(
			s string,
			c string,
			params url.Values,
		) (*models.CategoryResponse, error) {

			assert.Equal(t, slug, s)
			assert.Equal(t, country, c)
			assert.Equal(t, expectedParams, params)

			return expected, nil
		},
	)

	patches.ApplyFunc(
		requests.GetURLFromConfig,
		func(key string) (string, error) {
			assert.Equal(t, "image_base_url", key)

			return "https://images.test.com", nil
		},
	)

	defer patches.Reset()

	result, err := GetCategory(
		slug,
		country,
		expectedParams,
	)

	assert.NoError(t, err)
	assert.NotNil(t, result)

	assert.Equal(
		t,
		"Hotels in Dhaka",
		result.Result.Sections[0].Title,
	)

	assert.Equal(
		t,
		"Stay at Dhaka",
		result.Result.Sections[0].SubTitle,
	)

	assert.Equal(
		t,
		"https://images.test.com/hotel.jpg",
		result.Result.Sections[0].Items[0].Property.FeatureImage,
	)
}

func TestGetCategory_RequestError(t *testing.T) {

	expectedErr := errors.New("request failed")

	patches := gomonkey.ApplyFunc(
		requests.GetCategoryRequest,
		func(
			string,
			string,
			url.Values,
		) (*models.CategoryResponse, error) {

			return nil, expectedErr
		},
	)

	defer patches.Reset()

	result, err := GetCategory(
		"dhaka",
		"BD",
		nil,
	)

	assert.Nil(t, result)
	assert.Equal(t, expectedErr, err)
}

func TestGetCategory_ConfigError(t *testing.T) {

	response := &models.CategoryResponse{
		GeoInfo: models.GeoInfo{
			ShortName: "Dhaka",
		},
	}

	expectedErr := errors.New("config failed")

	patches := gomonkey.NewPatches()

	patches.ApplyFunc(
		requests.GetCategoryRequest,
		func(
			string,
			string,
			url.Values,
		) (*models.CategoryResponse, error) {

			return response, nil
		},
	)

	patches.ApplyFunc(
		requests.GetURLFromConfig,
		func(string) (string, error) {
			return "", expectedErr
		},
	)

	defer patches.Reset()

	result, err := GetCategory(
		"dhaka",
		"BD",
		nil,
	)

	assert.Nil(t, result)
	assert.Equal(t, expectedErr, err)
}

func TestGetCategory_SubCategoryParams(t *testing.T) {

	slug := "bangladesh"
	country := "BD"

	subCategoryParams := url.Values{
		"amenities": {"11"},
	}

	expected := &models.CategoryResponse{
		GeoInfo: models.GeoInfo{
			ShortName: "Bangladesh",
		},
	}

	patches := gomonkey.NewPatches()

	patches.ApplyFunc(
		requests.GetCategoryRequest,
		func(
			s string,
			c string,
			params url.Values,
		) (*models.CategoryResponse, error) {

			assert.Equal(t, slug, s)
			assert.Equal(t, country, c)

			// Verify that the sub-category parameters
			// reached the request layer unchanged.
			assert.Equal(
				t,
				subCategoryParams,
				params,
			)

			return expected, nil
		},
	)

	patches.ApplyFunc(
		requests.GetURLFromConfig,
		func(string) (string, error) {
			return "https://images.test.com", nil
		},
	)

	defer patches.Reset()

	result, err := GetCategory(
		slug,
		country,
		subCategoryParams,
	)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}