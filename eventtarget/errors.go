package eventtarget

import "errors"

var (
	ErrNotImplemented = errors.New("Browser not implemented EventTarget")
	//ErrNotAnEventTarget ErrNotAnEventTarget error
	ErrNotAnEventTarget = errors.New("Object is not an EventTarget")
)
