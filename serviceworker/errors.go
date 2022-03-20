package serviceworker

import "errors"

var (
	ErrNotImplemented    = errors.New("Browser not implemented Storage")
	ErrNotAServiceWorker = errors.New("Object is not a ServiceWorker object")
)
