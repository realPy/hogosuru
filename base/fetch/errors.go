package fetch

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented = errors.New("Browser not implemented Fetch")
	//ErrNotAnEventTarget ErrNotAnEventTarget error
	ErrNotAFetch        = errors.New("Object is not a FetchObject")
	ErrSignalNotManaged = errors.New("Signal is not managed by this instance")
)
