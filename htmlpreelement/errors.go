package htmlpreelement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented      = errors.New("Browser not implemented HTMLPreElement")
	ErrNotAnHTMLPreElement = errors.New("Object is not an HTMLPreElement")
)
