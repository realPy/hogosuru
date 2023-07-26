package htmlstyleelement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented        = errors.New("Browser not implemented HTMLStyleElement")
	ErrNotAnHTMLStyleElement = errors.New("Object is not an HTMLStyleElement")
)
