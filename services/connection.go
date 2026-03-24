package services

import (
	"context"
	"fmt"
	"m5adminapi/internal/client"
	"m5adminapi/internal/options"
	"m5adminapi/models"
)

type ConnectionService struct{ client *client.Client }

func NewConnectionService(c *client.Client) *ConnectionService {
	return &ConnectionService{client: c}
}

func (s *ConnectionService) Get(ctx context.Context, id uint32, opt ...string) (*models.Connection, *models.SdkError) {

	url := fmt.Sprintf("%s/connection/%d", s.client.Path, id)

	rs := client.RequestSettings{
		URL:                options.QueryParamMerge(url, opt),
		Method:             "GET",
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return nil, &models.SdkError{Err: err}
	}

	return client.ResponseConvert[models.Connection](resp)

}

func (s *ConnectionService) List(ctx context.Context, opt ...string) (*models.ConnectionList, *models.SdkError) {

	url := fmt.Sprintf("%s/connection", s.client.Path)

	rs := client.RequestSettings{
		URL:                options.QueryParamMerge(url, opt),
		Method:             "GET",
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return nil, &models.SdkError{Err: err}
	}

	return client.ResponseConvert[models.ConnectionList](resp)
}
