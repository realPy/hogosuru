package datatranferitemlist

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented           = errors.New("Browser not implemented DataTransferItemList")
	ErrNotADataTransferItemList = errors.New("Object is not a DataTransferItemList")
)
