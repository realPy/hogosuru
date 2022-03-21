package worker

import "errors"

var (
	ErrNotImplemented = errors.New("Browser not implemented Worker")
	ErrNotAWorker     = errors.New("Object is not a Worker object")
)
