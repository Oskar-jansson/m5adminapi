package models

type User struct {
	Id                    *uint32       `json:"id" validate:"gt=0" cmp:"skip"`
	Firstname             *string       `json:"firstname" validate:"required"`
	Lastname              *string       `json:"lastname" validate:"required"`
	Address               *string       `json:"address"`
	City                  *string       `json:"city"`
	Postalcode            *string       `json:"postalcode"`
	Phone1                *string       `json:"phone1"`
	Phone2                *string       `json:"phone2"`
	Phone3                *string       `json:"phone3"`
	Phoneconnectionnumber *int          `json:"phoneconnectionnumber"`
	Phoneconnectiontext   *string       `json:"phoneconnectiontext"`
	Email                 *string       `json:"email"`
	Extra                 *string       `json:"extra"`
	Employmentnumber      *uint32       `json:"employmentnumber"`
	Employmenttext        *string       `json:"employmenttext"`
	Customfield1          *string       `json:"customfield1"`
	Customfield2          *string       `json:"customfield2"`
	Customfield3          *string       `json:"customfield3"`
	Customfield4          *string       `json:"customfield4"`
	Customfield5          *string       `json:"customfield5"`
	Webpassword           *string       `json:"webpassword" cmp:"skip"` // non-readable
	Type                  *uint32       `json:"type"`
	Fkdepartment          *uint32       `json:"fkdepartment"`
	Fkusergroup           *uint32       `json:"fkusergroup"`
	Rastamp               *Rastamp      `json:"rastamp" validate:"required" cmp:"skip"`
	Fkfloor               *uint32       `json:"fkfloor"`
	Status                *uint32       `json:"status" validate:"required"`
	Startdate             *string       `json:"startdate" validate:"required"`
	Enddate               *string       `json:"enddate" validate:"required"`
	Refguid               *string       `json:"refguid" validate:"required"`
	Customfield6          *string       `json:"customfield6"`
	Customfield7          *string       `json:"customfield7"`
	Customfield8          *string       `json:"customfield8"`
	Customfield9          *string       `json:"customfield9"`
	Customfield10         *string       `json:"customfield10"`
	Apireference          *string       `json:"apireference"`
	Changedby             *ChangedBy    `json:"changedby" validate:"required" cmp:"skip"`
	Changeddate           *ChangedDate  `json:"changeddate" validate:"required" cmp:"skip"`
	Createdby             *CreatedBy    `json:"createdby" validate:"required" cmp:"skip"`
	Createddate           *CreatedDate  `json:"createddate" validate:"required" cmp:"skip"`
	Showregister          *bool         `json:"showregister" validate:"required"`
	Cardgroupname         *string       `json:"cardgroupname"`
	Message               *string       `json:"message"`
	Cardgroupstamp        *string       `json:"cardgroupstamp"`
	Cardtype              *int          `json:"cardtype"`
	Readinfomsg           *string       `json:"readinfomsg"`
	Remindpass            *int          `json:"remindpass"`
	Remindmachine         *int          `json:"remindmachine"`
	Balance               *float64      `json:"balance"`
	Lastcalc              *string       `json:"lastcalc"`
	Showmeasure           *int          `json:"showmeasure"`
	Showbooked            *int          `json:"showbooked"`
	Language              *int          `json:"language"`
	Cards                 []Card        `json:"cards"`
	Department            *Department   `json:"department"`
	Usergroup             *Usergroup    `json:"usergroup"`
	Floor                 *Floor        `json:"floor"`
	Accessgroups          []Accessgroup `json:"accessgroups"`
}

// Writeable fields in user
// Uses omitempty to allow incomplete structs to exclude said fields.
type UserInput struct {
	Firstname             *string `json:"firstname,omitempty"`
	Lastname              *string `json:"lastname,omitempty"`
	Address               *string `json:"address,omitempty"`
	City                  *string `json:"city,omitempty"`
	Postalcode            *string `json:"postalcode,omitempty"`
	Phone1                *string `json:"phone1,omitempty"`
	Phone2                *string `json:"phone2,omitempty"`
	Phone3                *string `json:"phone3,omitempty"`
	Phoneconnectionnumber *int    `json:"phoneconnectionnumber,omitempty"`
	Phoneconnectiontext   *string `json:"phoneconnectiontext,omitempty"`
	Email                 *string `json:"email,omitempty"`
	Extra                 *string `json:"extra,omitempty"`
	Employmentnumber      *uint32 `json:"employmentnumber,omitempty"`
	Employmenttext        *string `json:"employmenttext,omitempty"`
	Customfield1          *string `json:"customfield1,omitempty"`
	Customfield2          *string `json:"customfield2,omitempty"`
	Customfield3          *string `json:"customfield3,omitempty"`
	Customfield4          *string `json:"customfield4,omitempty"`
	Customfield5          *string `json:"customfield5,omitempty"`
	Webpassword           *string `json:"webpassword,omitempty"`

	// non writeable. But needed for valid PATCH requests.
	// Should always be current Rastamp of said user.
	Rastamp *Rastamp `json:"rastamp,omitempty"`

	Type          *uint32 `json:"type,omitempty"`
	Fkdepartment  *uint32 `json:"fkdepartment,omitempty"`
	Fkusergroup   *uint32 `json:"fkusergroup,omitempty"`
	Fkfloor       *uint32 `json:"fkfloor,omitempty"`
	Status        *uint32 `json:"status,omitempty"`
	Startdate     *string `json:"startdate,omitempty"`
	Enddate       *string `json:"enddate,omitempty"`
	Customfield6  *string `json:"customfield6,omitempty"`
	Customfield7  *string `json:"customfield7,omitempty"`
	Customfield8  *string `json:"customfield8,omitempty"`
	Customfield9  *string `json:"customfield9,omitempty"`
	Customfield10 *string `json:"customfield10,omitempty"`
	Apireference  *string `json:"apireference,omitempty"`
	Showregister  *bool   `json:"showregister,omitempty"`
	Cardgroupname *string `json:"cardgroupname,omitempty"`
	Message       *string `json:"message,omitempty"`
	Cardtype      *int    `json:"cardtype,omitempty"`

	//Cards        *[]Card
	//Department   *Department
	//Usergroup    *Usergroup
	//Floor        *Floor
	//Accessgroups *[]Accessgroup
}

type UserList struct {
	Users []User `json:"users"`
	Count *int   `json:"count,omitempty"` // only included with ?count='true'
}
