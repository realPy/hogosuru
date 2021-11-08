package htmlfieldsetelement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented           = errors.New("Browser not implemented HTMLFieldSetElement")
	ErrNotAnHtmlFieldSetElement = errors.New("Object is not an HTMLFieldSetElement")
	ErrNoForm                   = errors.New("No form attached to this object")
)
