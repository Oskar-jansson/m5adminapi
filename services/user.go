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
		return nil, &models.SdkError{Err: fmt.Errorf("Could not marshal input object to json")}
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
		return nil, &models.SdkError{Err: fmt.Errorf("Could not marshal input object to json")}
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
		Id:                    coalesce(u2.Id, u1.Id),
		Firstname:             coalesce(u2.Firstname, u1.Firstname),
		Lastname:              coalesce(u2.Lastname, u1.Lastname),
		Address:               coalesce(u2.Address, u1.Address),
		City:                  coalesce(u2.City, u1.City),
		Postalcode:            coalesce(u2.Postalcode, u1.Postalcode),
		Phone1:                coalesce(u2.Phone1, u1.Phone1),
		Phone2:                coalesce(u2.Phone2, u1.Phone2),
		Phone3:                coalesce(u2.Phone3, u1.Phone3),
		Phoneconnectionnumber: coalesce(u2.Phoneconnectionnumber, u1.Phoneconnectionnumber),
		Phoneconnectiontext:   coalesce(u2.Phoneconnectiontext, u1.Phoneconnectiontext),
		Email:                 coalesce(u2.Email, u1.Email),
		Extra:                 coalesce(u2.Extra, u1.Extra),
		Employmentnumber:      coalesce(u2.Employmentnumber, u1.Employmentnumber),
		Employmenttext:        coalesce(u2.Employmenttext, u1.Employmenttext),
		Customfield1:          coalesce(u2.Customfield1, u1.Customfield1),
		Customfield2:          coalesce(u2.Customfield2, u1.Customfield2),
		Customfield3:          coalesce(u2.Customfield3, u1.Customfield3),
		Customfield4:          coalesce(u2.Customfield4, u1.Customfield4),
		Customfield5:          coalesce(u2.Customfield5, u1.Customfield5),
		Webpassword:           coalesce(u2.Webpassword, u1.Webpassword),
		Type:                  coalesce(u2.Type, u1.Type),
		Fkdepartment:          coalesce(u2.Fkdepartment, u1.Fkdepartment),
		Fkusergroup:           coalesce(u2.Fkusergroup, u1.Fkusergroup),
		Rastamp:               coalesce(u2.Rastamp, u1.Rastamp),
		Fkfloor:               coalesce(u2.Fkfloor, u1.Fkfloor),
		Status:                coalesce(u2.Status, u1.Status),
		Startdate:             coalesce(u2.Startdate, u1.Startdate),
		Enddate:               coalesce(u2.Enddate, u1.Enddate),
		Refguid:               coalesce(u2.Refguid, u1.Refguid),
		Customfield6:          coalesce(u2.Customfield6, u1.Customfield6),
		Customfield7:          coalesce(u2.Customfield7, u1.Customfield7),
		Customfield8:          coalesce(u2.Customfield8, u1.Customfield8),
		Customfield9:          coalesce(u2.Customfield9, u1.Customfield9),
		Customfield10:         coalesce(u2.Customfield10, u1.Customfield10),
		Apireference:          coalesce(u2.Apireference, u1.Apireference),
		Changedby:             coalesce(u2.Changedby, u1.Changedby),
		Changeddate:           coalesce(u2.Changeddate, u1.Changeddate),
		Createdby:             coalesce(u2.Createdby, u1.Createdby),
		Createddate:           coalesce(u2.Createddate, u1.Createddate),
		Showregister:          coalesce(u2.Showregister, u1.Showregister),
		Cardgroupname:         coalesce(u2.Cardgroupname, u1.Cardgroupname),
		Message:               coalesce(u2.Message, u1.Message),
		Cardgroupstamp:        coalesce(u2.Cardgroupstamp, u1.Cardgroupstamp),
		Cardtype:              coalesce(u2.Cardtype, u1.Cardtype),
		Readinfomsg:           coalesce(u2.Readinfomsg, u1.Readinfomsg),
		Remindpass:            coalesce(u2.Remindpass, u1.Remindpass),
		Remindmachine:         coalesce(u2.Remindmachine, u1.Remindmachine),
		Balance:               coalesce(u2.Balance, u1.Balance),
		Lastcalc:              coalesce(u2.Lastcalc, u1.Lastcalc),
		Showmeasure:           coalesce(u2.Showmeasure, u1.Showmeasure),
		Showbooked:            coalesce(u2.Showbooked, u1.Showbooked),
		Language:              coalesce(u2.Language, u1.Language),
		Cards:                 coalesceSlice(u2.Cards, u1.Cards),
		Department:            coalesce(u2.Department, u1.Department),
		Usergroup:             coalesce(u2.Usergroup, u1.Usergroup),
		Floor:                 coalesce(u2.Floor, u1.Floor),
		Accessgroups:          coalesceSlice(u2.Accessgroups, u1.Accessgroups),
	}
}
