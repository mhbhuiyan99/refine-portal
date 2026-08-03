package requests

import (
	"errors"
	"net/http"
	"refine-portal/models"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
)


func TestGetLocationRequest_Success(t *testing.T) {
	patches := gomonkey.NewPatches()
	defer patches.Reset()

	patches.ApplyFunc(GetURLFromConfig,
		func(string) (string, error) {
			return "https://example.com", nil
		},
	)

	patches.ApplyFunc(NewGETRequest,
		func(url string) (*http.Request, error) {
			return http.NewRequest(http.MethodGet, url, nil)
		},
	)

	patches.ApplyFunc(DoRequest,
		func(req *http.Request, target any) error {
			resp := target.(*models.LocationResponse)

			resp.Success = true
			resp.GeoInfo.Name = "Dhaka"

			assert.Contains(t, req.URL.String(), "keyword=dhaka")

			return nil
		},
	)

	resp, err := GetLocationRequest("dhaka")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.True(t, resp.Success)
	assert.Equal(t, "Dhaka", resp.GeoInfo.Name)
}

func TestGetLocationRequest_ConfigError(t *testing.T) {
	patches := gomonkey.NewPatches()
	defer patches.Reset()

	patches.ApplyFunc(GetURLFromConfig,
		func(string) (string, error) {
			return "", errors.New("config error")
		},
	)

	resp, err := GetLocationRequest("dhaka")

	assert.Nil(t, resp)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "config error")
}