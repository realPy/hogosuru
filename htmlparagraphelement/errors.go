package htmlparagraphelement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented            = errors.New("Browser not implemented HTMLParagraphElement")
	ErrNotAnHTMLParagraphElement = errors.New("Object is not an HTMLParagraphElement")
)
