package stylesheet

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented  = errors.New("Browser not implemented StyleSheet")
	ErrNotAnStyleSheet = errors.New("Object is not a StyleSheet")
)
