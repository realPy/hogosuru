package decompressionstream

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented          = errors.New("Browser not implemented DecompressionStream")
	ErrNotADecompressionStream = errors.New("Object is not a DecommpressionStream")
)
