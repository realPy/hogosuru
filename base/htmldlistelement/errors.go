package htmldlistelement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented        = errors.New("Browser not implemented HTMLDistElement")
	ErrNotAnHtmlDListElement = errors.New("Object is not an HTMLDistElement")
)
