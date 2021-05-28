package domrect

import "errors"

var (
	ErrNotImplemented = errors.New("Browser not implemented DOMRect")
	ErrNotAnDOMRect   = errors.New("The given value must be an DOMRect")
)
