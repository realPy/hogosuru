package datatransferitem

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented       = errors.New("Browser not implemented DataTransferItem")
	ErrNotADataTransferItem = errors.New("Object is not a DataTransferItem")
)
