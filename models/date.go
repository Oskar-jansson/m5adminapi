package models

type Date struct {
	Id           *uint32 `json:"id" validate:"gt=0" cmp:"skip"`
	Calendardate *string `json:"calendardate"`
	Name         *string `json:"name"`
	Rasystem     *uint32 `json:"rasystem"`
	Rastamp      *string `json:"rastamp" cmp:"skip"`
}

type DateList struct {
	Dates []Date `json:"dates"`
	Count *int   `json:"count,omitempty"` // only included with ?count='true'
}
