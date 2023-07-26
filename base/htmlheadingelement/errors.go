package htmlheadingelement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented          = errors.New("Browser not implemented HtmlHeadingElement")
	ErrNotAnHtmlHeadingElement = errors.New("Object is not an HtmlHeadingElement")
)
