package models

type Domain struct {
	Id               *uint32 `json:"id" validate:"required"`
	Name             *string `json:"name" validate:"required"`
	Useaccessversion *bool   `json:"useaccessversion" validate:"required"`
	Versionstarttime *string `json:"versionstarttime" cmp:"skip"`
}

type DomainList struct {
	Domains []Domain `json:"domains"`
	Count   *int     `json:"count,omitempty"` // only included with ?count='true'
}
