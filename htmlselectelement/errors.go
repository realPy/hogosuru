package htmlselectelement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented         = errors.New("Browser not implemented HTMLSelectElement")
	ErrNotAnHTMLSelectElement = errors.New("Object is not an HTMLSelectElement")
)
