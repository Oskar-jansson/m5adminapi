package client

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/Oskar-jansson/m5adminapi/models"

	// validates schema
	val "github.com/go-playground/validator/v10"
)

var (

	// Enum.
	// Dictates usage of strict parsing or a more lenient version
	parseStrategyStorage ParseStrategy // "strict", "lenient". default is "lenient"
	mu                   sync.Mutex
)

func init() {
	parseStrategyStorage = ParseStrategyLenient
}

func SetParseStrategy(v ParseStrategy) error {
	if !v.IsValid() {
		return fmt.Errorf("invalid parse strategy: %q", v)
	}

	mu.Lock()
	defer mu.Unlock()
	parseStrategyStorage = v
	return nil
}

func GetParseStrategy() ParseStrategy {
	mu.Lock()
	defer mu.Unlock()
	return parseStrategyStorage
}

func (c *Client) SetAccessToken(token string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.AccessToken = token
}

func (c *Client) GetAccessToken() string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.AccessToken
}

func (c *Client) SetAddress(v string) *Client {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Address = v
	return c
}

func (c *Client) GetAddress() string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.Address
}

func (c *Client) SetPath(v string) *Client {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Path = v
	return c
}

func (c *Client) GetPath() string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.Path
}

func (c *Client) SetCredentials(v models.Credentials) *Client {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Credentials = v
	return c
}

func (c *Client) GetCredentials() models.Credentials {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.Credentials
}

func (c *Client) SetTlsConfig(v *tls.Config) *Client {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.TLSConfig = v
	return c
}

func (c *Client) GetTlsConfig() *tls.Config {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.TLSConfig
}

func (c *Client) AddPreflightEventTrigger(fn func(*RequestContext) (ShouldRetry, error)) *Client {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.PreFlightValidation = append(c.PreFlightValidation, fn)
	return c
}

func (c *Client) AddPostflightEventTrigger(fn func(*RequestContext, *HttpResponse, error) (ShouldRetry, error)) *Client {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.PostFlightRetryTrigger = append(c.PostFlightRetryTrigger, fn)
	return c
}

func (c *Client) HttpRequest(ctx context.Context, rs RequestSettings) (*HttpResponse, error) {
	if rs.URL == "" {
		return nil, fmt.Errorf("URL cannot be empty")
	}

	var finalErr error

	httpClient := c.newHTTPClient(rs.TimeOut)

	fullUrl, err := assembleUrl(c.Address, rs.URL)
	if err != nil {
		return nil, err
	}

	// Object returned by hooks
	reqCtx := &RequestContext{
		Method:        rs.Method,
		URL:           fullUrl,
		Body:          rs.Body,
		Headers:       make(map[string]string),
		MaxRetries:    rs.Retry.MaxRetries,
		StartTime:     time.Now(),
		RetrySettings: rs.Retry,
	}

	if rs.IncludeAccessToken {
		reqCtx.Headers["Access-Token"] = c.GetAccessToken()
	}

	// Retry loop
retryLoop:
	for attempt := 0; attempt <= rs.Retry.MaxRetries; attempt++ {

		var shouldRetry ShouldRetry
		var shouldError error

		if ctx.Err() != nil {
			return nil, ctx.Err()
		}

		reqCtx.AttemptNumber = attempt + 1
		attemptStart := time.Now()

		for _, preTrigger := range c.PreFlightValidation {
			requestingRetry, triggerErr := preTrigger(reqCtx)

			if triggerErr != nil {
				shouldError = triggerErr
			}
			if requestingRetry {
				shouldRetry = requestingRetry
			}
		}
		if shouldError != nil {
			return nil, shouldError
		}
		if shouldRetry {
			if attempt < rs.Retry.MaxRetries {

				c.waitBeforeRetry(attempt, rs.Retry)
				continue retryLoop

			}

			// Max retries reached but custom logic still wants to retry
			return nil, fmt.Errorf("max retries reached, custom retry logic requested additional retry")

		}

		resp, err := c.buildAndExecuteRequest(ctx, httpClient, fullUrl, rs)
		if resp != nil {
			resp.Duration = time.Since(attemptStart)
		}

		// PostFlightRetryTrigger
		for _, postTrigger := range c.PostFlightRetryTrigger {
			requestingRetry, triggerErr := postTrigger(reqCtx, resp, err)

			if triggerErr != nil {
				shouldError = triggerErr
			}
			if requestingRetry {
				shouldRetry = requestingRetry
			}
		}
		if shouldError != nil {
			return resp, shouldError
		}
		if shouldRetry {
			if attempt < rs.Retry.MaxRetries {

				c.waitBeforeRetry(attempt, rs.Retry)
				continue retryLoop

			}

			// Max retries reached but custom logic still wants to retry
			return nil, fmt.Errorf("max retries reached, custom retry logic requested additional retry")

		}

		if err != nil {
			finalErr = err
			if attempt < rs.Retry.MaxRetries {
				c.waitBeforeRetry(attempt, rs.Retry)
				continue
			}
			break
		}

		return resp, nil
	}

	return nil, finalErr
}

// newHTTPClient creates and configures a temporary net/http.Client.
// It does not configure the wrapper Client itself.
func (c *Client) newHTTPClient(timeout time.Duration) *http.Client {
	transport := &http.Transport{}
	if c.TLSConfig != nil {
		transport.TLSClientConfig = c.TLSConfig
	}

	httpClient := &http.Client{Transport: transport}
	if timeout > 0 {
		httpClient.Timeout = timeout
	}

	return httpClient
}

// buildAndExecuteRequest build and performs a single HTTP request attempt
func (c *Client) buildAndExecuteRequest(ctx context.Context, client *http.Client, fullUrl string, rs RequestSettings) (*HttpResponse, error) {

	req, err := http.NewRequestWithContext(ctx, rs.Method, fullUrl, bytes.NewReader(rs.Body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	if rs.IncludeAccessToken {
		req.Header.Set("Access-Token", c.GetAccessToken())
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request execution failed: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return &HttpResponse{
		StatusCode: resp.StatusCode,
		Body:       string(bodyBytes),
	}, nil
}

// returns an exponential retry timeout
func (c *Client) waitBeforeRetry(attempt int, retry RetrySettings) {

	// Helper function
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	if retry.ExponentialRetry {
		exponent := min(attempt, 10)
		delay := time.Duration(1<<exponent) * time.Second
		time.Sleep(delay)
	} else {
		time.Sleep(retry.RetryDelay)
	}
}

// ResponseConvert converts an HttpResponse into either the expected type T or an SdkError
// Prioritizes error on bad statuses.
func ResponseConvert[T any](r *HttpResponse) (*T, *models.SdkError) {
	raw := []byte(r.Body)

	if len(raw) == 0 {
		return nil, nil
	}

	// Handle error status codes
	// Does not determine error, just prefered parse method.
	if r.StatusCode < 200 || r.StatusCode >= 300 {
		var em models.ErrorMessage

		err := json.Unmarshal(raw, &em)
		if err != nil {

			return nil, &models.SdkError{
				Err: err,
			}
		}
		return nil, &models.SdkError{
			ErrorMessage: &em,
		}

	}

	var obj T

	dec := json.NewDecoder(bytes.NewReader(raw))

	if GetParseStrategy() == ParseStrategyStrict {
		dec.DisallowUnknownFields()
	}

	if err := dec.Decode(&obj); err != nil {
		return nil, &models.SdkError{Err: fmt.Errorf("decode error: %w", err)}
	}

	if dec.More() {
		return nil, &models.SdkError{Err: fmt.Errorf("unexpected trailing data")}
	}

	if GetParseStrategy() == ParseStrategyStrict {
		validate := val.New()
		if err := validate.Struct(obj); err != nil {
			return nil, &models.SdkError{Err: fmt.Errorf("validation error: %w", err)}
		}
	}

	return &obj, nil
}

// ResponseToSdkError converts an HttpResponse an SdkError.
//
// Used when no body is expected in response.
func ResponseToSdkError(r *HttpResponse) *models.SdkError {
	raw := []byte(r.Body)

	if len(raw) == 0 {
		return nil
	}

	if r.StatusCode < 200 || r.StatusCode >= 300 {
		var em models.ErrorMessage
		err := json.Unmarshal(raw, &em)
		if err != nil {
			return &models.SdkError{
				Err: err,
			}
		}
		return &models.SdkError{
			ErrorMessage: &em,
		}
	}

	return nil
}
