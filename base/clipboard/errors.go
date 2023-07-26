package clipboard

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented = errors.New("Browser not implemented Clipboard")
	ErrNotAClipboard  = errors.New("Object is not a Clipboard")
)
