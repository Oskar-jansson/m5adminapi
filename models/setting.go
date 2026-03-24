package models

type Setting struct {
	Id      *uint32 `json:"id" validate:"gt=0" cmp:"skip"`
	Name    *string `json:"name" validate:"required"`
	Setting *string `json:"setting" validate:"required"`
}

type SettingList struct {
	Settings []Setting `json:"settings"`
	Count    *int      `json:"count,omitempty"` // only included with ?count='true'
}
