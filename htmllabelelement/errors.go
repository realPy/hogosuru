package htmllabelelement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented        = errors.New("Browser not implemented HTMLLabelElement")
	ErrNotAnHTMLLabelElement = errors.New("Object is not an HTMLLabelElement")
)
