package services

import (
	"context"
	"fmt"

	"github.com/Oskar-jansson/m5adminapi/internal/client"
	"github.com/Oskar-jansson/m5adminapi/internal/options"
	"github.com/Oskar-jansson/m5adminapi/models"
)

type DateService struct{ client *client.Client }

func NewDateService(c *client.Client) *DateService {
	return &DateService{client: c}
}

func (s *DateService) List(ctx context.Context, opt ...string) (*models.DateList, *models.SdkError) {

	url := fmt.Sprintf("%s/date", s.client.Path)

	rs := client.RequestSettings{
		URL:                options.QueryParamMerge(url, opt),
		Method:             "GET",
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return nil, &models.SdkError{Err: err}
	}

	return client.ResponseConvert[models.DateList](resp)

}
