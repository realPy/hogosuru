package htmldetailselement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented          = errors.New("Browser not implemented HTMLDetailsElement")
	ErrNotAnHtmlDetailsElement = errors.New("Object is not an HTMLDetailsElement")
)
