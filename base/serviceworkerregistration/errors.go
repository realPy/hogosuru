package serviceworkerregistration

import "errors"

var (
	ErrNotImplemented                = errors.New("Browser not implemented ServiceWorkerRegistration")
	ErrNotAServiceWorkerRegistration = errors.New("Object is not a ServiceWorkerRegistration object")
)
