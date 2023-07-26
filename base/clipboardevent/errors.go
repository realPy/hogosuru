package clipboardevent

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented   = errors.New("Browser not implemented ClipboardEvent")
	ErrNotAnCustomEvent = errors.New("Object is not a ClipboardEvent")
)
