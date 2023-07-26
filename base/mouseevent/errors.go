package mouseevent

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented = errors.New("Browser not implemented MouseEvent")
	ErrNotAMouseEvent = errors.New("Object is not a MouseEvent")
)
