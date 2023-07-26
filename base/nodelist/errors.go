package nodelist

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented = errors.New("Browser not implemented NodeList")
	ErrNotAnNodeList  = errors.New("Object is not a NodeList")
)
