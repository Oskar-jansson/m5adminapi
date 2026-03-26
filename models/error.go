package models

import "fmt"

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

// Returns the underlying error or API error message.
// Err is prioritized over ErrorMessage.
func (e *SdkError) AsError() error {

	if e == nil {
		return nil
	}

	if e.Err != nil {
		return e.Err
	}
	if e.ErrorMessage != nil {
		return fmt.Errorf("%s", e.ErrorMessage.Friendlymessage)
	}
	return nil
}

func (e *SdkError) Unwrap() error {
	if e == nil {
		return nil
	}

	return e.Err
}
