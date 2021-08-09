package htmloptionscollection

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented             = errors.New("Browser not implemented HTMLOptionsCollection")
	ErrNotAnHTMLOptionsCollection = errors.New("Object is not a HTMLOptionsCollection")
)
