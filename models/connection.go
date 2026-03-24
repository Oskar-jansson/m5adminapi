package models

type Connection struct {
	Id              *uint32 `json:"id" validate:"gt=0" cmp:"skip"`
	Name            *string `json:"name" validate:"required"`
	Type            *uint32 `json:"type" validate:"required"`
	Baseaddress     *uint32 `json:"baseaddress" validate:"required"`
	State           *uint32 `json:"state" validate:"required"`
	Version         *string `json:"version"`
	Created         *bool   `json:"created" validate:"required"`
	Rastamp         *string `json:"rastamp" validate:"required" cmp:"skip"`
	Rasystem        *uint32 `json:"rasystem" validate:"required"`
	Progloc         *uint32 `json:"progloc" validate:"required"`
	Getipfromnotify *uint32 `json:"getipfromnotify" validate:"required"`
	Afevents        *uint32 `json:"afevents" validate:"required"`
	Autoupdate      *uint32 `json:"autoupdate" validate:"required"`
	Localcon        *uint32 `json:"localcon" validate:"required"`
	Needupdate      *uint32 `json:"needupdate" validate:"required"`
	Remotename      *string `json:"remotename" validate:"required"`
	Remoteport      *string `json:"remoteport" validate:"required"`
	Localport       *string `json:"localport" validate:"required"`
	Lastupdated     *string `json:"lastupdated" validate:"required"`
	Lastupdatedtime *string `json:"lastupdatedtime" validate:"required"`
}

type ConnectionList struct {
	Connections []Connection `json:"connections"`
	Count       *int         `json:"count,omitempty"` // only included with ?count='true'
}
