package formdata

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented        = errors.New("Browser not implemented FormData")
	ErrNotAFormData          = errors.New("Object is not a FormData")
	ErrNotAFormValueNotFound = errors.New("Form Value not found")
)
