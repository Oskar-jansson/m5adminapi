package services

import (
	"context"
	"fmt"
	"m5adminapi/internal/client"
	"m5adminapi/internal/options"
	"m5adminapi/models"
)

type AccessgroupService struct{ client *client.Client }

func NewAccessgroupService(c *client.Client) *AccessgroupService {
	return &AccessgroupService{client: c}
}

// Gets Accessgroup by id.
func (s *AccessgroupService) Get(ctx context.Context, id uint32) (*models.Accessgroup, *models.SdkError) {

	url := fmt.Sprintf("%s/accessgroup/%d", s.client.Path, id)

	rs := client.RequestSettings{
		URL:                url,
		Method:             "GET",
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return nil, &models.SdkError{Err: err}
	}

	return client.ResponseConvert[models.Accessgroup](resp)

}

// lists accessgroups
func (s *AccessgroupService) List(ctx context.Context, opt ...string) (*models.AccessgroupList, *models.SdkError) {

	url := fmt.Sprintf("%s/accessgroup", s.client.Path)

	rs := client.RequestSettings{
		URL:                options.QueryParamMerge(url, opt),
		Method:             "GET",
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return nil, &models.SdkError{Err: err}
	}

	return client.ResponseConvert[models.AccessgroupList](resp)

}
