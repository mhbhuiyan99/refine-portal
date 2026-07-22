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
	"net/url"
	"strings"
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

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		logs.Error(
			"[RequestLayer] Unexpected HTTP status | url=%s | status=%d",
			req.URL.String(),
			resp.StatusCode,
		)
		return fmt.Errorf("unexpected HTTP status: %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(target); err != nil {
		logs.Error(
			"[RequestLayer] Decode failed | url=%s | err=%v",
			req.URL.String(),
			err,
		)

		return fmt.Errorf(
			"decode response failed: %w",
			err,
		)
	}

	return nil
}

// BuildURL constructs a complete request URL.
//
// Responsibilities:
//   - Parse the base URL.
//   - Append the API path.
//   - Encode query parameters.
//   - Return the final URL string.
func BuildURL(baseURL string, path string, queryParams url.Values) (string, error) {
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("parse base_url failed: %w", err)
	}

	parsedURL.Path = path

	if queryParams != nil {
		parsedURL.RawQuery = queryParams.Encode()
	}

	return parsedURL.String(), nil
}

// GetBaseURL retrieves the configured API base URL.
//
// Responsibilities:
//   - Read the base URL from the application configuration.
//   - Validate that the base URL is not empty.
//   - Return the configured base URL.
func GetBaseURL() (string, error) {

    baseURL, err := web.AppConfig.String("base_url")
    if err != nil {
        return "", fmt.Errorf("failed to read base_url: %w", err)
    }

    if strings.TrimSpace(baseURL) == "" {
        return "", fmt.Errorf("base_url is empty")
    }

    return baseURL, nil
}

// NewGETRequest creates a configured HTTP GET request.
//
// Responsibilities:
//   - Create a new HTTP GET request.
//   - Apply the application's default authentication.
//   - Apply the application's default request headers.
//   - Return the configured request.
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

// setDefaultHeaders applies the application's default
// authentication and HTTP headers to a request.
//
// Responsibilities:
//   - Apply Basic Authentication.
//   - Set common request headers.
//   - Set the API key.
//   - Prepare the request for external API communication.
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
