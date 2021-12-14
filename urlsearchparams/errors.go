package urlsearchparams

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented       = errors.New("Browser not implemented URLSearchParams")
	ErrNotAnURLSearchParams = errors.New("Object is not a URLSearchParams")
)
