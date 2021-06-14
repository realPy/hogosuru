package indexeddb

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented = errors.New("Browser not implemented IDBFactory")
	//ErrNotAnIDBFactory ErrNotAnIDBFactory error
	ErrNotAnIDBFactory = errors.New("Object is not an IDBFactory")
	//ErrNotAnIDBOpenRequest ErrNotAnIDBOpenRequest error
	ErrNotAnIDBOpenRequest = errors.New("Object is not an IDBOpenrequest")
)
