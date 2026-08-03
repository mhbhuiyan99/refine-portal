package requests

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/beego/beego/v2/server/web"
	"github.com/stretchr/testify/assert"
)

func TestBuildImageURL(t *testing.T) {
	tests := []struct {
		name          string
		imageBaseURL  string
		imageName     string
		expectedImage string
	}{
		{
			name:          "normal url",
			imageBaseURL:  "https://images.example.com",
			imageName:     "room.jpg",
			expectedImage: "https://images.example.com/room.jpg",
		},
		{
			name:          "base url ends with slash",
			imageBaseURL:  "https://images.example.com/",
			imageName:     "room.jpg",
			expectedImage: "https://images.example.com/room.jpg",
		},
		{
			name:          "image starts with slash",
			imageBaseURL:  "https://images.example.com",
			imageName:     "/room.jpg",
			expectedImage: "https://images.example.com/room.jpg",
		},
		{
			name:          "both contain slash",
			imageBaseURL:  "https://images.example.com/",
			imageName:     "/room.jpg",
			expectedImage: "https://images.example.com/room.jpg",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := BuildImageURL(tc.imageBaseURL, tc.imageName)

			assert.Equal(t, tc.expectedImage, result)
		})
	}
}

func TestBuildURL(t *testing.T) {
	tests := []struct {
		name        string
		baseURL     string
		path        string
		query       url.Values
		expectedURL string
		expectError bool
	}{
		{
			name:    "url with query parameters",
			baseURL: "https://api.example.com",
			path:    "/api/location/v1",
			query: url.Values{
				"keyword": []string{"dhaka"},
				"page":    []string{"1"},
			},
			expectedURL: "https://api.example.com/api/location/v1?keyword=dhaka&page=1",
			expectError: false,
		},
		{
			name:        "url without query",
			baseURL:     "https://api.example.com",
			path:        "/api/location/v1",
			query:       nil,
			expectedURL: "https://api.example.com/api/location/v1",
			expectError: false,
		},
		{
			name:        "invalid base url",
			baseURL:     "://invalid-url",
			path:        "/api/location/v1",
			query:       nil,
			expectedURL: "",
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			result, err := BuildURL(
				tc.baseURL,
				tc.path,
				tc.query,
			)

			if tc.expectError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.expectedURL, result)
		})
	}
}

type testResponse struct {
	Name string `json:"name"`
}

func TestDoRequest_Success(t *testing.T) {
	// Create a fake HTTP server
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http. Request) {
			// Verify the request reached the server
			assert.Equal(t, http.MethodGet, r.Method)

			// Return a successful JSON response
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			_, err := w.Write([]byte(`{"name":"Alice"}`))
			assert.NoError(t, err)
		}),
	)

	defer server.Close()

	// Create a request to the fake server
	req, err := http.NewRequest(
		http.MethodGet,
		server.URL,
		nil,
	)

	assert.NoError(t, err)

	// Response will be decoded into this struct
	var result testResponse

	// Call the function we're testing
	err = DoRequest(req, &result)

	// Verify the result
	assert.NoError(t, err)
	assert.Equal(t, "Alice", result.Name)
}

func TestDoRequest_HTTPError(t *testing.T) {

	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.WriteHeader(http.StatusInternalServerError)

		}),
	)

	defer server.Close()

	req, err := http.NewRequest(
		http.MethodGet,
		server.URL,
		nil,
	)

	assert.NoError(t, err)

	var result testResponse

	err = DoRequest(req, &result)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unexpected HTTP status")
}

func TestDoRequest_InvalidJSON(t *testing.T) {

	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			_, err := w.Write([]byte(`{"name":`))
			assert.NoError(t, err)
		}),
	)

	defer server.Close()

	req, err := http.NewRequest(
		http.MethodGet,
		server.URL,
		nil,
	)
	assert.NoError(t, err)

	var result testResponse

	err = DoRequest(req, &result)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "decode response failed")
}

func TestDoRequest_RequestFailed(t *testing.T) {

	// Why port 1?
	// Normally, no application is listening on port 1.
	// When Go tries: 127.0.0.1:1 it immediately gets 'connection refused'
	// which causes httpClient.Do(req) to return an error.
	req, err := http.NewRequest(
		http.MethodGet,
		"http://127.0.0.1:1",
		nil,
	)
	assert.NoError(t, err)

	var result testResponse

	err = DoRequest(req, &result)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "request failed")
}

func TestNewGETRequest_Success(t *testing.T) {
	patches := gomonkey.ApplyMethod(
		reflect.TypeOf(web.AppConfig),
		"String",
		func(_ interface{}, key string) (string, error) {
			switch key {
			case "username":
				return "admin", nil
			case "password":
				return "secret", nil
			case "api_key":
				return "apikey123", nil
			default:
				return "", nil
			}
		},
	)
	defer patches.Reset()

	req, err := NewGETRequest("https://example.com")

	assert.NoError(t, err)
	assert.Equal(t, http.MethodGet, req.Method)
	assert.Equal(t, "application/json", req.Header.Get("Accept"))
	assert.Equal(t, "apikey123", req.Header.Get("x-api-key"))

	user, pass, ok := req.BasicAuth()
	assert.True(t, ok)
	assert.Equal(t, "admin", user)
	assert.Equal(t, "secret", pass)
}

func TestNewGETRequest_ConfigError(t *testing.T) {

	patches := gomonkey.ApplyMethod(
		reflect.TypeOf(web.AppConfig),
		"String",
		func(_ interface{}, key string) (string, error) {
			return "", errors.New("config error")
		},
	)
	defer patches.Reset()

	req, err := NewGETRequest("https://example.com")

	assert.Nil(t, req)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "set default headers")
}

func TestSetDefaultHeaders_Success(t *testing.T) {

	patches := gomonkey.ApplyMethod(
		reflect.TypeOf(web.AppConfig),
		"String",
		func(_ interface{}, key string) (string, error) {
			switch key {
			case "username":
				return "admin", nil
			case "password":
				return "secret", nil
			case "api_key":
				return "apikey123", nil
			default:
				return "", nil
			}
		},
	)
	defer patches.Reset()

	req, _ := http.NewRequest(http.MethodGet, "https://example.com", nil)

	err := setDefaultHeaders(req)

	assert.NoError(t, err)

	user, pass, ok := req.BasicAuth()
	assert.True(t, ok)
	assert.Equal(t, "admin", user)
	assert.Equal(t, "secret", pass)

	assert.Equal(t, "application/json", req.Header.Get("Accept"))
	assert.Equal(t, "application/json", req.Header.Get("Content-Type"))
	assert.Equal(t, "en-US", req.Header.Get("Accept-Language"))
	assert.Equal(t, "desktop", req.Header.Get("User-Agent"))
	assert.Equal(t, "XMLHttpRequest", req.Header.Get("X-Requested-With"))
	assert.Equal(t, "123presto-MS-ROW.com", req.Header.Get("Origin"))
	assert.Equal(t, "apikey123", req.Header.Get("x-api-key"))
}

func TestSetDefaultHeaders_ConfigError(t *testing.T) {

	patches := gomonkey.ApplyMethod(
		reflect.TypeOf(web.AppConfig),
		"String",
		func(_ interface{}, key string) (string, error) {
			return "", errors.New("config error")
		},
	)
	defer patches.Reset()

	req, _ := http.NewRequest(http.MethodGet, "https://example.com", nil)

	err := setDefaultHeaders(req)

	assert.Error(t, err)
}