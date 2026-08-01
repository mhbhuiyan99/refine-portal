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

# Why httptest was chosen for DoRequest()

The `DoRequest()` function is responsible for:

- Sending HTTP requests
- Receiving HTTP responses
- Validating status codes
- Decoding JSON responses

Since its responsibility is HTTP communication, `httptest` is the most suitable testing tool.

It provides a real local HTTP server, allowing the complete request/response flow to be tested without depending on external services.

Compared with interface-based mocking frameworks or function patching tools, `httptest` exercises more production code while keeping tests isolated and repeatable.

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