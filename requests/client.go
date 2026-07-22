// Package requests contains all external API communication logic.
// It is responsible for creating HTTP requests, executing them,
// decoding responses, and returning typed models.
//
// Controllers and services should not communicate with external APIs
// directly. All API interactions should go through this package.
package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

// DoRequest sends an HTTP request and decodes the JSON response.
//
// Responsibilities:
//   - Execute the HTTP request.
//   - Validate the HTTP status code.
//   - Decode the JSON response into target.
//   - Return descriptive errors for request, status, or decode failures.
func DoRequest(
    req *http.Request,
    target any,
) error {
	resp, err := httpClient.Do(req)
	if err != nil {
		logs.Error(
			"[RequestLayer] HTTP request failed | url=%s | err=%v",
			req.URL.String(),
			err,
		)

		return fmt.Errorf("request failed: %w", err)
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(target); err != nil {
		logs.Error(
			"[RequestLayer] Decode failed | err=%v",
			err,
		)

		return fmt.Errorf(
			"decode response failed: %w",
			err,
		)
	}

	return nil
}


func setDefaultHeaders(request *http.Request) error {
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

// NewGETRequest creates an HTTP GET request with the
// application's default authentication and headers.
func NewGETRequest(url string) (*http.Request, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}

	if err := setDefaultHeaders(request); err != nil {
		return nil, fmt.Errorf("set default headers: %w", err)
	}

	return request, nil
}