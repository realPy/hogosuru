package htmlcollection

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented      = errors.New("Browser not implemented HTMLCollection")
	ErrNotAnHTMLCollection = errors.New("Object is not a HTMLCollection")
)
