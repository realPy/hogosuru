package iterator

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	EOI           = errors.New("End of iterator")
	NotAnIterator = errors.New("Object Not an iterator")
)
