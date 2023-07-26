package htmlmapelement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented      = errors.New("Browser not implemented HTMLMapElement")
	ErrNotAnHTMLMapElement = errors.New("Object is not an HTMLMapElement")
)
