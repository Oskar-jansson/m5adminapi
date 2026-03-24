package services

import (
	"context"
	"fmt"

	"github.com/Oskar-jansson/m5adminapi/internal/client"
	"github.com/Oskar-jansson/m5adminapi/internal/options"
	"github.com/Oskar-jansson/m5adminapi/models"
)

type UnitService struct{ client *client.Client }

func NewUnitService(c *client.Client) *UnitService {
	return &UnitService{client: c}
}

func (s *UnitService) Get(ctx context.Context, id uint32, opt ...string) (*models.Unit, *models.SdkError) {

	url := fmt.Sprintf("%s/unit/%d", s.client.Path, id)

	rs := client.RequestSettings{
		URL:                options.QueryParamMerge(url, opt),
		Method:             "GET",
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return nil, &models.SdkError{Err: err}
	}

	return client.ResponseConvert[models.Unit](resp)

}

func (s *UnitService) List(ctx context.Context, opt ...string) (*models.UnitList, *models.SdkError) {

	url := fmt.Sprintf("%s/unit", s.client.Path)

	rs := client.RequestSettings{
		URL:                options.QueryParamMerge(url, opt),
		Method:             "GET",
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return nil, &models.SdkError{Err: err}
	}

	return client.ResponseConvert[models.UnitList](resp)

}
