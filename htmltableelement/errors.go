package htmltableelement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented           = errors.New("Browser not implemented HTMLTableElement")
	ErrNotAnHTMLTableColElement = errors.New("Object is not an HTMLTableElement")
)
