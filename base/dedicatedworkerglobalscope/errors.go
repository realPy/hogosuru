package dedicatedworkerglobalscope

import "errors"

var (
	ErrNotImplemented                 = errors.New("Browser not implemented DedicatedWorkerGlobalScope")
	ErrNotADedicatedWorkerGlobalScope = errors.New("Object is not a DedicatedWorkerGlobalScope object")
)
