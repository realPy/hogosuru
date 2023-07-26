package dragevent

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented = errors.New("Browser not implemented DragEvent")
	ErrNotAnDragEvent = errors.New("Object is not an DragEvent")
)
