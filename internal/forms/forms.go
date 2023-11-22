package forms

import (
	"net/http"
	"net/url"
)

// Form creates a custom Form struct and embeds an url.Values object
type Form struct{
	url.Values
	Errors errors
}


// NewForm initializes the From struct
func NewForm (data url.Values) *Form{
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Has checks if Form field is in post and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	x:=r.FormValue(field)
	if x==""{
		return false
	}
	return true
}