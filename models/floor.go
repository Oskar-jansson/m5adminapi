package models

type Floor struct {
	Id        *uint32 `json:"id" validate:"gt=0" cmp:"skip"`
	Floorname *string `json:"floorname" validate:"required"`
	Rastamp   *string `json:"rastamp" validate:"required" cmp:"skip"`
	Users     []User  `json:"users"`
}

type FloorInput struct {
	Floorname *string `json:"floorname,omitempty"`
	Rastamp   *string `json:"rastamp,omitempty"`
	Users     []User  `json:"users,omitempty"`
}

type FloorList struct {
	Floors []Floor `json:"floors"`
	Count  *int    `json:"count,omitempty"` // only included with ?count='true'
}
