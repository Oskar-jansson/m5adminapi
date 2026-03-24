package client

import (
	"crypto/tls"
	"sync"
	"time"

	"github.com/Oskar-jansson/m5adminapi/models"
)

// Root object of http requests wrapper
type Client struct {

	// Destination of request. ex http://localhost
	// Must contain protocol (http:// or https://)
	Address string

	// Path from address. ex /m5adminApi/api
	Path string

	// store credentials with client
	Credentials models.Credentials

	AccessToken string
	TLSConfig   *tls.Config

	mu sync.Mutex

	// Injected logic middleware, allows to alter behavior of http wrapper.
	// Useful for injecting custom error handling or input validation
	PreFlightValidation    []func(*RequestContext) (ShouldRetry, error)
	PostFlightRetryTrigger []func(*RequestContext, *HttpResponse, error) (ShouldRetry, error)
}

// named bool for readability
type ShouldRetry bool

type ParseStrategy string

const (
	ParseStrategyLenient ParseStrategy = "lenient"
	ParseStrategyStrict  ParseStrategy = "strict"
)

func (p ParseStrategy) IsValid() bool {
	return p == ParseStrategyLenient || p == ParseStrategyStrict
}

// RequestContext contains all information about the current request
// Made available to PreFlight and PostFlight middleware
// Output from Hooks
type RequestContext struct {
	Method        string
	URL           string
	Body          []byte
	Headers       map[string]string
	AttemptNumber int
	MaxRetries    int
	StartTime     time.Time
	RetrySettings RetrySettings
}

// RequestSettings holds configuration for an HTTP request
// Input only
type RequestSettings struct {
	URL                string // Target URL (will be appended to Client.Address)
	IncludeAccessToken bool
	Method             string        // HTTP method (GET, POST, PATCH, DELETE, etc.)
	Body               []byte        // Request body
	Retry              RetrySettings // Retry configuration
	TimeOut            time.Duration // Request timeout
}

// RetrySettings configures retry behavior for HTTP requests
type RetrySettings struct {
	MaxRetries       int           // Maximum number of retry attempts
	RetryDelay       time.Duration // Base delay between retries
	ExponentialRetry bool          // Use exponential backoff
}

// HttpResponse represents an HTTP response
type HttpResponse struct {
	StatusCode int           // HTTP status code
	Body       string        // Response body
	Duration   time.Duration // Request duration
}
