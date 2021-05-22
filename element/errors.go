package element

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented  = errors.New("Browser not implemented Node")
	ErrNotAnElement    = errors.New("Object is not an Element")
	ErrElementNoChilds = errors.New("Element has no childs")
)
