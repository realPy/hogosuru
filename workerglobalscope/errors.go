package workerglobalscope

import "errors"

var (
	ErrNotImplemented        = errors.New("Browser not implemented WorkerGlobalScope")
	ErrNotAWorkerGlobalScope = errors.New("Object is not a WorkerGlobalScope object")
)
