package gocardless

import (
	"fmt"
	"net/url"
)

type Error struct {
	Message        string                   `json:"message,omitempty"`
	HTTPStatusCode int                      `json:"code,omitempty"`
	Errors         map[string][]interface{} `json:"errors,omitempty"`
	URL            *url.URL                 `json:"url,omitempty"`
}

func (r *Error) Error() string {
	return fmt.Sprintf("%d %s: %v %+v",
		r.HTTPStatusCode, r.Message, r.URL, r.Errors)
}

func (r *Error) GetErrorCode() string {
	return fmt.Sprintf("%d", r.HTTPStatusCode)
}

func (r *Error) GetErrorMessage() string {
	return fmt.Sprintf("%s", r.Message)
}

// prints the struct in a default format with the fields name
func (r *Error) GetErrors() string {
	return fmt.Sprintf("%+v", r.Errors)
}

// prints the error URL in the default format
func (r *Error) GetErrorURL() string {
	return fmt.Sprintf("%v", r.URL)
}
