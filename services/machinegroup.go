package services

import (
	"context"
	"fmt"

	"github.com/Oskar-jansson/m5adminapi/internal/client"
	"github.com/Oskar-jansson/m5adminapi/internal/options"
	"github.com/Oskar-jansson/m5adminapi/models"
)

type MachinegroupService struct{ client *client.Client }

func NewMachinegroupService(c *client.Client) *MachinegroupService {
	return &MachinegroupService{client: c}
}

func (s *MachinegroupService) Get(ctx context.Context, id uint32, opt ...string) (*models.Machinegroup, *models.SdkError) {

	url := fmt.Sprintf("%s/machinegroup/%d", s.client.Path, id)

	rs := client.RequestSettings{
		URL:                options.QueryParamMerge(url, opt),
		Method:             "GET",
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return nil, &models.SdkError{Err: err}
	}

	return client.ResponseConvert[models.Machinegroup](resp)

}

func (s *MachinegroupService) List(ctx context.Context, opt ...string) (*models.MachinegroupList, *models.SdkError) {

	url := fmt.Sprintf("%s/machinegroup", s.client.Path)

	rs := client.RequestSettings{
		URL:                options.QueryParamMerge(url, opt),
		Method:             "GET",
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return nil, &models.SdkError{Err: err}
	}

	return client.ResponseConvert[models.MachinegroupList](resp)

}
