package services

import (
	"context"
	"fmt"

	"github.com/Oskar-jansson/m5adminapi/internal/client"
	"github.com/Oskar-jansson/m5adminapi/internal/options"
	"github.com/Oskar-jansson/m5adminapi/models"
)

type OfflineunitService struct{ client *client.Client }

func NewOfflineunitService(c *client.Client) *OfflineunitService {
	return &OfflineunitService{client: c}
}

func (s *OfflineunitService) Get(ctx context.Context, id uint32, opt ...string) (*models.Offlineunit, *models.SdkError) {

	url := fmt.Sprintf("%s/offlineunit/%d", s.client.Path, id)

	rs := client.RequestSettings{
		URL:                options.QueryParamMerge(url, opt),
		Method:             "GET",
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return nil, &models.SdkError{Err: err}
	}

	return client.ResponseConvert[models.Offlineunit](resp)

}

func (s *OfflineunitService) List(ctx context.Context, opt ...string) (*models.OfflineunitList, *models.SdkError) {

	url := fmt.Sprintf("%s/offlineunit", s.client.Path)

	rs := client.RequestSettings{
		URL:                options.QueryParamMerge(url, opt),
		Method:             "GET",
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return nil, &models.SdkError{Err: err}
	}

	return client.ResponseConvert[models.OfflineunitList](resp)

}

func (s *OfflineunitService) StepAccessVersion(ctx context.Context, offlineunitId int) *models.SdkError {

	url := fmt.Sprintf("%s/offlineunit/%d/stepaccessversion", s.client.Path, offlineunitId)

	rs := client.RequestSettings{
		URL:                url,
		Method:             "POST",
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return &models.SdkError{Err: err}
	}

	return client.ResponseToSdkError(resp)
}
