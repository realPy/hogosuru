package htmlembedelement

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented        = errors.New("Browser not implemented HTMLEmbedElement")
	ErrNotAnHtmlEmbedElement = errors.New("Object is not an HTMLEmbedElement")
)
