package models

type Timezone struct {
	Id          *uint32 `json:"id" validate:"gt=0" cmp:"skip"`
	Name        *string `json:"name" validate:"required"`
	Activedays  *string `json:"activedays" validate:"required"`
	Times       *string `json:"times" validate:"required"`
	Timeperiods *string `json:"timeperiods" validate:"required"`
	Rastamp     *string `json:"rastamp" validate:"required" cmp:"skip"`
	Allocstatus *uint32 `json:"allocstatus"` // undocumented
	Rasystem    *uint32 `json:"rasystem" validate:"required"`
}

type TimezoneList struct {
	Timezones []Timezone `json:"timezones"`
	Count     *int       `json:"count,omitempty"` // only included with ?count='true'
}
