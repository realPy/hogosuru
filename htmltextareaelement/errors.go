package htmltextareaelement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented           = errors.New("Browser not implemented HTMLTextAreaElement")
	ErrNotAnHTMLTextAreaElement = errors.New("Object is not an HTMLTextAreaElement")
)
