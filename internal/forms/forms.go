package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

// Form creates a custom Form struct and embeds an url.Values object
type Form struct {
	url.Values
	Errors errors
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// NewForm initializes the From struct
func NewForm(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

//Required checks for required fields
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field can't be blank!")
		}
	}
}

// Has checks if Form field is in post and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.FormValue(field)
	if x == "" {
		f.Errors.Add(field, "This field can not be blank")
		return false
	}
	return true
}

// MinLen checks minimal length of a field
func (f *Form) MinLen(field string, length int, r *http.Request) bool {
	minLen := length
	x := r.FormValue(field)
	if len(x) < minLen {
		f.Errors.Add(field, fmt.Sprintf("Too Short! %s filed must be NLT %d characters!", field, minLen))
		return false
	}
	return true
}

func (f *Form) MaxLen(field string, length int, r *http.Request) bool {
	maxLen := length
	x := r.FormValue(field)
	if len(x) > maxLen {
		f.Errors.Add(field, fmt.Sprintf("Too Long! %s filed must be NMT %d characters!", field, maxLen))
		return false
	}
	return true
}



//manually written emal validators
func IsValidEmail(email string) bool {
	// Regular expression pattern for basic email validation
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	match, _ := regexp.MatchString(emailPattern, email)
	return match
}

func (f *Form) EmailFormat(field string, r *http.Request) bool {
	x := r.FormValue(field)
	if !IsValidEmail(x) {
		f.Errors.Add(field, "Wrong Format of Email! Must be name@domen.com")
		return false
	} else {
		return true
	}

}
