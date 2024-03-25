package compressionstream

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented        = errors.New("Browser not implemented CompressionStream")
	ErrNotACompressionStream = errors.New("Object is not a CommpressionStream")
)
