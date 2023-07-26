package htmllegendelement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented         = errors.New("Browser not implemented HTMLLegendElement")
	ErrNotAnHTMLLegendElement = errors.New("Object is not an HTMLLegendElement")
)
