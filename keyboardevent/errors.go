package keyboardevent

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented    = errors.New("Browser not implemented KeyboardEvent")
	ErrNotAKeyboardEvent = errors.New("Object is not a KeyboardEvent")
)
