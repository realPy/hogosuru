package htmlheadelement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented       = errors.New("Browser not implemented HtmlHeadElement")
	ErrNotAnHtmlHeadElement = errors.New("Object is not an HtmlHeadElement")
)
