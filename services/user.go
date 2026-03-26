package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Oskar-jansson/m5adminapi/internal/client"
	"github.com/Oskar-jansson/m5adminapi/internal/options"
	"github.com/Oskar-jansson/m5adminapi/models"
)

type UserService struct{ client *client.Client }

func NewUserService(c *client.Client) *UserService {
	return &UserService{client: c}
}

// Get specific User
// GET /user/1
func (s *UserService) Get(ctx context.Context, id uint32, opt ...string) (*models.User, *models.SdkError) {

	url := fmt.Sprintf("%s/user/%d", s.client.Path, id)

	rs := client.RequestSettings{
		URL:                options.QueryParamMerge(url, opt),
		Method:             "GET",
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return nil, &models.SdkError{Err: err}
	}

	return client.ResponseConvert[models.User](resp)

}

func (s *UserService) List(ctx context.Context, opt ...string) (*models.UserList, *models.SdkError) {

	url := fmt.Sprintf("%s/user", s.client.Path)

	rs := client.RequestSettings{
		URL:                options.QueryParamMerge(url, opt),
		Method:             "GET",
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return nil, &models.SdkError{Err: err}
	}

	return client.ResponseConvert[models.UserList](resp)

}

// Edit() allows to Patch users.
//
// Needs to use current User.Rastamp in body of userInput to be a valid request.
func (s *UserService) Edit(ctx context.Context, id uint32, changes *models.UserInput, opt ...string) (*models.User, *models.SdkError) {

	body, err := json.Marshal(changes)
	if err != nil {
		return nil, &models.SdkError{Err: fmt.Errorf("could not marshal input object to json")}
	}

	url := fmt.Sprintf("%s/user/%d", s.client.Path, id)

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

	return client.ResponseConvert[models.User](resp)

}

func (s *UserService) Create(ctx context.Context, user *models.UserInput, opt ...string) (*models.User, *models.SdkError) {
	body, err := json.Marshal(user)
	if err != nil {
		return nil, &models.SdkError{Err: fmt.Errorf("could not marshal input object to json")}
	}

	url := fmt.Sprintf("%s/user", s.client.Path)

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

	return client.ResponseConvert[models.User](resp)

}

func (s *UserService) Delete(ctx context.Context, id uint32, opt ...string) *models.SdkError {

	url := fmt.Sprintf("%s/user/%d", s.client.Path, id)

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

func (s *UserService) AssignAccessgroup(ctx context.Context, userId uint32, accessgroupId uint32) *models.SdkError {
	url := fmt.Sprintf("%s/user/%d/accessgroup/%d", s.client.Path, userId, accessgroupId)

	rs := client.RequestSettings{
		URL:                url,
		Method:             "PUT",
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return &models.SdkError{Err: err}
	}

	return client.ResponseToSdkError(resp)

}

func (s *UserService) DeleteAccessgroup(ctx context.Context, userId uint32, accessgroupId uint32) *models.SdkError {
	url := fmt.Sprintf("%s/user/%d/accessgroup/%d", s.client.Path, userId, accessgroupId)

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

func (s *UserService) ConvUserinputToUser(u models.UserInput) *models.User {
	return &models.User{
		Firstname:             u.Firstname,
		Lastname:              u.Lastname,
		Address:               u.Address,
		City:                  u.City,
		Postalcode:            u.Postalcode,
		Phone1:                u.Phone1,
		Phone2:                u.Phone2,
		Phone3:                u.Phone3,
		Phoneconnectionnumber: u.Phoneconnectionnumber,
		Phoneconnectiontext:   u.Phoneconnectiontext,
		Email:                 u.Email,
		Extra:                 u.Extra,
		Employmentnumber:      u.Employmentnumber,
		Employmenttext:        u.Employmenttext,
		Customfield1:          u.Customfield1,
		Customfield2:          u.Customfield2,
		Customfield3:          u.Customfield3,
		Customfield4:          u.Customfield4,
		Customfield5:          u.Customfield5,
		Webpassword:           u.Webpassword,
		Rastamp:               u.Rastamp,
		Type:                  u.Type,
		Fkdepartment:          u.Fkdepartment,
		Fkusergroup:           u.Fkusergroup,
		Fkfloor:               u.Fkfloor,
		Status:                u.Status,
		Startdate:             u.Startdate,
		Enddate:               u.Enddate,
		Customfield6:          u.Customfield6,
		Customfield7:          u.Customfield7,
		Customfield8:          u.Customfield8,
		Customfield9:          u.Customfield9,
		Customfield10:         u.Customfield10,
		Apireference:          u.Apireference,
		Showregister:          u.Showregister,
		Cardgroupname:         u.Cardgroupname,
		Message:               u.Message,
		Cardtype:              u.Cardtype,
	}
}

// u2 fields take precedence over u1 fields (if u2 field is non-nil, use it).
func (s *UserService) MergeUsers(u1, u2 models.User) *models.User {
	return &models.User{
		Id:                    coalesceUint32(u2.Id, u1.Id),
		Firstname:             coalesceString(u2.Firstname, u1.Firstname),
		Lastname:              coalesceString(u2.Lastname, u1.Lastname),
		Address:               coalesceString(u2.Address, u1.Address),
		City:                  coalesceString(u2.City, u1.City),
		Postalcode:            coalesceString(u2.Postalcode, u1.Postalcode),
		Phone1:                coalesceString(u2.Phone1, u1.Phone1),
		Phone2:                coalesceString(u2.Phone2, u1.Phone2),
		Phone3:                coalesceString(u2.Phone3, u1.Phone3),
		Phoneconnectionnumber: coalesceInt(u2.Phoneconnectionnumber, u1.Phoneconnectionnumber),
		Phoneconnectiontext:   coalesceString(u2.Phoneconnectiontext, u1.Phoneconnectiontext),
		Email:                 coalesceString(u2.Email, u1.Email),
		Extra:                 coalesceString(u2.Extra, u1.Extra),
		Employmentnumber:      coalesceUint32(u2.Employmentnumber, u1.Employmentnumber),
		Employmenttext:        coalesceString(u2.Employmenttext, u1.Employmenttext),
		Customfield1:          coalesceString(u2.Customfield1, u1.Customfield1),
		Customfield2:          coalesceString(u2.Customfield2, u1.Customfield2),
		Customfield3:          coalesceString(u2.Customfield3, u1.Customfield3),
		Customfield4:          coalesceString(u2.Customfield4, u1.Customfield4),
		Customfield5:          coalesceString(u2.Customfield5, u1.Customfield5),
		Webpassword:           coalesceString(u2.Webpassword, u1.Webpassword),
		Type:                  coalesceUint32(u2.Type, u1.Type),
		Fkdepartment:          coalesceUint32(u2.Fkdepartment, u1.Fkdepartment),
		Fkusergroup:           coalesceUint32(u2.Fkusergroup, u1.Fkusergroup),
		Rastamp:               coalesceString(u2.Rastamp, u1.Rastamp),
		Fkfloor:               coalesceUint32(u2.Fkfloor, u1.Fkfloor),
		Status:                coalesceUint32(u2.Status, u1.Status),
		Startdate:             coalesceString(u2.Startdate, u1.Startdate),
		Enddate:               coalesceString(u2.Enddate, u1.Enddate),
		Refguid:               coalesceString(u2.Refguid, u1.Refguid),
		Customfield6:          coalesceString(u2.Customfield6, u1.Customfield6),
		Customfield7:          coalesceString(u2.Customfield7, u1.Customfield7),
		Customfield8:          coalesceString(u2.Customfield8, u1.Customfield8),
		Customfield9:          coalesceString(u2.Customfield9, u1.Customfield9),
		Customfield10:         coalesceString(u2.Customfield10, u1.Customfield10),
		Apireference:          coalesceString(u2.Apireference, u1.Apireference),
		Changedby:             coalesceString(u2.Changedby, u1.Changedby),
		Changeddate:           coalesceString(u2.Changeddate, u1.Changeddate),
		Createdby:             coalesceString(u2.Createdby, u1.Createdby),
		Createddate:           coalesceString(u2.Createddate, u1.Createddate),
		Showregister:          coalesceBool(u2.Showregister, u1.Showregister),
		Cardgroupname:         coalesceString(u2.Cardgroupname, u1.Cardgroupname),
		Message:               coalesceString(u2.Message, u1.Message),
		Cardgroupstamp:        coalesceString(u2.Cardgroupstamp, u1.Cardgroupstamp),
		Cardtype:              coalesceInt(u2.Cardtype, u1.Cardtype),
		Readinfomsg:           coalesceString(u2.Readinfomsg, u1.Readinfomsg),
		Remindpass:            coalesceInt(u2.Remindpass, u1.Remindpass),
		Remindmachine:         coalesceInt(u2.Remindmachine, u1.Remindmachine),
		Balance:               coalesceFloat64(u2.Balance, u1.Balance),
		Lastcalc:              coalesceString(u2.Lastcalc, u1.Lastcalc),
		Showmeasure:           coalesceInt(u2.Showmeasure, u1.Showmeasure),
		Showbooked:            coalesceInt(u2.Showbooked, u1.Showbooked),
		Language:              coalesceInt(u2.Language, u1.Language),

		//nolint:all
		Cards: coalesceSlice(u2.Cards, u1.Cards),

		Department:   coalesce[models.Department](u2.Department, u1.Department),
		Usergroup:    coalesce[models.Usergroup](u2.Usergroup, u1.Usergroup),
		Floor:        coalesce[models.Floor](u2.Floor, u1.Floor),
		Accessgroups: coalesceSlice(u2.Accessgroups, u1.Accessgroups),
	}
}
