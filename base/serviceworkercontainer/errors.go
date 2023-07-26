package serviceworkercontainer

import "errors"

var (
	ErrNotImplemented             = errors.New("Browser not implemented ServiceWorkerContainer")
	ErrNotAServiceWorkerContainer = errors.New("Object is not a ServiceWorkerContainer object")
	ErrControllerNotDefined       = errors.New("Controller not defined")
)
