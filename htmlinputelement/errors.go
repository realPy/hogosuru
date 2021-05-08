package htmlinputelement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented        = errors.New("Browser not implemented HtmlInputElement")
	ErrNotAnHtmlInputElement = errors.New("Object is not an HTMLInputElement")
)
