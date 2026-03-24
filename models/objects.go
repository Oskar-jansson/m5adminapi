package models

// Keep all API resource structs here.

// Credentials used to authenticate
type Credentials struct {
	Id       string `json:"id"`
	System   string `json:"system"`
	Lang     string `json:"lang"`
	User     string `json:"user"`
	Password string `json:"password"`
	Apitype  string `json:"apitype"`
	Apikey   string `json:"apikey"`
}

// Response on defined errors from api side
type ErrorMessage struct {
	Errorcode        int    `json:"errorcode,omitempty"`
	Friendlymessage  string `json:"friendlymessage,omitempty"`
	Developermessage string `json:"developermessage,omitempty"`
	Moreinfo         string `json:"moreinfo,omitempty"`
	Reject           int    `json:"reject,omitempty"`
}

// genrelized error from services.
// Used instead of error, gives access to API error message directly
type SdkError struct {
	Err          error         // General error
	ErrorMessage *ErrorMessage // From Api
}
