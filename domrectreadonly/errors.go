package domrectreadonly

import "errors"

var (
	ErrNotImplemented       = errors.New("Browser not implemented DOMRectReadOnly")
	ErrNotAnDOMRectReadOnly = errors.New("The given value must be an DOMRectReadOnly")
)
