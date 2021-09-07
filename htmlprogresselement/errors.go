package htmlprogresselement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented           = errors.New("Browser not implemented HtmlProgressElement")
	ErrNotAnHtmlProgressElement = errors.New("Object is not an HtmlProgressElement")
)
