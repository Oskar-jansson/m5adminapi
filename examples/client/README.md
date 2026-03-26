# Client

The `Client` is the core HTTP client wrapper for the M5 Admin API SDK. It handles HTTP requests, authentication, TLS configuration, and provides middleware support for custom request processing.

## Overview

The `Client` struct encapsulates all the necessary configuration for making HTTP requests to the M5 Admin API, including:

- Server address and API path
- Authentication credentials and access tokens
- TLS configuration for secure connections
- Middleware hooks for request/response processing
- Retry logic and timeout handling

## Creating a Client

### Using the Constructor

```go
// Create a new client instance
client := m5adminapi.NewClient()
```

### Configuration Methods

The client provides fluent setter methods for configuration:

```go
client := m5adminapi.NewClient().
    SetAddress("https://api.example.com").
    SetPath("/m5adminapi/api").
    SetCredentials(models.Credentials{
        Id:       "1",
        System:   "RASYSTEM",
        Lang:     "en",
        User:     "operator",
        Password: "password",
        Apitype:  "main",
        Apikey:   "your-api-key",
    }).
    SetTlsConfig(&tls.Config{
        // tls settings here
    })
```

## Client Fields

### Core Configuration

- `Address` (string): The base URL of the API server (e.g., "http://localhost" or "https://api.example.com")
- `Path` (string): The API path prefix (e.g., "/m5adminapi/api")
- `Credentials` (models.Credentials): Authentication credentials for API access
- `AccessToken` (string): OAuth-style access token for authenticated requests
- `TLSConfig` (*tls.Config): TLS configuration for HTTPS connections

### Middleware

- `PreFlightValidation` ([]func(*RequestContext) (ShouldRetry, error)): Functions called before each request for validation or modification
- `PostFlightRetryTrigger` ([]func(*RequestContext, *HttpResponse, error) (ShouldRetry, error)): Functions called after each request to determine retry behavior

## Key Methods

### Set

- `SetAddress(v string) *Client`: Set the server address
- `SetPath(v string) *Client`: Set the API path
- `SetCredentials(v models.Credentials) *Client`: Set authentication credentials
- `SetTlsConfig(v *tls.Config) *Client`: Set TLS configuration
- `SetAccessToken(token string)`: Set the access token
- `AddPreflightEventTrigger(fn func(*RequestContext) (ShouldRetry, error)) *Client`: Add pre-flight validation function
- `AddPostflightEventTrigger(fn func(*RequestContext, *HttpResponse, error) (ShouldRetry, error)) *Client`: Add post-flight retry trigger

### Get

- `GetAddress() string`: Get the current server address
- `GetPath() string`: Get the current API path
- `GetCredentials() models.Credentials`: Get the current credentials
- `GetTlsConfig() *tls.Config`: Get the current TLS configuration
- `GetAccessToken() string`: Get the current access token

## Usage with AdminApiConnection

## Middleware

The client supports middleware functions for custom request processing:

### Pre-flight Validation

Called before each request, allowing validation and cancelling of said request.

```go
client.AddPreflightEventTrigger(func(ctx *RequestContext) (ShouldRetry, error) {

    if ctx.Method == "POST" {
        // Custom validation logic
        // Telemtry, etc...
    }
    return false, nil // ShouldRetry: false, error: nil
})
```

### Post-flight Retry Triggers

Called after each request, can determine if retry is needed. Does not overwrite `retrySettings` provided.

```go
client.AddPostflightEventTrigger(func(ctx *RequestContext, resp *HttpResponse, err error) (ShouldRetry, error) {
    if resp != nil && resp.StatusCode == 429 { // Rate limited
        return true, nil // Retry
    }
    return false, nil // Don't retry
})
```

## Thread Safety

The `Client` uses a `sync.Mutex` to protect shared state, making it safe for concurrent use. However, middleware functions should be added during initialization to avoid race conditions.
