package domstringlist

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented     = errors.New("Browser not implemented DOMStringList")
	ErrNotAnDOMStringList = errors.New("Object is not a DOMStringList")
)
