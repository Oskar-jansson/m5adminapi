package services

import (
	"context"
	"fmt"
	"m5adminapi/internal/client"
	"m5adminapi/models"
)

type SystemService struct {
	client *client.Client
}

func NewSystemService(c *client.Client) *SystemService {
	return &SystemService{client: c}
}

func (s *SystemService) List(ctx context.Context) (*models.System, *models.SdkError) {
	var systems *models.System

	url := fmt.Sprintf("%s/system", s.client.Path)

	rs := client.RequestSettings{
		URL:                url,
		Method:             "GET",
		IncludeAccessToken: false,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return systems, &models.SdkError{Err: err}
	}

	return client.ResponseConvert[models.System](resp)

}
