package htmldatalistelement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented           = errors.New("Browser not implemented HTMLDataListElement")
	ErrNotAnHtmlDataListElement = errors.New("Object is not an HTMLDataListElement")
)
