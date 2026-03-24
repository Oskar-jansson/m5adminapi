package services

import (
	"context"
	"fmt"

	"github.com/Oskar-jansson/m5adminapi/internal/client"
	"github.com/Oskar-jansson/m5adminapi/internal/options"
	"github.com/Oskar-jansson/m5adminapi/models"
)

type MachinegrouptypeService struct{ client *client.Client }

func NewMachinegrouptypeService(c *client.Client) *MachinegrouptypeService {
	return &MachinegrouptypeService{client: c}
}

func (s *MachinegrouptypeService) Get(ctx context.Context, id uint32, opt ...string) (*models.Machinegrouptype, *models.SdkError) {

	url := fmt.Sprintf("%s/machinegrouptype/%d", s.client.Path, id)

	rs := client.RequestSettings{
		URL:                options.QueryParamMerge(url, opt),
		Method:             "GET",
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return nil, &models.SdkError{Err: err}
	}

	return client.ResponseConvert[models.Machinegrouptype](resp)

}

func (s *MachinegrouptypeService) List(ctx context.Context, opt ...string) (*models.MachinegrouptypeList, *models.SdkError) {

	url := fmt.Sprintf("%s/machinegrouptype", s.client.Path)

	rs := client.RequestSettings{
		URL:                options.QueryParamMerge(url, opt),
		Method:             "GET",
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return nil, &models.SdkError{Err: err}
	}

	return client.ResponseConvert[models.MachinegrouptypeList](resp)

}
