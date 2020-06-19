package models

import "errors"

var (
	// ErrRedirectNotFound error if a redirect code not found
	ErrRedirectNotFound = errors.New("Redirect Not Found")
	// ErrRedirectInvalid error if a redirect is invalid
	ErrRedirectInvalid = errors.New("Redirect Invalid")
)

// Redirect is a model to construct code and redirect url
type Redirect struct {
	Code      string `json:"code"`
	URL       string `json:"url" validate:"empty=false & format=url"`
	CreatedAt int64  `json:"created_at"`
}
