package cssstyledeclaration

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented           = errors.New("Browser not implemented CSSStyleDeclaration")
	ErrNotAnCSSStyleDeclaration = errors.New("Object is not a CSSStyleDeclaration")
)
