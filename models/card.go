package models

type Card struct {
	Id                *uint32 `json:"id" validate:"gt=0" cmp:"skip"`
	Fkusernumber      *uint32 `json:"fkusernumber" validate:"required"`
	Name              *string `json:"name" validate:"required"`
	Type              *uint32 `json:"type" validate:"required"`
	Cardidentity      *string `json:"cardidentity"`
	Regnumber         *string `json:"regnumber"`
	Pincode           *string `json:"pincode"`
	Startdate         *string `json:"startdate"`
	Enddate           *string `json:"enddate"`
	Alarmoff          *string `json:"alarmoff"`
	Alarmon           *string `json:"alarmon"`
	Timecode          *uint32 `json:"timecode"`
	Timecodetype      *uint32 `json:"timecodetype"`
	Timebookings      *string `json:"timebookings"`
	Booktype          *string `json:"booktype"`
	Rastamp           *string `json:"rastamp" validate:"required" cmp:"skip"`
	Phoneshort        *string `json:"phoneshort"`
	Phonetele         *string `json:"phonetele"`
	Phoneteleall      *string `json:"phoneteleall" validate:"required"`
	Phoneuseexternal  *uint32 `json:"phoneuseexternal"`
	Disability        *uint32 `json:"disability"`
	Fieldex1          *string `json:"fieldex1"`
	Fieldex2          *string `json:"fieldex2"`
	Cardidentityraw   *string `json:"cardidentityraw" validate:"required"`
	Refguid           *string `json:"refguid" validate:"required"`
	Selectpindatetime *string `json:"selectpindatetime"`
	Apireference      *string `json:"apireference"`
	Changedby         *string `json:"changedby" validate:"required" cmp:"skip"`
	Changeddate       *string `json:"changeddate" validate:"required" cmp:"skip"`
	Createdby         *string `json:"createdby" validate:"required" cmp:"skip"`
	Createddate       *string `json:"createddate" validate:"required" cmp:"skip"`
	Asciicard         *bool   `json:"asciicard" validate:"required"`
	Showregister      *bool   `json:"showregister" validate:"required"`
	Pinblocked        *bool   `json:"pinblocked" validate:"required"`
	Isblocked         *bool   `json:"isblocked" validate:"required"`
	Inherituseraccess *bool   `json:"inherituseraccess" validate:"required"`
	Expired           *bool   `json:"expired" validate:"required"`
	User              *User   `json:"user"`
	Parentcard        *uint32 `json:"parentcard"` // undocumented
	Vacation          *uint32 `json:"vacation"`   // undocumented

	// exposed when using ?include=accessalarms
	Accessalarms *[]struct {
		Fkaccessgroup uint32 `json:"fkaccessgroup"`
		Alarmon       uint32 `json:"alarmon"`
		Alarmoff      uint32 `json:"alarmoff"`
	} `json:"accessalarms"`

	Accessgroups []Accessgroup  `json:"accessgroups"`
	Readeraccess []Readeraccess `json:"readeraccess"`
}

type CardInput struct {
	Fkusernumber      *uint32 `json:"fkusernumber,omitempty"`
	Name              *string `json:"name,omitempty"`
	Type              *uint32 `json:"type,omitempty"`
	Cardidentity      *string `json:"cardidentity,omitempty"`
	Regnumber         *string `json:"regnumber,omitempty"`
	Pincode           *string `json:"pincode,omitempty"`
	Startdate         *string `json:"startdate,omitempty"`
	Enddate           *string `json:"enddate,omitempty"`
	Alarmoff          *string `json:"alarmoff,omitempty"`
	Alarmon           *string `json:"alarmon,omitempty"`
	Timecode          *uint32 `json:"timecode,omitempty"`
	Timecodetype      *uint32 `json:"timecodetype,omitempty"`
	Timebookings      *string `json:"timebookings,omitempty"`
	Booktype          *string `json:"booktype"`
	Rastamp           *string `json:"rastamp,omitempty"`
	Phoneshort        *string `json:"phoneshort,omitempty"`
	Phonetele         *string `json:"phonetele,omitempty"`
	Phoneteleall      *string `json:"phoneteleall,omitempty"`
	Phoneuseexternal  *uint32 `json:"phoneuseexternal,omitempty"`
	Disability        *uint32 `json:"disability,omitempty"`
	Fieldex1          *string `json:"fieldex1,omitempty"`
	Fieldex2          *string `json:"fieldex2,omitempty"`
	Cardidentityraw   *string `json:"cardidentityraw,omitempty"`
	Refguid           *string `json:"refguid,omitempty"`
	Selectpindatetime *string `json:"selectpindatetime,omitempty"`
	Apireference      *string `json:"apireference,omitempty"`
	Asciicard         *bool   `json:"asciicard,omitempty"`
	Showregister      *bool   `json:"showregister,omitempty"`
	Pinblocked        *bool   `json:"pinblocked,omitempty"`
	Isblocked         *bool   `json:"isblocked,omitempty"`
	Inherituseraccess *bool   `json:"inherituseraccess,omitempty"`

	//User              User    `json:",omitempty"`
	// Accessgroups      []Accessgroup
}

type CardList struct {
	Cards []Card `json:"cards"`
	Count *int   `json:"count,omitempty"` // only included with ?count='true'
}
