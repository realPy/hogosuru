package datatransfert

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented = errors.New("Browser not implemented DataTransfert")
	//ErrNotADataTransfert ErrNotADataTransfert
	ErrNotADataTransfert = errors.New("The given value must be a dataTransfert")
)
