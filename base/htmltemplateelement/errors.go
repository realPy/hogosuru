package htmltemplateelement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented           = errors.New("Browser not implemented HTMLTemplateElement")
	ErrNotAnHTMLTemplateElement = errors.New("Object is not an HTMLTemplateElement")
)
