package location

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented = errors.New("Browser not implemented Location")
	ErrNotALocation   = errors.New("Object is not a Location")
)
