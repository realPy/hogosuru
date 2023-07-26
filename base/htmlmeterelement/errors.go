package htmlmeterelement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented        = errors.New("Browser not implemented HTMLMeterElement")
	ErrNotAnHTMLMeterElement = errors.New("Object is not an HTMLMeterElement")
)
