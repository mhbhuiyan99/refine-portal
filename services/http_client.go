package services

import (
	"fmt"
	"net/http"
	"time"

	"github.com/beego/beego/v2/server/web"
)

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

func SetDefaultHeaders(request *http.Request) error {
	username, err := web.AppConfig.String("username")
	if err != nil {
		return err
	}

	password, err := web.AppConfig.String("password")
	if err != nil {
		return err
	}

	apiKey, err := web.AppConfig.String("api_key")
	if err != nil {
		return err
	}

	request.SetBasicAuth(username, password)

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Accept-Language", "en-US")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-Requested-With", "XMLHttpRequest")
	request.Header.Set("User-Agent", "desktop")
	request.Header.Set("Origin", "123presto-MS-ROW.com")
	request.Header.Set("x-api-key", apiKey)

	return nil
}

func NewGETRequest(url string) (*http.Request, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}

	if err := SetDefaultHeaders(request); err != nil {
		return nil, fmt.Errorf("set default headers: %w", err)
	}

	return request, nil
}