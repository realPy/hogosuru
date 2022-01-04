package eventsource

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented = errors.New("Browser not implemented Event Source")
	//ErrNotAnEventTarget ErrNotAnEventTarget error
	ErrNotAnEventSource = errors.New("Object is not an EventSource")
)
