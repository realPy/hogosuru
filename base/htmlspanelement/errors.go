package htmlspanelement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented       = errors.New("Browser not implemented HTMLSpanElement")
	ErrNotAnHTMLSpanElement = errors.New("Object is not an HTMLSpanElement")
)
