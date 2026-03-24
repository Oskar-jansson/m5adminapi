package services

import (
	"context"
	"fmt"
	"m5adminapi/internal/client"
	"m5adminapi/internal/options"
	"m5adminapi/models"
)

type DomainService struct{ client *client.Client }

func NewDomainService(c *client.Client) *DomainService {
	return &DomainService{client: c}
}

func (s *DomainService) Get(ctx context.Context, id uint32, opt ...string) (*models.Domain, *models.SdkError) {

	url := fmt.Sprintf("%s/domain/%d", s.client.Path, id)

	rs := client.RequestSettings{
		URL:                options.QueryParamMerge(url, opt),
		Method:             "GET",
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return nil, &models.SdkError{Err: err}
	}

	return client.ResponseConvert[models.Domain](resp)

}

func (s *DomainService) List(ctx context.Context, opt ...string) (*models.DomainList, *models.SdkError) {

	url := fmt.Sprintf("%s/domain", s.client.Path)

	rs := client.RequestSettings{
		URL:                options.QueryParamMerge(url, opt),
		Method:             "GET",
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return nil, &models.SdkError{Err: err}
	}

	return client.ResponseConvert[models.DomainList](resp)

}
