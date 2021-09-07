package htmltablerowelement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented           = errors.New("Browser not implemented HTMLTableRowElement")
	ErrNotAnHTMLTableRowElement = errors.New("Object is not an HTMLTableRowElement")
)
