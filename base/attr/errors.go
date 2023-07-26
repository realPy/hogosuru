package attr

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented = errors.New("Browser not implemented Attr")
	ErrNotAnAttr      = errors.New("Object is not an Attr")
	ErrNoOwnerElement = errors.New("No owner element")
)
