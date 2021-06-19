package domrectlist

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented   = errors.New("Browser not implemented DOMRectList")
	ErrNotAnDOMRectList = errors.New("Object is not a DOMRectList")
)
