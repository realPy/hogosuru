package window

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented = errors.New("Browser not implemented Window")
	ErrNotAWindow     = errors.New("Object is not a Window")
)
