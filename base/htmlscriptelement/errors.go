package htmlscriptelement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented         = errors.New("Browser not implemented HTMLScriptElement")
	ErrNotAnHTMLScriptElement = errors.New("Object is not an HTMLScriptElement")
)
