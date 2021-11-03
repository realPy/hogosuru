package htmliframeelement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented         = errors.New("Browser not implemented HtmlIFrameElement")
	ErrNotAnHtmlIFrameElement = errors.New("Object is not an HtmlIFrameElement")
	ErrNoContentDocument      = errors.New("No Content Document")
)
