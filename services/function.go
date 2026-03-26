package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Oskar-jansson/m5adminapi/internal/client"
	"github.com/Oskar-jansson/m5adminapi/internal/options"
	"github.com/Oskar-jansson/m5adminapi/models"
)

type FunctionService struct{ client *client.Client }

func NewFunctionService(c *client.Client) *FunctionService {
	return &FunctionService{client: c}
}

func (s *FunctionService) Get(ctx context.Context, id uint32, opt ...string) (*models.Function, *models.SdkError) {

	url := fmt.Sprintf("%s/function/%d", s.client.Path, id)

	rs := client.RequestSettings{
		URL:                options.QueryParamMerge(url, opt),
		Method:             "GET",
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return nil, &models.SdkError{Err: err}
	}

	return client.ResponseConvert[models.Function](resp)

}

func (s *FunctionService) List(ctx context.Context, opt ...string) (*models.FunctionList, *models.SdkError) {

	url := fmt.Sprintf("%s/function", s.client.Path)

	rs := client.RequestSettings{
		URL:                options.QueryParamMerge(url, opt),
		Method:             "GET",
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return nil, &models.SdkError{Err: err}
	}

	return client.ResponseConvert[models.FunctionList](resp)

}

func (s *FunctionService) Edit(ctx context.Context, id uint32, changes *models.FunctionInput, opt ...string) (*models.Function, *models.SdkError) {

	body, err := json.Marshal(changes)
	if err != nil {
		return nil, &models.SdkError{Err: fmt.Errorf("Could not marshal input object to json")}
	}

	url := fmt.Sprintf("%s/function/%d", s.client.Path, id)

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

	return client.ResponseConvert[models.Function](resp)

}

func (s *FunctionService) Create(ctx context.Context, function *models.FunctionInput, opt ...string) (*models.Function, *models.SdkError) {
	body, err := json.Marshal(function)
	if err != nil {
		return nil, &models.SdkError{Err: fmt.Errorf("Could not marshal input object to json")}
	}

	url := fmt.Sprintf("%s/function", s.client.Path)

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

	return client.ResponseConvert[models.Function](resp)

}

func (s *FunctionService) Delete(ctx context.Context, id uint32, opt ...string) *models.SdkError {

	url := fmt.Sprintf("%s/function/%d", s.client.Path, id)

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

func (s *FunctionService) ConvFunctioninputToFunction(f models.FunctionInput) *models.Function {
	return &models.Function{
		Fkaccessgroup: f.Fkaccessgroup,
		Fkunit:        f.Fkunit,
		Times:         f.Times,
		Timeperiods:   f.Timeperiods,
		Activedays:    f.Activedays,
		Type:          f.Type,
		Comment:       f.Comment,
		Rastamp:       f.Rastamp,
	}
}

// f2 fields take precedence over f1 fields (if f2 field is non-nil, use it).
func (s *FunctionService) MergeFunctions(f1, f2 models.Function) *models.Function {
	return &models.Function{
		Id:            coalesceUint32(f2.Id, f1.Id),
		Fkaccessgroup: coalesceUint32(f2.Fkaccessgroup, f1.Fkaccessgroup),
		Fkunit:        coalesceUint32(f2.Fkunit, f1.Fkunit),
		Times:         coalesceString(f2.Times, f1.Times),
		Timeperiods:   coalesceString(f2.Timeperiods, f1.Timeperiods),
		Activedays:    coalesceString(f2.Activedays, f1.Activedays),
		Type:          coalesceUint32(f2.Type, f1.Type),
		Comment:       coalesceString(f2.Comment, f1.Comment),
		Rastamp:       coalesceString(f2.Rastamp, f1.Rastamp),
		Unit:          coalesce[models.Unit](f2.Unit, f1.Unit),
		Accessgroup:   coalesce[models.Accessgroup](f2.Accessgroup, f1.Accessgroup),
	}
}
