package htmltablecaptionelement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented               = errors.New("Browser not implemented HTMLTableCaptionElement")
	ErrNotAnHTMLTableCaptionElement = errors.New("Object is not an HTMLTableCaptionElement")
)
