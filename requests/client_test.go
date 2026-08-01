package requests

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

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
	Name string `json: "name"`
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