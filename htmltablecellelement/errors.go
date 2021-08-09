package htmltablecellelement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented            = errors.New("Browser not implemented HTMLTableCellElement")
	ErrNotAnHTMLTableCellElement = errors.New("Object is not an HTMLTableCellElement")
)
