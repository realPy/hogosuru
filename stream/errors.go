package stream

import "errors"

var (
	//ErrNotAReadableStream ErrNotAnBlob error
	ErrNotAReadableStream = errors.New("Object is not a ReadableStream")
)
