package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Oskar-jansson/m5adminapi/internal/client"
	"github.com/Oskar-jansson/m5adminapi/internal/options"
	"github.com/Oskar-jansson/m5adminapi/models"
)

type UsergroupService struct{ client *client.Client }

func NewUsergroupService(c *client.Client) *UsergroupService {
	return &UsergroupService{client: c}
}

func (s *UsergroupService) Get(ctx context.Context, id uint32, opt ...string) (*models.Usergroup, *models.SdkError) {

	url := fmt.Sprintf("%s/usergroup/%d", s.client.Path, id)

	rs := client.RequestSettings{
		URL:                options.QueryParamMerge(url, opt),
		Method:             "GET",
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return nil, &models.SdkError{Err: err}
	}

	return client.ResponseConvert[models.Usergroup](resp)

}

func (s *UsergroupService) List(ctx context.Context, opt ...string) (*models.UsergroupList, *models.SdkError) {

	url := fmt.Sprintf("%s/usergroup", s.client.Path)

	rs := client.RequestSettings{
		URL:                options.QueryParamMerge(url, opt),
		Method:             "GET",
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return nil, &models.SdkError{Err: err}
	}

	return client.ResponseConvert[models.UsergroupList](resp)

}

// Edit() allows to Patch usergroups.
//
// Needs to use current usergroup.Rastamp in body of usergroupInput to be a valid request.
func (s *UsergroupService) Edit(ctx context.Context, id uint32, changes *models.UsergroupInput, opt ...string) (*models.Usergroup, *models.SdkError) {

	body, err := json.Marshal(changes)
	if err != nil {
		return nil, &models.SdkError{Err: fmt.Errorf("could not marshal input object to json")}
	}

	url := fmt.Sprintf("%s/usergroup/%d", s.client.Path, id)

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

	return client.ResponseConvert[models.Usergroup](resp)

}

func (s *UsergroupService) Create(ctx context.Context, usergroup *models.UsergroupInput, opt ...string) (*models.Usergroup, *models.SdkError) {
	body, err := json.Marshal(usergroup)
	if err != nil {
		return nil, &models.SdkError{Err: fmt.Errorf("could not marshal input object to json")}
	}

	url := fmt.Sprintf("%s/usergroup", s.client.Path)

	rs := client.RequestSettings{
		URL:                options.QueryParamMerge(url, opt),
		Method:             "POST",
		Body:               body,
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return nil, &models.SdkError{Err: err}
	}

	return client.ResponseConvert[models.Usergroup](resp)

}

func (s *UsergroupService) Delete(ctx context.Context, id uint32, opt ...string) *models.SdkError {

	url := fmt.Sprintf("%s/usergroup/%d", s.client.Path, id)

	rs := client.RequestSettings{
		URL:                options.QueryParamMerge(url, opt),
		Method:             "DELETE",
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return &models.SdkError{Err: err}
	}

	return client.ResponseToSdkError(resp)

}

func (s *UsergroupService) ConvUsergroupinputToUsergroup(u models.UsergroupInput) *models.Usergroup {
	return &models.Usergroup{
		Groupname: u.Groupname,
		Rastamp:   u.Rastamp,
		Groupdata: u.Groupdata,
	}
}

// u2 fields take precedence over u1 fields (if u2 field is non-nil, use it).
func (s *UsergroupService) MergeUsergroups(u1, u2 models.Usergroup) *models.Usergroup {
	return &models.Usergroup{
		Id:        coalesceUint32(u2.Id, u1.Id),
		Groupname: coalesceString(u2.Groupname, u1.Groupname),
		Rastamp:   coalesceString(u2.Rastamp, u1.Rastamp),
		Groupdata: coalesceString(u2.Groupdata, u1.Groupdata),
		Users:     coalesceSlice(u2.Users, u1.Users),
	}
}
