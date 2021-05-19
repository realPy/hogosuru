package htmlelement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented   = errors.New("Browser not implemented HtmlElement")
	ErrNotAnHtmlElement = errors.New("Object is not an HTMLElement")
)
