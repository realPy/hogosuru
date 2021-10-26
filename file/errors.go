package file

import "errors"

var (
	ErrNotImplemented = errors.New("Browser not implemented File")
	ErrNotAFile       = errors.New("Object is not a File")
)
