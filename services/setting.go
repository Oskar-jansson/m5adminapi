package services

import (
	"context"
	"fmt"
	"m5adminapi/internal/client"
	"m5adminapi/internal/options"
	"m5adminapi/models"
)

type SettingService struct{ client *client.Client }

func NewSettingService(c *client.Client) *SettingService {
	return &SettingService{client: c}
}

func (s *SettingService) Get(ctx context.Context, id uint32, opt ...string) (*models.Setting, *models.SdkError) {

	url := fmt.Sprintf("%s/setting/%d", s.client.Path, id)

	rs := client.RequestSettings{
		URL:                options.QueryParamMerge(url, opt),
		Method:             "GET",
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return nil, &models.SdkError{Err: err}
	}

	return client.ResponseConvert[models.Setting](resp)

}

func (s *SettingService) List(ctx context.Context, opt ...string) (*models.SettingList, *models.SdkError) {

	url := fmt.Sprintf("%s/setting", s.client.Path)

	rs := client.RequestSettings{
		URL:                options.QueryParamMerge(url, opt),
		Method:             "GET",
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return nil, &models.SdkError{Err: err}
	}

	return client.ResponseConvert[models.SettingList](resp)

}
