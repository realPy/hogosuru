package htmltablecolelement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented           = errors.New("Browser not implemented HTMLTableColElement")
	ErrNotAnHTMLTableColElement = errors.New("Object is not an HTMLTableColElement")
)
