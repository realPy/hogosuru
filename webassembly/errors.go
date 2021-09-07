package webassembly

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented  = errors.New("Browser not implemented WebAssembly")
	ErrNotAWebAssembly = errors.New("Object is not a WebAssembly")
)
