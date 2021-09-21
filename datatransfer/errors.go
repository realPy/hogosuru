package datatransfer

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented = errors.New("Browser not implemented DataTransfer")
	//ErrNotADataTransfer ErrNotADataTransfer
	ErrNotADataTransfer = errors.New("The given value must be a dataTransfer")
)
