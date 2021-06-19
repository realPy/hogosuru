package progressevent

import "errors"

var (
	ErrNotImplemented = errors.New("Browser not implemented ProgressEvent")
	//ErrNotAnEventTarget ErrNotAnEventTarget error
	ErrNotAnProgressEvent = errors.New("Object is not a Progress Event")
)
