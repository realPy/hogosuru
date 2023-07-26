package htmltimeelement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented       = errors.New("Browser not implemented HTMLTimeElement")
	ErrNotAnHTMLTimeElement = errors.New("Object is not an HTMLTimeElement")
)
