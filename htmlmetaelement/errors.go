package htmlmetaelement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented       = errors.New("Browser not implemented HTMLMetaElement")
	ErrNotAnHTMLMetaElement = errors.New("Object is not an HTMLMetaElement")
)
