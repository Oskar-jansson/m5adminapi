package services

import (
	"context"
	"fmt"

	"github.com/Oskar-jansson/m5adminapi/internal/client"
	"github.com/Oskar-jansson/m5adminapi/internal/options"
	"github.com/Oskar-jansson/m5adminapi/models"
)

type AdministratorService struct{ client *client.Client }

func NewAdministratorService(c *client.Client) *AdministratorService {
	return &AdministratorService{client: c}
}

// Gets administrator/operator by id
func (s *AdministratorService) Get(ctx context.Context, id uint32, opt ...string) (*models.Administrator, *models.SdkError) {

	url := fmt.Sprintf("%s/administrator/%d", s.client.Path, id)

	rs := client.RequestSettings{
		URL:                options.QueryParamMerge(url, opt),
		Method:             "GET",
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return nil, &models.SdkError{Err: err}
	}

	return client.ResponseConvert[models.Administrator](resp)

}

// List administrators/operators
func (s *AdministratorService) List(ctx context.Context, opt ...string) (*models.AdministratorList, *models.SdkError) {

	url := fmt.Sprintf("%s/administrator", s.client.Path)

	rs := client.RequestSettings{
		URL:                options.QueryParamMerge(url, opt),
		Method:             "GET",
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return nil, &models.SdkError{Err: err}
	}

	return client.ResponseConvert[models.AdministratorList](resp)

}
