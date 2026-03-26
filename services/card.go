package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Oskar-jansson/m5adminapi/internal/client"
	"github.com/Oskar-jansson/m5adminapi/internal/options"
	"github.com/Oskar-jansson/m5adminapi/models"
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
		return nil, &models.SdkError{Err: fmt.Errorf("could not marshal input object to json")}
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
		return nil, &models.SdkError{Err: fmt.Errorf("could not marshal input object to json")}
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
		return &models.SdkError{Err: fmt.Errorf("could not marshal input object to json")}
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
		Id:                coalesceUint32(c2.Id, c1.Id),
		Fkusernumber:      coalesceUint32(c2.Fkusernumber, c1.Fkusernumber),
		Name:              coalesceString(c2.Name, c1.Name),
		Type:              coalesceUint32(c2.Type, c1.Type),
		Cardidentity:      coalesceString(c2.Cardidentity, c1.Cardidentity),
		Regnumber:         coalesceString(c2.Regnumber, c1.Regnumber),
		Pincode:           coalesceString(c2.Pincode, c1.Pincode),
		Startdate:         coalesceString(c2.Startdate, c1.Startdate),
		Enddate:           coalesceString(c2.Enddate, c1.Enddate),
		Alarmoff:          coalesceString(c2.Alarmoff, c1.Alarmoff),
		Alarmon:           coalesceString(c2.Alarmon, c1.Alarmon),
		Timecode:          coalesceUint32(c2.Timecode, c1.Timecode),
		Timecodetype:      coalesceUint32(c2.Timecodetype, c1.Timecodetype),
		Timebookings:      coalesceString(c2.Timebookings, c1.Timebookings),
		Booktype:          coalesceString(c2.Booktype, c1.Booktype),
		Rastamp:           coalesceString(c2.Rastamp, c1.Rastamp),
		Phoneshort:        coalesceString(c2.Phoneshort, c1.Phoneshort),
		Phonetele:         coalesceString(c2.Phonetele, c1.Phonetele),
		Phoneteleall:      coalesceString(c2.Phoneteleall, c1.Phoneteleall),
		Phoneuseexternal:  coalesceUint32(c2.Phoneuseexternal, c1.Phoneuseexternal),
		Disability:        coalesceUint32(c2.Disability, c1.Disability),
		Fieldex1:          coalesceString(c2.Fieldex1, c1.Fieldex1),
		Fieldex2:          coalesceString(c2.Fieldex2, c1.Fieldex2),
		Cardidentityraw:   coalesceString(c2.Cardidentityraw, c1.Cardidentityraw),
		Refguid:           coalesceString(c2.Refguid, c1.Refguid),
		Selectpindatetime: coalesceString(c2.Selectpindatetime, c1.Selectpindatetime),
		Apireference:      coalesceString(c2.Apireference, c1.Apireference),
		Changedby:         coalesceString(c2.Changedby, c1.Changedby),
		Changeddate:       coalesceString(c2.Changeddate, c1.Changeddate),
		Createdby:         coalesceString(c2.Createdby, c1.Createdby),
		Createddate:       coalesceString(c2.Createddate, c1.Createddate),
		Asciicard:         coalesceBool(c2.Asciicard, c1.Asciicard),
		Showregister:      coalesceBool(c2.Showregister, c1.Showregister),
		Pinblocked:        coalesceBool(c2.Pinblocked, c1.Pinblocked),
		Isblocked:         coalesceBool(c2.Isblocked, c1.Isblocked),
		Inherituseraccess: coalesceBool(c2.Inherituseraccess, c1.Inherituseraccess),
		Expired:           coalesceBool(c2.Expired, c1.Expired),
		User:              coalesce[models.User](c2.User, c1.User),
		Parentcard:        coalesceUint32(c2.Parentcard, c1.Parentcard),
		Vacation:          coalesceUint32(c2.Vacation, c1.Vacation),

		//nolint:all
		Accessalarms: coalesce(c2.Accessalarms, c1.Accessalarms),

		//nolint:all
		Accessgroups: coalesceSlice(c2.Accessgroups, c1.Accessgroups),

		//nolint:all
		Readeraccess: coalesceSlice(c2.Readeraccess, c1.Readeraccess),
	}
}
