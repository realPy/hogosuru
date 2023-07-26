package blob

import "errors"

var (
	//ErrNotAnBlob ErrNotAnBlob error
	ErrNotABlob       = errors.New("Object is not a Blob")
	ErrNotImplemented = errors.New("Browser not implemented Blob")
)
