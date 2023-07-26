package json

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented = errors.New("Browser not implemented JSON")
	ErrNotAJson       = errors.New("Object is not a Json")
)
