package htmloptionelement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented         = errors.New("Browser not implemented HTMLOptionElement")
	ErrNotAnHTMLOptionElement = errors.New("Object is not an HTMLOptionElement")
)
