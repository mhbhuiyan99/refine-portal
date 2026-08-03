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

Different testing tools are selected based on the type of dependency being tested.

---

# Testing Tools

## Testify

Purpose:

- Assertions
- Simplifies comparison of expected and actual values
- Readable test failures

Example:

```go
assert.Equal(t, expected, actual)
assert.NoError(t, err)
assert.Error(t, err)
```


## httptest

Purpose:

Simulates an HTTP server during testing.

Instead of calling a real external API, a local HTTP server is created that returns predefined responses.

This allows testing of:

- HTTP request execution
- HTTP status validation
- JSON decoding
- Error handling

without requiring internet connectivity or real API credentials.

## Function Variable Injection

Purpose:

Replace selected dependencies during testing without relying on runtime method patching.

Package-level function variables allow tests to substitute implementations with lightweight test doubles while keeping the production code unchanged.

Some functions depend on application configuration (for example, reading values from `app.conf`). Instead of patching framework methods, these dependencies are exposed through package-level function variables.

Example:

```go
var getConfig = GetURLFromConfig
```
This approach follows the dependency injection principle and makes configuration-dependent code easier to unit test without modifying production behavior.

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

`gomonkey` is used when testing code that directly depends on functions in other packages, such as request-layer or service-layer functions.

It allows these dependencies to be replaced with predefined implementations so that the function under test can be executed in isolation.

Typical examples include:

- Mocking request-layer functions from the service layer.
- Mocking service-layer functions from controllers.
- Simulating success and failure scenarios.

Configuration access is handled through function variable injection instead of `gomonkey`. This avoids runtime method patching for configuration reads, simplifies the tests, and reduces dependence on framework internals.

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

The request layer currently includes tests for:

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

`NewGETRequest()` and `setDefaultHeaders()` are tested using function variable injection to replace configuration access during unit testing.


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

- `testify`
- Function variable injection

These tests replace the configuration reader through function variable injection, allowing request construction and header configuration to be verified without requiring a real application configuration file.

---

### Configuration Helper

#### Function Tested

- `GetURLFromConfig()`


#### Testing Tool

- `gomonkey` (used for mocking `web.AppConfig.String()`)


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

- CategoryController.Get()
- LocationAPIController.Get()
- PropertyAPIController.GetList()
- PropertyAPIController.GetDetails()
- PropertyImageController.Get()
- RefineController.Get()

#### Purpose

The controller validates incoming HTTP requests, invokes the service layer, and returns an appropriate HTTP response.

#### Testing Tool

- `httptest`
- `gomonkey`
- `testify`

#### Why `httptest`?

Controllers are HTTP handlers. `httptest` creates an in-memory HTTP request and response recorder, allowing the controller behavior to be tested without starting a real web server.

#### Why `gomonkey`?

The controller depends on the service layer. During unit testing, service-layer functions such as `services.GetLocation()`, `services.GetProperties()`, `services.GetPropertyDetails()`, and `services.GetPropertyImages()` are patched using `gomonkey`. This allows each controller to be tested independently of the service implementation.

#### Implemented Test Scenarios

The controller tests verify:

- Successful request handling.
- Validation of required query parameters.
- Default values for optional query parameters.
- Proper invocation of the service layer.
- Correct JSON responses.
- Appropriate HTTP status codes.
- Proper handling of service-layer errors.

### RefineController

- Returns the provided search and sorting parameters.
- Uses default values when query parameters are not provided.
- Populates the template data correctly.
- Sets the expected template name.

### PropertyImageController

- Returns the property images when a valid property ID is provided.
- Returns **HTTP 400 Bad Request** when the required `propertyId` parameter is missing.
- Returns **HTTP 500 Internal Server Error** when the service layer returns an error.

### PropertyAPIController

#### GetList()

Verified:

- Returns the property list successfully.
- Validates the required `category` parameter.
- Validates the required `location` parameter.
- Uses default values for optional query parameters when omitted.
- Returns HTTP 500 when the service layer returns an error.

#### GetDetails()

Verified:

- Returns property details successfully.
- Validates the required `propertyIdList` parameter.
- Splits the property ID list correctly before passing it to the service layer.
- Returns HTTP 500 when the service layer returns an error.

### CategoryController

Verified:

- Extracts the category slug from the request URL.
- Converts the URL slug into the format required by the Category API.
- Resolves the country code using the Location service.
- Retrieves category data successfully.
- Populates the template data.
- Sets the expected template name.
- Returns HTTP 500 when the Location service returns an error.
- Returns HTTP 500 when the Category service returns an error.

#### Note

Beego's `CustomAbort()` intentionally triggers a panic after writing the HTTP response. The unit tests use `assert.Panics()` to verify this expected framework behavior while still validating the returned HTTP status code and response body.

## Notes for Controller Tests

Controller tests use `gomonkey` to replace service-layer functions.

Because `gomonkey` performs runtime function patching, Go compiler optimizations such as function inlining may prevent patches from being applied correctly.

When running tests that depend on `gomonkey`, use:

```bash
go test ./controllers -gcflags=all=-l -v
```

---

# How to Run Tests

Run all tests

```bash
go test ./... -gcflags=all=-l
```

Run a specific package

```bash
go test ./requests -gcflags=all=-l -v
```

Generate coverage

```bash
go test ./... -gcflags=all=-l -coverprofile=coverage.out
```

View coverage

```bash
go tool cover -func=coverage.out
```

---

# Current Coverage

The current implementation focuses on the project's primary controllers, services, request layer, and helper functions. Additional packages can be covered using the same testing approach as the project grows.

---

# Notes

This project intentionally avoids calling real external APIs during unit tests.

External dependencies are isolated using appropriate testing tools, ensuring that tests remain:

- Fast
- Repeatable
- Independent
- Easy to maintain and extend