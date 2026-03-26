package client

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Oskar-jansson/m5adminapi/models"
)

func (c *Client) RequestStream(ctx context.Context, url string) (*bufio.Scanner, func() error, *models.SdkError) {

	var req *http.Request
	var err error

	client := c.newHTTPClient(0)

	req, err = http.NewRequestWithContext(ctx, "GET", url, bytes.NewBufferString(""))
	if err != nil {
		return nil, nil, &models.SdkError{Err: err}
	}

	req.Header.Set("Access-Token", fmt.Sprint(c.AccessToken))

	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, &models.SdkError{Err: err}
	}

	if resp.StatusCode != 200 {
		_ = resp.Body.Close() // disregard error as stream was never open. Instead use custom error.
		return nil, nil, &models.SdkError{Err: fmt.Errorf("status not ok, stream was not initiated")}
	}

	closer := resp.Body.Close
	return bufio.NewScanner(resp.Body), closer, nil
}

func (c *Client) ReadStream(ctx context.Context, scanner *bufio.Scanner, handler func(*models.Event), scannerClose func() error) {
	defer func() {
		if err := scannerClose(); err != nil {

			// attempt to send error over stream, otherwise void
			handler(&models.Event{InternalError: err})
		}
	}()

	for scanner.Scan() {
		if ctx.Err() != nil {
			handler(&models.Event{InternalError: fmt.Errorf("context was closed")})
			return
		}

		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		var event models.Event
		if err := json.Unmarshal([]byte(line), &event); err != nil {
			continue
		}
		handler(&event)
	}

	if err := scanner.Err(); err != nil {
		handler(&models.Event{InternalError: err})
	}
}
