package forms

import (
	"net/url"
)

// Defining a custom form struct
type Form struct {
	url.Values
	Errors errors
}

// Initializes a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Function to check if a field in a form is empty
func (f *Form) Has(key string) bool {
	x := f.Get(key)
	if x == "" {
		f.Errors.Add(key, "This field cannot be empty")
		return false
	}
	return true
}
