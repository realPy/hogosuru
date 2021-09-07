package namednodemap

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented   = errors.New("Browser not implemented NamedNodeMap")
	ErrNotANamedNodeMap = errors.New("Object is not a NamedNodeMap")
	ErrNotNameAttr      = errors.New("No Name attribue find")
)
