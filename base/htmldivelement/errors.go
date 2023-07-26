package htmldivelement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented      = errors.New("Browser not implemented HTMLDivElement")
	ErrNotAnHtmlDivElement = errors.New("Object is not an HTMLDivElement")
)
