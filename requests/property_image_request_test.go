package requests

import (
	"errors"
	"net/http"
	"refine-portal/models"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetPropertyImagesRequest_Success(t *testing.T) {
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

			assert.Contains(t, req.URL.String(), "propertyId=123")

			resp := target.(*models.PropertyImagesResponse)
			resp.Success = true
			resp.Images = []string{
				"image1.jpg",
				"image2.jpg",
			}

			return nil
		},
	)

	resp, err := GetPropertyImagesRequest("123")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.True(t, resp.Success)
	assert.Len(t, resp.Images, 2)
	assert.Equal(t, "image1.jpg", resp.Images[0])
}

func TestGetPropertyImagesRequest_ConfigError(t *testing.T) {
	patches := gomonkey.NewPatches()
	defer patches.Reset()

	patches.ApplyFunc(GetURLFromConfig,
		func(string) (string, error) {
			return "", errors.New("config error")
		},
	)

	resp, err := GetPropertyImagesRequest("123")

	assert.Nil(t, resp)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "config error")
}