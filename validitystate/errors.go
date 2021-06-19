package validitystate

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented     = errors.New("Browser not implemented HtmlInputElement")
	ErrNotAnValidityState = errors.New("Object is not an ValidityState")
)
