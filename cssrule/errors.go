package cssrule

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented = errors.New("Browser not implemented CSSRule")
	ErrNotAnCSSRule   = errors.New("Object is not a CSSRule")
)
