package htmlbuttonelement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented         = errors.New("Browser not implemented HTMLButtonElement")
	ErrNotAnHtmlButtonElement = errors.New("Object is not an HTMLButtonElement")
)
