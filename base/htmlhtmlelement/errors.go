package htmlhtmlelement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented       = errors.New("Browser not implemented HTMLHtmlElement")
	ErrNotAnHtmlHtmlElement = errors.New("Object is not an HTMLHtmlElement")
)
