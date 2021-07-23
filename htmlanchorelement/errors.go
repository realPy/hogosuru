package htmlanchorelement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented         = errors.New("Browser not implemented HtmlAnchorElement")
	ErrNotAnHtmlAnchorElement = errors.New("Object is not an HtmlAnchorElement")
)
