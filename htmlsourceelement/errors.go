package htmlsourceelement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented         = errors.New("Browser not implemented HTMLSourceElement")
	ErrNotAnHTMLSourceElement = errors.New("Object is not an HTMLSourceElement")
)
