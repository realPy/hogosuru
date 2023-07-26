package jserror

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented = errors.New("Browser not implemented Error")
	ErrNotAnError     = errors.New("Object is not Error")
)
