package services

import (
	"context"
	"encoding/json"
	"fmt"
	"m5adminapi/internal/client"
	"m5adminapi/internal/options"
	"m5adminapi/models"
)

type FloorService struct{ client *client.Client }

func NewFloorService(c *client.Client) *FloorService {
	return &FloorService{client: c}
}

func (s *FloorService) Get(ctx context.Context, id uint32, opt ...string) (*models.Floor, *models.SdkError) {

	url := fmt.Sprintf("%s/floor/%d", s.client.Path, id)

	rs := client.RequestSettings{
		URL:                options.QueryParamMerge(url, opt),
		Method:             "GET",
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return nil, &models.SdkError{Err: err}
	}

	return client.ResponseConvert[models.Floor](resp)

}

func (s *FloorService) List(ctx context.Context, opt ...string) (*models.FloorList, *models.SdkError) {

	url := fmt.Sprintf("%s/floor", s.client.Path)

	rs := client.RequestSettings{
		URL:                options.QueryParamMerge(url, opt),
		Method:             "GET",
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return nil, &models.SdkError{Err: err}
	}

	return client.ResponseConvert[models.FloorList](resp)

}

func (s *FloorService) Edit(ctx context.Context, id uint32, changes *models.FloorInput, opt ...string) (*models.Floor, *models.SdkError) {

	body, err := json.Marshal(changes)
	if err != nil {
		return nil, &models.SdkError{Err: fmt.Errorf("Could not marshal input object to json")}
	}

	url := fmt.Sprintf("%s/floor/%d", s.client.Path, id)

	rs := client.RequestSettings{
		URL:                options.QueryParamMerge(url, opt),
		Method:             "PATCH",
		Body:               body,
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return nil, &models.SdkError{Err: err}
	}

	return client.ResponseConvert[models.Floor](resp)

}
