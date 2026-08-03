package requests

import (
	"errors"
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/beego/beego/v2/server/web"
	"github.com/stretchr/testify/assert"
)

func TestGetURLFromConfig_Success(t *testing.T) {
	patches := gomonkey.ApplyMethod(
		reflect.TypeOf(web.AppConfig),
		"String",
		func(_ interface{}, key string) (string, error) {
			return "https://example.com", nil
		},
	)
	defer patches.Reset()

	url, err := GetURLFromConfig("base_url")

	assert.NoError(t, err)
	assert.Equal(t, "https://example.com", url)
}

func TestGetURLFromConfig_ReadError(t *testing.T) {
	patches := gomonkey.ApplyMethod(
		reflect.TypeOf(web.AppConfig),
		"String",
		func(_ interface{}, key string) (string, error) {
			return "", errors.New("config not found")
		},
	)
	defer patches.Reset()

	url, err := GetURLFromConfig("base_url")

	assert.Error(t, err)
	assert.Empty(t, url)
	assert.Contains(t, err.Error(), "failed to read")
}

func TestGetURLFromConfig_EmptyValue(t *testing.T) {
	patches := gomonkey.ApplyMethod(
		reflect.TypeOf(web.AppConfig),
		"String",
		func(_ interface{}, key string) (string, error) {
			return "   ", nil
		},
	)
	defer patches.Reset()

	url, err := GetURLFromConfig("base_url")

	assert.Error(t, err)
	assert.Empty(t, url)
	assert.Contains(t, err.Error(), "base_url is empty")
}