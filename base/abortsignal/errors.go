package abortsignal

import "errors"

var (
	ErrNotAnAbortSignal = errors.New("The given value must be an AbortSignal")
	ErrNotImplemented   = errors.New("Browser not implemented AbortSignal")
)
