package htmlhrelement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented     = errors.New("Browser not implemented HtmlHrElement")
	ErrNotAnHtmlHrElement = errors.New("Object is not an HtmlHrElement")
)
