package url

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented = errors.New("Browser not implemented URL")
	ErrNotAURL        = errors.New("Object is not an URL")
)
