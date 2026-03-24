package services

import (
	"context"
	"encoding/json"
	"fmt"
	"m5adminapi/internal/client"
	"m5adminapi/internal/options"
	"m5adminapi/models"
)

type CardService struct{ client *client.Client }

func NewCardService(c *client.Client) *CardService {
	return &CardService{client: c}
}

func (s *CardService) Get(ctx context.Context, id uint32, opt ...string) (*models.Card, *models.SdkError) {

	url := fmt.Sprintf("%s/card/%d", s.client.Path, id)

	rs := client.RequestSettings{
		URL:                options.QueryParamMerge(url, opt),
		Method:             "GET",
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return nil, &models.SdkError{Err: err}
	}

	return client.ResponseConvert[models.Card](resp)

}

func (s *CardService) List(ctx context.Context, opt ...string) (*models.CardList, *models.SdkError) {

	url := fmt.Sprintf("%s/card", s.client.Path)

	rs := client.RequestSettings{
		URL:                options.QueryParamMerge(url, opt),
		Method:             "GET",
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return nil, &models.SdkError{Err: err}
	}

	return client.ResponseConvert[models.CardList](resp)

}

func (s *CardService) Edit(ctx context.Context, id uint32, changes *models.CardInput) (*models.Card, *models.SdkError) {

	body, err := json.Marshal(changes)
	if err != nil {
		return nil, &models.SdkError{Err: fmt.Errorf("Could not marshal input object to json")}
	}

	url := fmt.Sprintf("%s/card/%d", s.client.Path, id)

	rs := client.RequestSettings{
		URL:                url,
		Method:             "PATCH",
		Body:               body,
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return nil, &models.SdkError{Err: err}
	}

	return client.ResponseConvert[models.Card](resp)

}

func (s *CardService) Create(ctx context.Context, card *models.CardInput) (*models.Card, *models.SdkError) {
	body, err := json.Marshal(card)
	if err != nil {
		return nil, &models.SdkError{Err: fmt.Errorf("Could not marshal input object to json")}
	}

	url := fmt.Sprintf("%s/card", s.client.Path)

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

	return client.ResponseConvert[models.Card](resp)

}

func (s *CardService) Delete(ctx context.Context, id uint32) *models.SdkError {

	url := fmt.Sprintf("%s/card/%d", s.client.Path, id)

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

func (s *CardService) AssignAccessgroup(ctx context.Context, cardId uint32, accessgroupId uint32) *models.SdkError {
	url := fmt.Sprintf("%s/card/%d/accessgroup/%d", s.client.Path, cardId, accessgroupId)

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

func (s *CardService) DeleteAccessgroup(ctx context.Context, cardId uint32, accessgroupId uint32) *models.SdkError {
	url := fmt.Sprintf("%s/card/%d/accessgroup/%d", s.client.Path, cardId, accessgroupId)

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

func (s *CardService) BlockCards(ctx context.Context, list []int) *models.SdkError {

	// Assemble request body from list array
	cards := make([]map[string]int, len(list))
	for i, id := range list {
		cards[i] = map[string]int{"id": id}
	}

	mapList := map[string]any{
		"cards": cards,
	}

	body, err := json.Marshal(mapList)
	if err != nil {
		return &models.SdkError{Err: fmt.Errorf("Could not marshal input object to json")}
	}

	url := fmt.Sprintf("%s/card/block", s.client.Path)

	rs := client.RequestSettings{
		URL:                url,
		Method:             "POST",
		Body:               body,
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return &models.SdkError{Err: err}
	}

	return client.ResponseToSdkError(resp)
}

func (s *CardService) ConvCardinputToCard(c models.CardInput) *models.Card {
	return &models.Card{
		Fkusernumber:      c.Fkusernumber,
		Name:              c.Name,
		Type:              c.Type,
		Cardidentity:      c.Cardidentity,
		Regnumber:         c.Regnumber,
		Pincode:           c.Pincode,
		Startdate:         c.Startdate,
		Enddate:           c.Enddate,
		Alarmoff:          c.Alarmoff,
		Alarmon:           c.Alarmon,
		Timecode:          c.Timecode,
		Timecodetype:      c.Timecodetype,
		Timebookings:      c.Timebookings,
		Booktype:          c.Booktype,
		Rastamp:           c.Rastamp,
		Phoneshort:        c.Phoneshort,
		Phonetele:         c.Phonetele,
		Phoneteleall:      c.Phoneteleall,
		Phoneuseexternal:  c.Phoneuseexternal,
		Disability:        c.Disability,
		Fieldex1:          c.Fieldex1,
		Fieldex2:          c.Fieldex2,
		Cardidentityraw:   c.Cardidentityraw,
		Refguid:           c.Refguid,
		Selectpindatetime: c.Selectpindatetime,
		Apireference:      c.Apireference,
		Asciicard:         c.Asciicard,
		Showregister:      c.Showregister,
		Pinblocked:        c.Pinblocked,
		Isblocked:         c.Isblocked,
		Inherituseraccess: c.Inherituseraccess,
	}
}

// c2 fields take precedence over c1 fields (if c2 field is non-nil, use it).
func (s *CardService) MergeCards(c1, c2 models.Card) *models.Card {
	return &models.Card{
		Id:                coalesce(c2.Id, c1.Id),
		Fkusernumber:      coalesce(c2.Fkusernumber, c1.Fkusernumber),
		Name:              coalesce(c2.Name, c1.Name),
		Type:              coalesce(c2.Type, c1.Type),
		Cardidentity:      coalesce(c2.Cardidentity, c1.Cardidentity),
		Regnumber:         coalesce(c2.Regnumber, c1.Regnumber),
		Pincode:           coalesce(c2.Pincode, c1.Pincode),
		Startdate:         coalesce(c2.Startdate, c1.Startdate),
		Enddate:           coalesce(c2.Enddate, c1.Enddate),
		Alarmoff:          coalesce(c2.Alarmoff, c1.Alarmoff),
		Alarmon:           coalesce(c2.Alarmon, c1.Alarmon),
		Timecode:          coalesce(c2.Timecode, c1.Timecode),
		Timecodetype:      coalesce(c2.Timecodetype, c1.Timecodetype),
		Timebookings:      coalesce(c2.Timebookings, c1.Timebookings),
		Booktype:          coalesce(c2.Booktype, c1.Booktype),
		Rastamp:           coalesce(c2.Rastamp, c1.Rastamp),
		Phoneshort:        coalesce(c2.Phoneshort, c1.Phoneshort),
		Phonetele:         coalesce(c2.Phonetele, c1.Phonetele),
		Phoneteleall:      coalesce(c2.Phoneteleall, c1.Phoneteleall),
		Phoneuseexternal:  coalesce(c2.Phoneuseexternal, c1.Phoneuseexternal),
		Disability:        coalesce(c2.Disability, c1.Disability),
		Fieldex1:          coalesce(c2.Fieldex1, c1.Fieldex1),
		Fieldex2:          coalesce(c2.Fieldex2, c1.Fieldex2),
		Cardidentityraw:   coalesce(c2.Cardidentityraw, c1.Cardidentityraw),
		Refguid:           coalesce(c2.Refguid, c1.Refguid),
		Selectpindatetime: coalesce(c2.Selectpindatetime, c1.Selectpindatetime),
		Apireference:      coalesce(c2.Apireference, c1.Apireference),
		Changedby:         coalesce(c2.Changedby, c1.Changedby),
		Changeddate:       coalesce(c2.Changeddate, c1.Changeddate),
		Createdby:         coalesce(c2.Createdby, c1.Createdby),
		Createddate:       coalesce(c2.Createddate, c1.Createddate),
		Asciicard:         coalesce(c2.Asciicard, c1.Asciicard),
		Showregister:      coalesce(c2.Showregister, c1.Showregister),
		Pinblocked:        coalesce(c2.Pinblocked, c1.Pinblocked),
		Isblocked:         coalesce(c2.Isblocked, c1.Isblocked),
		Inherituseraccess: coalesce(c2.Inherituseraccess, c1.Inherituseraccess),
		Expired:           coalesce(c2.Expired, c1.Expired),
		User:              coalesce(c2.User, c1.User),
		Parentcard:        coalesce(c2.Parentcard, c1.Parentcard),
		Vacation:          coalesce(c2.Vacation, c1.Vacation),
		Accessalarms:      coalesce(c2.Accessalarms, c1.Accessalarms),
		Accessgroups:      coalesceSlice(c2.Accessgroups, c1.Accessgroups),
		Readeraccess:      coalesceSlice(c2.Readeraccess, c1.Readeraccess),
	}
}
