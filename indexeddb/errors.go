package indexeddb

import "errors"

var (
	//ErrNotImplemented ErrNotImplemented error
	ErrNotImplemented = errors.New("Browser not implemented IDBFactory")
	//ErrNotAnIDBFactory ErrNotAnIDBFactory error
	ErrNotAnIDBFactory = errors.New("Object is not an IDBFactory")
	//ErrNotAnIDBRequest ErrNotAnIDBOpenRequest error
	ErrNotAnIDBRequest = errors.New("Object is not an IDBRequest")
	//ErrNotAnIDBOpenRequest ErrNotAnIDBOpenRequest error
	ErrNotAnIDBOpenDBRequest = errors.New("Object is not an IDBOpenrequest")
	//ErrNotAnIDBTransaction ErrNotAnIDBOpenRequest error
	ErrNotAnIDBTransaction = errors.New("Object is not an IDBTransaction")
	//ErrNotAnIDBDatabase ErrNotAnIDBOpenRequest error
	ErrNotAnIDBDatabase = errors.New("Object is not an IDBDatabase")
	//ErrBadStoreType ErrBadStoreType error
	ErrBadStoreType = errors.New("Store type must be stringof array or string")
	//ErrNotAnIDBObjectStore ErrNotAnIDBObjectStore error
	ErrNotAnIDBObjectStore = errors.New("Object is not an ObjectStore")
	//ErrNotAnIDBIndex ErrNotAnIDBIndex error
	ErrNotAnIDBIndex = errors.New("Object is not IDBIndex")
	//ErrNotAnIDBIndex ErrNotAnIDBIndex error
	ErrNotAnIDBKeyRange = errors.New("Object is not IDBKeyRange")
)
