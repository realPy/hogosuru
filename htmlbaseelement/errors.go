package htmlbaseelement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented       = errors.New("Browser not implemented HtmlBaseElement")
	ErrNotAnHtmlBaseElement = errors.New("Object is not an HtmlBaseElement")
)
