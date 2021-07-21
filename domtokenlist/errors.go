package domtokenlist

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented    = errors.New("Browser not implemented DOMTokenList")
	ErrNotAnDOMTokenList = errors.New("Object is not a DOMTokenList")
)
