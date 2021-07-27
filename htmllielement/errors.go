package htmllielement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented     = errors.New("Browser not implemented HTMLLIElement")
	ErrNotAnHTMLLIElement = errors.New("Object is not an HTMLLIElement")
)
