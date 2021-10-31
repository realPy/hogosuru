package indexeddb

// https://developer.mozilla.org/fr/docs/Web/API/IDBDatabase

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/array"
	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/domstringlist"
	"github.com/realPy/hogosuru/event"
	"github.com/realPy/hogosuru/eventtarget"
)

//IDBDatabase struct
type IDBDatabase struct {
	eventtarget.EventTarget
}

type IDBDatabaseFrom interface {
	IDBDatabase_() IDBDatabase
}

func (i IDBDatabase) IDBDatabase_() IDBDatabase {
	return i
}

var singletonIDBDatabase sync.Once

var idbdatabaseinterface js.Value

func IDBDatabaseGetInterface() js.Value {

	singletonIDBDatabase.Do(func() {

		var err error
		if idbrequestinterface, err = baseobject.Get(js.Global(), "IDBDatabase"); err != nil {
			idbrequestinterface = js.Undefined()
		}

		baseobject.Register(idbrequestinterface, func(v js.Value) (interface{}, error) {
			return IDBDatabaseNewFromJSObject(v)
		})
	})

	return idbrequestinterface
}

func IDBDatabaseNewFromJSObject(obj js.Value) (IDBDatabase, error) {
	var i IDBDatabase
	var err error
	if ai := IDBDatabaseGetInterface(); !ai.IsUndefined() {
		if obj.IsUndefined() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(ai) {
				i.BaseObject = i.SetObject(obj)
			} else {
				err = ErrNotAnIDBDatabase
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return i, err
}

func IDBDatabaseNewFromObject(obj baseobject.BaseObject) (IDBDatabase, error) {

	return IDBDatabaseNewFromJSObject(obj.JSObject())
}

func (i IDBDatabase) Close() error {
	var err error
	_, err = i.Call("close")
	return err
}

func (i IDBDatabase) DeleteObjectStore(name string) error {
	var err error
	_, err = i.Call("deleteObjectStore", js.ValueOf(name))
	return err
}

func (i IDBDatabase) CreateObjectStore(name string, options ...map[string]interface{}) (IDBObjectStore, error) {
	var err error
	var obj js.Value
	var arrayJS []interface{}
	var s IDBObjectStore
	arrayJS = append(arrayJS, js.ValueOf(name))

	if len(options) > 0 {
		arrayJS = append(arrayJS, js.ValueOf(options[0]))
	}
	if obj, err = i.Call("createObjectStore", arrayJS...); err == nil {
		s, err = IDBObjectStoreNewFromJSObject(obj)
	}

	return s, err
}

func (i IDBDatabase) Transaction(store interface{}, mode ...string) (IDBTransaction, error) {
	var err error
	var obj js.Value
	var arrayJS []interface{}
	var t IDBTransaction

	//array of string ['my-store-name']
	if arr, ok := store.(array.Array); ok {
		arrayJS = append(arrayJS, arr.JSObject())
		//store name
	} else if storename, ok := store.(string); ok {
		arrayJS = append(arrayJS, js.ValueOf(storename))
	} else {
		err = ErrBadStoreType
	}

	if len(mode) > 0 {
		arrayJS = append(arrayJS, js.ValueOf(mode[0]))
	}

	if obj, err = i.Call("transaction", arrayJS...); err == nil {
		t, err = IDBTransactionNewFromJSObject(obj)
	}
	return t, err
}

func (i IDBDatabase) getAttributeInt(attribute string) (int64, error) {

	var err error
	var obj js.Value
	var ret int64

	if obj, err = i.Get(attribute); err == nil {

		if obj.Type() == js.TypeNumber {
			ret = int64(obj.Float())
		} else {
			err = baseobject.ErrObjectNotNumber
		}
	}
	return ret, err
}

func (i IDBDatabase) Name() (string, error) {
	return i.GetAttributeString("name")
}

func (i IDBDatabase) Version() (int64, error) {
	return i.getAttributeInt("version")
}

func (i IDBDatabase) ObjectStoreNames() (domstringlist.DOMStringList, error) {

	var err error
	var obj js.Value
	var d domstringlist.DOMStringList

	if obj, err = i.Get("objectStoreNames"); err == nil {
		d, err = domstringlist.NewFromJSObject(obj)
	}
	return d, err
}

func (i IDBDatabase) OnAbort(handler func(e event.Event)) (js.Func, error) {

	return i.AddEventListener("onabort", handler)
}

func (i IDBDatabase) OnError(handler func(e event.Event)) (js.Func, error) {

	return i.AddEventListener("onerror", handler)
}

func (i IDBDatabase) OnVersionChange(handler func(e event.Event)) (js.Func, error) {

	return i.AddEventListener("onversionchange", handler)
}
