package services

import (
	"context"
	"fmt"
	"m5adminapi/internal/client"
	"m5adminapi/models"
	"strconv"
	"strings"
)

type VersionService struct {
	client *client.Client
}

func NewVersionService(c *client.Client) *VersionService {
	return &VersionService{client: c}
}

func (s *VersionService) Get(ctx context.Context) (*models.Version, *models.SdkError) {

	var version models.Version

	url := fmt.Sprintf("%s/version", s.client.Path)

	rs := client.RequestSettings{
		URL:                url,
		Method:             "GET",
		IncludeAccessToken: false,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return &version, &models.SdkError{Err: err}
	}

	// Expected response
	type versionResponse struct {
		V string `json:"version" validate:"required"`
	}

	r, em := client.ResponseConvert[versionResponse](resp)
	if r != nil {

		// parse version into models.version
		arr := strings.Split(r.V, ".")

		major, err := strconv.Atoi(arr[0])
		middle, err := strconv.Atoi(arr[1])
		minor, err := strconv.Atoi(arr[2])
		full, err := strconv.Atoi(strings.Replace(r.V, ".", "", -1))

		if err != nil {
			return nil, &models.SdkError{Err: fmt.Errorf("Could not parse response into object")}
		}

		version = models.Version{
			Major:    major,
			Middle:   middle,
			Minor:    minor,
			AsString: r.V,
			AsInt:    full,
		}

	}
	if em != nil {
		return &version, em
	}

	return &version, nil
}
