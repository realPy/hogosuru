package promise

import "errors"

var (
	ErrNotImplemented     = errors.New("Browser not implemented Promise")
	ErrNotAPromise        = errors.New("Object is not a Promise")
	ErrResultPromiseError = errors.New("Result error promise")
)
