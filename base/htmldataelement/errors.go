package htmldataelement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented       = errors.New("Browser not implemented HTMLDataElement")
	ErrNotAnHtmlDataElement = errors.New("Object is not an HTMLDataElement")
)
