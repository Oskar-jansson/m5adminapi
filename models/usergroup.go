package models

type Usergroup struct {
	Id        *uint32 `json:"id" validate:"gt=0" cmp:"skip"`
	Groupname *string `json:"groupname" validate:"required"`
	Rastamp   *string `json:"rastamp" validate:"required" cmp:"skip"`
	Groupdata *string `json:"groupdata"`
	Users     []User  `json:"users"`
}

type UsergroupInput struct {
	Groupname *string `json:"groupname,omitempty"`
	Rastamp   *string `json:"rastamp,omitempty"`
	Groupdata *string `json:"groupdata,omitempty"`
}

type UsergroupList struct {
	Usergroups []Usergroup `json:"usergroups"`
	Count      *int        `json:"count,omitempty"` // only included with ?count='true'
}
