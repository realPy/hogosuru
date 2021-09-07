package xmlhttprequest

import (
	"errors"
)

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented     = errors.New("Browser not implemented XMLHTTPRequest")
	ErrNotAXMLHTTPRequest = errors.New("Object is not a XMLHTTPRequest")
)
