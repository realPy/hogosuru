package htmlbodyelement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented       = errors.New("Browser not implemented HtmlBodyElement")
	ErrNotAnHtmlBodyElement = errors.New("Object is not an HtmlBodyElement")
)
