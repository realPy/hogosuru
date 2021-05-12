package websocket

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented   = errors.New("Browser not implemented WebSocket")
	ErrSendUnknownType  = errors.New("Unknown type send data provide to send method")
	ErrSetBadBinaryType = errors.New("Bad Binary Type set")
)
