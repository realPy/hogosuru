package event

import "errors"

var (
	ErrNotAnEvent = errors.New("Object is not an Event")
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented = errors.New("Browser not implemented Event")
)
