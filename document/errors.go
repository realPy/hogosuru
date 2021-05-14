package document

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented = errors.New("Browser not implemented Document")
	//ErrNotADocument ErrNotADocument
	ErrNotADocument = errors.New("The given value must be a document")
)
