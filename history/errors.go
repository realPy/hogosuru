package history

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented = errors.New("Browser not implemented History")

	ErrCantImplementedHistory = errors.New("Can't be implemented")
)
