package idbdatabase

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotAnIDBDatabase = errors.New("Object is not a IDBDatabase")
)
