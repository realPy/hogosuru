package htmlareaelement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented       = errors.New("Browser not implemented HtmlAreaElement")
	ErrNotAnHtmlAreaElement = errors.New("Object is not an HtmlAreaElement")
)
