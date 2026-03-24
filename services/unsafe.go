package services

import (
	"context"
	"m5adminapi/internal/client"
	"m5adminapi/models"
)

type UnsafeService struct {
	client *client.Client
}

func New(c *client.Client) *UnsafeService {
	return &UnsafeService{client: c}
}

// all-purpose function to test malformed API-requests
func (s *UnsafeService) Request(ctx context.Context, endpoint string, method string, body []byte, IncludeAccesstoken bool) (*client.HttpResponse, *models.SdkError) {

	rs := client.RequestSettings{
		URL:                endpoint,
		Method:             method,
		IncludeAccessToken: IncludeAccesstoken,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return nil, &models.SdkError{Err: err}
	}

	return resp, nil

}
