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

# Testing Tools

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

### Why gomonkey?

The service and request layers depend on lower-level functions such as configuration readers, HTTP request builders, and external API calls.

During unit testing, these dependencies are replaced with controlled implementations using `gomonkey`. This allows each function to be tested in isolation without requiring:

- External HTTP services
- Application configuration files
- Network connectivity

As a result, the tests remain fast, deterministic, and focused only on the business logic of the function under test.

---

# Current Test Coverage

Current unit tests cover:

- Controllers
- Services
- Request layer
- Helper functions

Additional controllers and remaining packages can be tested following the same testing approach.

## Services

Overall Coverage: 97.6%

| Function | Coverage |
|----------|---------:|
| chunkStrings() | 100% |
| GetLocation() | 100% |
| GetProperties() | 100% |
| GetPropertyDetails() | 97.6% |
| GetPropertyImages() | 100% |
| GetCategory() | 95.0% |

### Testing Tools

- Testify
- gomonkey

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


---

#### `GetLocation()`

Verified:

- Request layer returns a successful response.
- Request layer returns an error.


---

#### `GetProperties()`

Verified:

- Returns the property list when the request layer succeeds.
- Returns the request layer error without modification.

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

---

#### `GetPropertyImages()`

Verified:

- Calls the request layer with the correct property ID.
- Returns the image response from the request layer.
- Propagates errors returned by the request layer.

---

#### `GetCategory()`

Verified:

- Calls the request layer with the correct parameters.
- Replaces location placeholders in section titles and subtitles.
- Builds complete image URLs using the configured image base URL.
- Propagates request-layer errors.
- Returns an error when image URL configuration cannot be loaded.

---

## Requests

| Function | Coverage |
|----------|---------:|
| DoRequest() | 100% |
| BuildURL() | 100% |
| BuildImageURL() | 100% |
| GetURLFromConfig() | 100% |
| GetCategoryRequest() | 100% |
| GetPropertyListRequest() | 87.0% |
| GetPropertyDetailsRequest() | 83.3% |
| GetPropertyImagesRequest() | 81.2% |
| GetLocationRequest() | 80.0% |
| NewGETRequest() | 83.3% |
| setDefaultHeaders() | 88.9% |

### Request Functions

Each request function is responsible for:

- Reading configuration.
- Building the request URL.
- Creating an HTTP request.
- Executing the request.
- Returning the decoded response.

---

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

#### Functions Tested

- `DoRequest()`
- `NewGETRequest()`
- `setDefaultHeaders()`


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

#### Additional Request Construction Tests

The request helper functions were also tested.

`NewGETRequest()`

Verified:

- Creates an HTTP GET request.
- Applies the application's default headers.
- Returns an error when default headers cannot be applied.

`setDefaultHeaders()`

Verified:

- Applies Basic Authentication.
- Sets all required HTTP headers.
- Returns an error when configuration values cannot be loaded.

Testing Tool:

- `gomonkey`
- `testify`

These tests verify request construction and header configuration without requiring a real application configuration file.

---

### Configuration Helper

#### Function Tested

- `GetURLFromConfig()`


#### Testing Tool

- `gomonkey`


#### Implemented Test Scenarios

- Configuration value exists.
- Configuration key returns an error.
- Configuration value is empty or contains only whitespace.

---

### Location Request

#### Function Tested

- `GetLocationRequest()`

#### Testing Tool

- `gomonkey`
- `testify`

#### Implemented Test Scenarios

- Successful request flow.
- Configuration read failure.

---

### Property Details Request

#### Function Tested

- `GetPropertyDetailsRequest()`

#### Testing Tool

- `gomonkey`
- `testify`


#### Implemented Test Scenarios

- Successful request flow.
- Configuration read failure.

---

### Property Images Request

#### Function Tested

- `GetPropertyImagesRequest()`


#### Testing Tool

- `gomonkey`
- `testify`


#### Implemented Test Scenarios

- Successful request flow.
- Configuration read failure.

---

### Property List Request

#### Function Tested

- `GetPropertyListRequest()`

#### Testing Tool

- `gomonkey`
- `testify`


#### Implemented Test Scenarios

- Successful property list retrieval.
- Configuration read failure.
- URL construction failure.
- Request creation failure.
- HTTP request execution failure.

#### Why not `httptest`?

`httptest` is most appropriate when testing actual HTTP communication.

In this case, the HTTP behavior has already been tested inside `DoRequest()`. Using `httptest` here would repeat the same HTTP flow and make the test slower without increasing confidence.

Using `gomonkey` keeps this test focused on the request orchestration logic while remaining fast and independent.

---

### Category Request

#### Function Tested

- `GetCategoryRequest()`

#### Testing Tool

- `gomonkey`
- `testify`


#### Implemented Test Scenarios

- Successful category retrieval.
- Configuration read failure.
- URL construction failure.
- HTTP request creation failure.
- HTTP request execution failure.

---

## Controllers

### Functions Tested

- `LocationAPIController.Get()`
- `RefineController.Get()`
- `PropertyImageController.Get()`

#### Purpose

The controller validates incoming HTTP requests, invokes the service layer, and returns an appropriate HTTP response.

#### Testing Tool

- `httptest`
- `gomonkey`
- `testify`

#### Why `httptest`?

Controllers are HTTP handlers. `httptest` creates an in-memory HTTP request and response recorder, allowing the controller behavior to be tested without starting a real web server.

#### Why `gomonkey`?

The controller depends on the service layer. During unit testing, `services.GetLocation()` is patched using `gomonkey` so that the controller can be tested independently of the service implementation.

#### Implemented Test Scenarios

- Successful request.
- Missing required query parameter.
- Service returns an error.

### RefineController

- Returns the provided search and sorting parameters.
- Uses default values when query parameters are not provided.
- Populates the template data correctly.
- Sets the expected template name.

### PropertyImageController

- Returns the property images when a valid property ID is provided.
- Returns **HTTP 400 Bad Request** when the required `propertyId` parameter is missing.
- Returns **HTTP 500 Internal Server Error** when the service layer returns an error.

#### Note

Beego's `CustomAbort()` intentionally triggers a panic after writing the HTTP response. The unit tests use `assert.Panics()` to verify this expected framework behavior while still validating the returned HTTP status code and response body.

## Notes for Controller Tests

Controller tests use `gomonkey` to replace service-layer functions.

Some service functions are very small wrappers and may be inlined by the Go compiler. Function inlining can prevent `gomonkey` from applying patches correctly.

When running controller tests that depend on function patching, the following command may be required:

```bash
go test ./controllers -gcflags=all=-l -v
```

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
- Easy to maintain and extend