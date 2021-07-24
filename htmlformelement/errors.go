package htmlformelement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented       = errors.New("Browser not implemented HtmlFormElement")
	ErrNotAnHtmlFormElement = errors.New("Object is not an HtmlFormElement")
)
