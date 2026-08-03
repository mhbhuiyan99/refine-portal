# Unit Testing Documentation

## Overview

This document describes the testing strategy, testing tools, implementation decisions, and overall testing approach used in the Refine Portal project.

The objective of the testing process is to verify the correctness of the application's business logic while keeping tests fast, repeatable, and independent from external services.

---

# Testing Strategy

The project follows Go's standard unit testing practices.

The testing strategy includes:

- Writing tests in `_test.go` files.
- Using table-driven tests where applicable.
- Testing core business logic rather than framework internals.
- Isolating external dependencies.
- Covering both successful and error scenarios.
- Keeping tests independent and deterministic.

Different testing tools are selected depending on the type of dependency being tested instead of using one tool for every situation.

---

# Current Testing Tools

## Testify

Purpose:

- Assertions
- Easy comparison of expected and actual values
- Readable test failures

Example:

```go
assert.Equal(t, expected, actual)
assert.NoError(t, err)
assert.Error(t, err)
```


## httptest

Purpose:

Simulate an HTTP server during testing.

Instead of calling a real external API, a local HTTP server is created that returns predefined responses.

This allows testing of:

- HTTP request execution
- HTTP status validation
- JSON decoding
- Error handling

without requiring internet connectivity or real API credentials.


## gomonkey

Purpose:

Replace (patch) functions or methods during testing.

Instead of executing the real implementation, `gomonkey` temporarily redirects the call to a custom implementation defined by the test.

This is useful when a function depends on another function that should not be executed during unit testing.

Typical use cases:

- Mocking request layer functions.
- Avoiding external dependencies.
- Simulating success and error scenarios.
- Isolating business logic from lower layers.

---

# Test Coverage

## Services

### Functions Tested

- chunkStrings()
- GetLocation()
- GetProperties()
- GetPropertyDetails()
- GetPropertyImages()
- GetCategory()

### Testing Focus

#### `chunkStrings()`

Verified:

- Normal inputs.
- Empty slice.
- Invalid batch size.
- Boundary conditions.

Tool Used:

- Testify

---

#### `GetLocation()`

Verified:

- Request layer returns a successful response.
- Request layer returns an error.

Tool Used:

- gomonkey
- Testify

Why gomonkey?

`GetLocation()` depends on `requests.GetLocationRequest()`. During unit testing, the request layer is replaced with a fake implementation using `gomonkey`.

This isolates the service layer and ensures the test verifies only the service logic without making real HTTP requests.

---

#### `GetProperties()`

Verified:

- Returns the property list when the request layer succeeds.
- Returns the request layer error without modification.

Tool Used:

- gomonkey
- Testify

Why gomonkey?

`GetProperties()` depends on `requests.GetPropertyListRequest()`. During unit testing, the request layer is replaced with a fake implementation using `gomonkey`.

This allows the service logic to be tested independently without making real HTTP requests.

---

#### `GetPropertyDetails()`

Verified:

- Splits property IDs into batches.
- Retrieves property details from the request layer.
- Merges batch responses into a single result.
- Builds complete image URLs.
- Attaches partner feed information to each property.
- Returns immediately when any request batch fails.
- Returns an error if the image base URL configuration cannot be loaded.
- Splits more than 50 property IDs into multiple batches before calling the request layer.

Tool Used:

- gomonkey
- Testify

Why gomonkey?

`GetPropertyDetails()` depends on external request-layer functions to retrieve property details and configuration values. These dependencies are replaced using `gomonkey` so that the service logic can be tested independently without making external API calls or reading application configuration.

---

#### `GetPropertyImages()`

Verified:

- Calls the request layer with the correct property ID.
- Returns the image response from the request layer.
- Propagates errors returned by the request layer.

Tool Used:

- gomonkey
- Testify

Why gomonkey?

`GetPropertyImages()` depends on the request layer to retrieve property images. During testing, the request-layer function is replaced with a fake implementation so that the service can be verified independently without making external API calls.

---

#### `GetCategory()`

Verified:

- Calls the request layer with the correct parameters.
- Replaces location placeholders in section titles and subtitles.
- Builds complete image URLs using the configured image base URL.
- Propagates request-layer errors.
- Returns an error when image URL configuration cannot be loaded.

Tool Used:

- gomonkey
- Testify

Why gomonkey?

`GetCategory()` depends on external request-layer functions and application configuration. These dependencies are replaced using `gomonkey`, allowing the service's business logic (placeholder replacement and image URL construction) to be tested independently without external API calls or configuration files.

---

## Requests

Current tests include:

### URL helpers

- `BuildURL()`
- `BuildImageURL()`

Testing focuses on:

- Correct URL generation
- Query parameter encoding
- Invalid URLs
- Slash normalization

---

### HTTP Client

#### Function Tested

- `DoRequest()`

#### Purpose

The `DoRequest()` function is responsible for:

- Sending an HTTP request.
- Validating the HTTP status code.
- Decoding the JSON response into a target structure.
- Returning descriptive errors when a request or decoding fails.

#### Testing Tool

- `net/http/httptest`

#### Why `httptest`?

`DoRequest()` communicates with an external HTTP service. Instead of calling a real API, the tests create a local HTTP server using `httptest`.

This allows the entire request/response flow to be tested while remaining independent from external systems.

Benefits:

- Uses the real `http.Client`.
- Uses the real HTTP request/response flow.
- Uses the real JSON decoder.
- No internet connection required.
- No API credentials required.
- Fast and repeatable.

#### Implemented Test Scenarios

- Successful HTTP response with valid JSON.
- HTTP 500 Internal Server Error.
- Invalid JSON response.
- Network failure.

---

### Configuration Helper

#### Function Tested

- `GetURLFromConfig()`

#### Purpose

The `GetURLFromConfig()` function reads a URL value from the Beego application configuration and validates that the value exists and is not empty.

#### Testing Tool

- `gomonkey`

#### Why `gomonkey`?

`GetURLFromConfig()` depends on `web.AppConfig.String()`, which reads configuration from Beego's global configuration object.

Instead of requiring a real `app.conf` file during testing, `gomonkey` temporarily patches the `String()` method to return predefined values. This isolates the function from external configuration and keeps the unit tests fast and deterministic.

#### Implemented Test Scenarios

- Configuration value exists.
- Configuration key returns an error.
- Configuration value is empty or contains only whitespace.

---

### Location Request

#### Function Tested

- `GetLocationRequest()`

#### Purpose

The `GetLocationRequest()` function orchestrates the complete flow for retrieving location suggestions. It loads the API base URL, builds the request URL, creates an HTTP request, executes it, and returns the decoded response.

#### Testing Tool

- `gomonkey`
- `testify`

#### Why `gomonkey`?

The function depends on several helper functions within the request layer. These dependencies are patched so that the test focuses only on verifying the orchestration logic without making real HTTP requests or requiring application configuration.

#### Implemented Test Scenarios

- Successful request flow.
- Configuration read failure.

---


# How to Run Tests

Run all tests

```bash
go test ./...
```

Run a specific package

```bash
go test ./requests -v
```

Generate coverage

```bash
go test ./... -coverprofile=coverage.out
```

View coverage

```bash
go tool cover -func=coverage.out
```

---

# Current Coverage

Coverage will increase as additional unit tests are implemented for:

- Controllers
- Services
- Request layer
- Helper functions

---

# Notes

This project intentionally avoids calling real external APIs during unit tests.

External dependencies are isolated using appropriate testing tools, ensuring that tests remain:

- Fast
- Repeatable
- Independent
- Easy to maintain