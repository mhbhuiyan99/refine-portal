package services

import (
	"errors"
	"testing"

	"refine-portal/models"
	"refine-portal/requests"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetCategory_Success(t *testing.T) {

	slug := "dhaka"
	country := "BD"

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
		func(s, c string) (*models.CategoryResponse, error) {

			assert.Equal(t, slug, s)
			assert.Equal(t, country, c)

			return expected, nil
		},
	)

	patches.ApplyFunc(
		requests.GetURLFromConfig,
		func(key string) (string, error) {
			return "https://images.test.com", nil
		},
	)

	defer patches.Reset()

	result, err := GetCategory(slug, country)

	assert.NoError(t, err)

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
		func(string, string) (*models.CategoryResponse, error) {
			return nil, expectedErr
		},
	)

	defer patches.Reset()

	result, err := GetCategory("dhaka", "BD")

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
		func(string, string) (*models.CategoryResponse, error) {
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

	result, err := GetCategory("dhaka", "BD")

	assert.Nil(t, result)
	assert.Equal(t, expectedErr, err)
}