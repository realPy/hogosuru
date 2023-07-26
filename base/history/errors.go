package history

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented = errors.New("Browser not implemented History")

	ErrNotAnHistory = errors.New("Object is not an History")
)
