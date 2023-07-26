package serviceworkerglobalscope

import "errors"

var (
	ErrNotImplemented               = errors.New("Browser not implemented ServiceWorkerGlobalScope")
	ErrNotAServiceWorkerGlobalScope = errors.New("Object is not a ServiceWorkerGlobalScope object")
)
