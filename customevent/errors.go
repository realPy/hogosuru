package customevent

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented   = errors.New("Browser not implemented EventCustom")
	ErrNotAnCustomEvent = errors.New("Object is not a customEvent")
)
