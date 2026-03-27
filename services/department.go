package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Oskar-jansson/m5adminapi/internal/client"
	"github.com/Oskar-jansson/m5adminapi/internal/options"
	"github.com/Oskar-jansson/m5adminapi/models"
)

type DepartmentService struct{ client *client.Client }

func NewDepartmentService(c *client.Client) *DepartmentService {
	return &DepartmentService{client: c}
}

func (s *DepartmentService) Get(ctx context.Context, id uint32, opt ...string) (*models.Department, *models.SdkError) {

	url := fmt.Sprintf("%s/Department/%d", s.client.Path, id)

	rs := client.RequestSettings{
		URL:                options.QueryParamMerge(url, opt),
		Method:             "GET",
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return nil, &models.SdkError{Err: err}
	}

	return client.ResponseConvert[models.Department](resp)

}

func (s *DepartmentService) List(ctx context.Context, opt ...string) (*models.DepartmentList, *models.SdkError) {

	url := fmt.Sprintf("%s/Department", s.client.Path)

	rs := client.RequestSettings{
		URL:                options.QueryParamMerge(url, opt),
		Method:             "GET",
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return nil, &models.SdkError{Err: err}
	}

	return client.ResponseConvert[models.DepartmentList](resp)

}

func (s *DepartmentService) Edit(ctx context.Context, id uint32, changes *models.DepartmentInput, opt ...string) (*models.Department, *models.SdkError) {

	body, err := json.Marshal(changes)
	if err != nil {
		return nil, &models.SdkError{Err: fmt.Errorf("could not marshal input object to json")}
	}

	url := fmt.Sprintf("%s/Department/%d", s.client.Path, id)

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

	return client.ResponseConvert[models.Department](resp)

}

func (s *DepartmentService) Create(ctx context.Context, Department *models.DepartmentInput, opt ...string) (*models.Department, *models.SdkError) {
	body, err := json.Marshal(Department)
	if err != nil {
		return nil, &models.SdkError{Err: fmt.Errorf("could not marshal input object to json")}
	}

	url := fmt.Sprintf("%s/Department", s.client.Path)

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

	return client.ResponseConvert[models.Department](resp)

}

func (s *DepartmentService) Delete(ctx context.Context, id uint32, opt ...string) *models.SdkError {

	url := fmt.Sprintf("%s/Department/%d", s.client.Path, id)

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

func (s *DepartmentService) ConvDepartmentinputToDepartment(u models.DepartmentInput) *models.Department {
	return &models.Department{

		Departmentname: u.Departmentname,
		Rastamp:        u.Rastamp,
		Departmentdata: u.Departmentdata,
	}
}

// u2 fields take precedence over u1 fields (if u2 field is non-nil, use it).
func (s *DepartmentService) MergeDepartments(u1, u2 models.Department) *models.Department {
	return &models.Department{
		Id:             coalesceUint32(u2.Id, u1.Id),
		Departmentname: coalesceString(u2.Departmentname, u1.Departmentname),
		Rastamp:        coalesce(u2.Rastamp, u1.Rastamp), //nolint:all
		Departmentdata: coalesceString(u2.Departmentdata, u1.Departmentdata),
	}
}
