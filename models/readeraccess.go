package models

type Readeraccess struct {
	Id                 *uint32      `json:"id" validate:"gt=0" cmp:"skip"`
	Cardid             *uint32      `json:"cardid" validate:"required"`
	Timezoneid         *uint32      `json:"timezoneid" validate:"required"`
	Offlinetimelimit   *uint32      `json:"offlinetimelimit" validate:"required"`
	Dooropendisability *bool        `json:"dooropendisability" validate:"required"`
	Name               *string      `json:"name" validate:"required"`
	Unitid             *uint32      `json:"unitid"`
	Offlineunitid      *uint32      `json:"offlineunitid"`
	Rasystem           *uint32      `json:"rasystem" validate:"required"`
	Accessversion      *uint32      `json:"accessversion" validate:"required"`
	Accessblocked      *bool        `json:"accessblocked" validate:"required"`
	Noaccessversion    *bool        `json:"noaccessversion" validate:"required"`
	Card               *Card        `json:"card"`
	Unit               *Unit        `json:"unit"`
	Offlineunit        *Offlineunit `json:"offlineunit"`
}

type ReaderaccessList struct {
	Readeraccess []Readeraccess `json:"readeraccess"`
	Count        *int           `json:"count,omitempty"` // only included with ?count='true'
}

type ReaderaccessInput struct {
	Cardid             *uint32      `json:"cardid,omitempty"`
	Timezoneid         *uint32      `json:"timezoneid,omitempty"`
	Offlinetimelimit   *uint32      `json:"offlinetimelimit,omitempty"`
	Dooropendisability *bool        `json:"dooropendisability,omitempty"`
	Name               *string      `json:"name,omitempty"`
	Unitid             *uint32      `json:"unitid,omitempty"`
	Offlineunitid      *uint32      `json:"offlineunitid,omitempty"`
	Rasystem           *uint32      `json:"rasystem,omitempty"`
	Accessversion      *uint32      `json:"accessversion,omitempty"`
	Accessblocked      *bool        `json:"accessblocked,omitempty"`
	Noaccessversion    *bool        `json:"noaccessversion,omitempty"`
	Card               *Card        `json:"card,omitempty"`
	Unit               *Unit        `json:"unit,omitempty"`
	Offlineunit        *Offlineunit `json:"offlineunit,omitempty"`
}
