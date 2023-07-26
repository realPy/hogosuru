package arraybuffer

import "errors"

var (
	ErrNotAnArrayBuffer = errors.New("The given value must be an arrayBuffer")
	ErrNotImplemented   = errors.New("Browser not implemented Arraybuffer")
)
