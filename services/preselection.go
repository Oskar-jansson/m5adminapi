package services

import (
	"context"
	"fmt"

	"github.com/Oskar-jansson/m5adminapi/internal/client"
	"github.com/Oskar-jansson/m5adminapi/internal/options"
	"github.com/Oskar-jansson/m5adminapi/models"
)

type PreselectionService struct{ client *client.Client }

func NewPreselectionService(c *client.Client) *PreselectionService {
	return &PreselectionService{client: c}
}

func (s *PreselectionService) Get(ctx context.Context, id uint32, opt ...string) (*models.Preselection, *models.SdkError) {

	url := fmt.Sprintf("%s/preselection/%d", s.client.Path, id)

	rs := client.RequestSettings{
		URL:                options.QueryParamMerge(url, opt),
		Method:             "GET",
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return nil, &models.SdkError{Err: err}
	}

	return client.ResponseConvert[models.Preselection](resp)

}

func (s *PreselectionService) List(ctx context.Context, opt ...string) (*models.PreselectionList, *models.SdkError) {

	url := fmt.Sprintf("%s/preselection", s.client.Path)

	rs := client.RequestSettings{
		URL:                options.QueryParamMerge(url, opt),
		Method:             "GET",
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return nil, &models.SdkError{Err: err}
	}

	return client.ResponseConvert[models.PreselectionList](resp)

}
