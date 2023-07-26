package stream

import "errors"

var (
	ErrNotImplemented     = errors.New("Browser not implemented ReadableStream")
	ErrNotAReadableStream = errors.New("Object is not a ReadableStream")
)
