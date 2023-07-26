package domexception

import "errors"

var (
	ErrNotImplemented   = errors.New("Browser not implemented DOMRectReadOnly")
	ErrNotADOMException = errors.New("The given value must be a DOMException")
)
