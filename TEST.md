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

---

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

---

# Test Coverage

## Services

Current tests include:

- Helper functions
    - `chunkStrings()`

Testing focuses on:

- Normal inputs
- Empty input
- Boundary conditions
- Invalid parameters

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

#### Planned Test Scenarios

- Network failure.

- `DoRequest()`

Current scenarios:

- Successful HTTP response with valid JSON

Additional scenarios planned:

- HTTP error responses
- Invalid JSON
- Network failures

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