package messageevent

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented   = errors.New("Browser not implemented MessageEvent")
	ErrNotAMessageEvent = errors.New("Object is not a MessageEvent")
)
