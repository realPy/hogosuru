package htmllinkelement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented       = errors.New("Browser not implemented HtmlAnchorElement")
	ErrNotAnHtmlLinkElement = errors.New("Object is not an HtmlLinkElement")
)
