package pushmanager

import "errors"

var (
	ErrNotImplemented  = errors.New("Browser not implemented PushManager")
	ErrNotAPushManager = errors.New("Object is not a PushManager object")
)
