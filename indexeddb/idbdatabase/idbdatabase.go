package idbdatabase

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/indexeddb/store"
)

var singleton sync.Once

var idbdatabaseinterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//GetJSInterface get the JS interface of formdata
func GetJSInterface() *JSInterface {

	singleton.Do(func() {
		var idbdatabaseinstance JSInterface
		var err error
		if idbdatabaseinstance.objectInterface, err = js.Global().GetWithErr("IDBDatabase"); err == nil {
			idbdatabaseinterface = &idbdatabaseinstance
		}
	})

	return idbdatabaseinterface
}

type IDBDatabase struct {
	baseobject.BaseObject
}

func NewFromJSObject(obj js.Value) (IDBDatabase, error) {

	var i IDBDatabase

	if idbi := GetJSInterface(); idbi != nil {
		if obj.InstanceOf(idbi.objectInterface) {
			i.BaseObject = i.SetObject(obj)
			return i, nil
		}

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
