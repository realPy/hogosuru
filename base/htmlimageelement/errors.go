package htmlimageelement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented       = errors.New("Browser not implemented HtmlImageElement")
	ErrNotAnHtmImageElement = errors.New("Object is not an HtmlImageElement")
)
