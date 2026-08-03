package requests

import (
	"errors"
	"net/http"
	"refine-portal/models"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetCategoryRequest_Success(t *testing.T) {
	patches := gomonkey.NewPatches()
	defer patches.Reset()

	expected := &models.CategoryResponse{
		GeoInfo: models.GeoInfo{
			Name: "Dhaka",
		},
	}

	patches.ApplyFunc(GetURLFromConfig,
		func(string) (string, error) {
			return "https://example.com", nil
		})

	patches.ApplyFunc(BuildURL,
		func(string, string, map[string][]string) (string, error) {
			return "https://example.com/category", nil
		})

	patches.ApplyFunc(NewGETRequest,
		func(string) (*http.Request, error) {
			return &http.Request{}, nil
		})

	patches.ApplyFunc(DoRequest,
		func(req *http.Request, target any) error {
			*(target.(*models.CategoryResponse)) = *expected
			return nil
		})

	result, err := GetCategoryRequest("bangladesh/dhaka", "BD")

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestGetCategoryRequest_ConfigError(t *testing.T) {
	patches := gomonkey.NewPatches()
	defer patches.Reset()

	patches.ApplyFunc(GetURLFromConfig,
		func(string) (string, error) {
			return "", errors.New("config error")
		})

	result, err := GetCategoryRequest("bangladesh/dhaka", "BD")

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestGetCategoryRequest_BuildURLError(t *testing.T) {
	patches := gomonkey.NewPatches()
	defer patches.Reset()

	patches.ApplyFunc(GetURLFromConfig,
		func(string) (string, error) {
			return "https://example.com", nil
		})

	patches.ApplyFunc(BuildURL,
		func(string, string, map[string][]string) (string, error) {
			return "", errors.New("url error")
		})

	result, err := GetCategoryRequest("bangladesh/dhaka", "BD")

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestGetCategoryRequest_NewRequestError(t *testing.T) {
	patches := gomonkey.NewPatches()
	defer patches.Reset()

	patches.ApplyFunc(GetURLFromConfig,
		func(string) (string, error) {
			return "https://example.com", nil
		})

	patches.ApplyFunc(BuildURL,
		func(string, string, map[string][]string) (string, error) {
			return "https://example.com/category", nil
		})

	patches.ApplyFunc(NewGETRequest,
		func(string) (*http.Request, error) {
			return nil, errors.New("request error")
		})

	result, err := GetCategoryRequest("bangladesh/dhaka", "BD")

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestGetCategoryRequest_DoRequestError(t *testing.T) {
	patches := gomonkey.NewPatches()
	defer patches.Reset()

	patches.ApplyFunc(GetURLFromConfig,
		func(string) (string, error) {
			return "https://example.com", nil
		})

	patches.ApplyFunc(BuildURL,
		func(string, string, map[string][]string) (string, error) {
			return "https://example.com/category", nil
		})

	patches.ApplyFunc(NewGETRequest,
		func(string) (*http.Request, error) {
			return &http.Request{}, nil
		})

	patches.ApplyFunc(DoRequest,
		func(*http.Request, any) error {
			return errors.New("request failed")
		})

	result, err := GetCategoryRequest("bangladesh/dhaka", "BD")

	assert.Error(t, err)
	assert.Nil(t, result)
}