package documentfragment

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented = errors.New("Browser not implemented DocumentFragment")
	//ErrNotADocument ErrNotADocument
	ErrNotADocumentFragment = errors.New("The given value must be a DocumentFragment")
)
