package stream

import "errors"

var (
	ErrNotImplementedReadableStream = errors.New("Browser not implemented ReadableStream")
	ErrNotImplementedWritableStream = errors.New("Browser not implemented WritableStream")
	ErrNotAReadableStream           = errors.New("Object is not a ReadableStream")
	ErrNotAWritableStream           = errors.New("Object is not a WritableStream")
)
