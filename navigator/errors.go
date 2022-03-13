package navigator

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented = errors.New("Browser not implemented Navigator")
	ErrNotANavigator  = errors.New("Object is not a Navigator")
)
