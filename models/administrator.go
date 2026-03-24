package models

type Administrator struct {
	Id        *uint32 `json:"id" validate:"gt=0" cmp:"skip"`
	Name      *string `json:"name" validate:"required"`
	Blocked   *bool   `json:"blocked" validate:"required"`
	Startdate *string `json:"startdate"`
	Enddate   *string `json:"enddate"`

	Apiadmin *uint32 `json:"apiadmin"` // undocumented
}

type AdministratorList struct {
	Administrators []Administrator `json:"administrators"`
	Count          *int            `json:"count,omitempty"` // only included with ?count='true'
}
