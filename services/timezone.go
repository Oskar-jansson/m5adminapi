package services

import (
	"context"
	"fmt"
	"m5adminapi/internal/client"
	"m5adminapi/internal/options"
	"m5adminapi/models"
)

type TimezoneService struct{ client *client.Client }

func NewTimezoneService(c *client.Client) *TimezoneService {
	return &TimezoneService{client: c}
}

func (s *TimezoneService) Get(ctx context.Context, id uint32, opt ...string) (*models.Timezone, *models.SdkError) {

	url := fmt.Sprintf("%s/timezone/%d", s.client.Path, id)

	rs := client.RequestSettings{
		URL:                options.QueryParamMerge(url, opt),
		Method:             "GET",
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return nil, &models.SdkError{Err: err}
	}

	return client.ResponseConvert[models.Timezone](resp)

}

func (s *TimezoneService) List(ctx context.Context, opt ...string) (*models.TimezoneList, *models.SdkError) {

	url := fmt.Sprintf("%s/timezone", s.client.Path)

	rs := client.RequestSettings{
		URL:                options.QueryParamMerge(url, opt),
		Method:             "GET",
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return nil, &models.SdkError{Err: err}
	}

	return client.ResponseConvert[models.TimezoneList](resp)

}
