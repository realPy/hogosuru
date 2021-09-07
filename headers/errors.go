package headers

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented = errors.New("Browser not implemented History")

	ErrNotAnHeaders = errors.New("Object is not an Headers")
)
