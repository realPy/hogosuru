package idbdatabase

import (
	"syscall/js"

	"github.com/realPy/hogosuru/indexeddb/store"
	"github.com/realPy/hogosuru/object"
)

type IDBDatabase struct {
	object.Object
}

func NewFromJSObject(obj js.Value) (IDBDatabase, error) {

	var i IDBDatabase

	if object.String(obj) == "[object IDBDatabase]" {
		i.Object = i.SetObject(obj)
		return i, nil
	}

	return i, ErrNotAnIDBDatabase

}

func (i IDBDatabase) CreateStore(name string, schema map[string]interface{}) (store.Store, error) {

	if storeObject, err := i.JSObject().CallWithErr("createObjectStore", js.ValueOf(name), schema); err == nil {

		return store.NewFromJSObject(storeObject)
	} else {
		return store.Store{}, err
	}

}

func (i IDBDatabase) GetObjectStore(table string, permission string) (store.Store, error) {
	if transaction, err := i.JSObject().CallWithErr("transaction", js.ValueOf(table), js.ValueOf(permission)); err == nil {

		if objectstore, err := transaction.CallWithErr("objectStore", js.ValueOf(table)); err == nil {
			return store.NewFromJSObject(objectstore)
		} else {
			return store.Store{}, err
		}
	} else {
		return store.Store{}, err
	}
}
