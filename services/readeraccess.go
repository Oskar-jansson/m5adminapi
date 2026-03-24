package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Oskar-jansson/m5adminapi/internal/client"
	"github.com/Oskar-jansson/m5adminapi/internal/options"
	"github.com/Oskar-jansson/m5adminapi/models"
)

type Readeraccess struct{ client *client.Client }

func NewReaderaccess(c *client.Client) *Readeraccess {
	return &Readeraccess{client: c}
}

func (s *Readeraccess) List(ctx context.Context, opt ...string) (*models.ReaderaccessList, *models.SdkError) {

	url := fmt.Sprintf("%s/readeraccess", s.client.Path)

	rs := client.RequestSettings{
		URL:                options.QueryParamMerge(url, opt),
		Method:             "GET",
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return nil, &models.SdkError{Err: err}
	}

	return client.ResponseConvert[models.ReaderaccessList](resp)

}

func (s *Readeraccess) Create(ctx context.Context, cardId int, Readeraccess *models.ReaderaccessInput) (*models.Readeraccess, *models.SdkError) {
	body, err := json.Marshal(Readeraccess)
	if err != nil {
		return nil, &models.SdkError{Err: fmt.Errorf("Could not marshal input object to json")}
	}

	url := fmt.Sprintf("%s/card/%d/readeraccess", s.client.Path, cardId)

	rs := client.RequestSettings{
		URL:                url,
		Method:             "POST",
		Body:               body,
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return nil, &models.SdkError{Err: err}
	}

	return client.ResponseConvert[models.Readeraccess](resp)

}

func (s *Readeraccess) Delete(ctx context.Context, cardId uint32, unitId uint32) *models.SdkError {

	url := fmt.Sprintf("%s/card/%d/readeraccess/%d", s.client.Path, cardId, unitId)

	rs := client.RequestSettings{
		URL:                url,
		Method:             "DELETE",
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return &models.SdkError{Err: err}
	}

	return client.ResponseToSdkError(resp)

}
