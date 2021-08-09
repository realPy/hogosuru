package console

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented = errors.New("Browser not implemented Console")
	ErrNotAConsole    = errors.New("Object is not a Console")
)
