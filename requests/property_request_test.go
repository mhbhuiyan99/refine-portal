package requests

import (
	"errors"
	"net/http"
	"refine-portal/models"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
)


func TestGetPropertyDetailsRequest_Success(t *testing.T) {
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

			assert.Contains(t, req.URL.String(), "propertyIdList=1%2C2")

			resp := target.(*models.PropertyDetailsResponse)
			resp.Success = true
			resp.Items = []models.PropertyDetails{
				{
					ID: "1",
				},
			}

			return nil
		},
	)

	resp, err := GetPropertyDetailsRequest([]string{"1", "2"})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.True(t, resp.Success)
	assert.Len(t, resp.Items, 1)
	assert.Equal(t, "1", resp.Items[0].ID)
}

func TestGetPropertyDetailsRequest_ConfigError(t *testing.T) {
	patches := gomonkey.NewPatches()
	defer patches.Reset()

	patches.ApplyFunc(GetURLFromConfig,
		func(string) (string, error) {
			return "", errors.New("config error")
		},
	)

	resp, err := GetPropertyDetailsRequest([]string{"1", "2"})

	assert.Nil(t, resp)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "config error")
}

