package htmlbrelement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented     = errors.New("Browser not implemented HTMLBRElement")
	ErrNotAnHtmlBrElement = errors.New("Object is not an HTMLBRElement")
)
