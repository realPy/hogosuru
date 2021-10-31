package abortcontroller

import "errors"

var (
	ErrNotAnAbortController = errors.New("The given value must be an AbortController")
	ErrNotImplemented       = errors.New("Browser not implemented AbortController")
)
