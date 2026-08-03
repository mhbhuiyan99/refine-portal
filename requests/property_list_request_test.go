package requests

import (
	"errors"
	"net/http"
	"refine-portal/models"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetPropertyListRequest_Success(t *testing.T) {

	expected := &models.PropertyListResponse{
		Result: models.PropertyResult{
			ItemIDs: []string{"1", "2"},
		},
	}

	patches := gomonkey.ApplyFunc(
		GetURLFromConfig,
		func(string) (string, error) {
			return "http://example.com", nil
		},
	)
	defer patches.Reset()

	patches.ApplyFunc(
		NewGETRequest,
		func(string) (*http.Request, error) {
			return &http.Request{}, nil
		},
	)

	patches.ApplyFunc(
		DoRequest,
		func(_ *http.Request, target any) error {
			*(target.(*models.PropertyListResponse)) = *expected
			return nil
		},
	)

	req := models.PropertyListRequest{
		Category:  "dhaka",
		Locations: "BD",
		Order:     1,
		Limit:     20,
		Items:     1,
		Device:    "desktop",
		Page:      1,
	}

	result, err := GetPropertyListRequest(req)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestGetPropertyListRequest_DoRequestError(t *testing.T) {

	patches := gomonkey.ApplyFunc(
		GetURLFromConfig,
		func(string) (string, error) {
			return "http://example.com", nil
		},
	)
	defer patches.Reset()

	patches.ApplyFunc(
		NewGETRequest,
		func(string) (*http.Request, error) {
			return &http.Request{}, nil
		},
	)

	patches.ApplyFunc(
		DoRequest,
		func(*http.Request, any) error {
			return errors.New("request failed")
		},
	)

	req := models.PropertyListRequest{
		Category: "dhaka",
	}

	result, err := GetPropertyListRequest(req)

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "property list request failed")
}