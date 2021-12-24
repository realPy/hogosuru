package indexeddb

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/eventtarget"
)

//https://developer.mozilla.org/fr/docs/Web/API/IDBObjectStores

//IDBObjectStore struct
type IDBObjectStore struct {
	eventtarget.EventTarget
}

type IDBObjectStoreFrom interface {
	IDBObjectStore_() IDBObjectStore
}

func (i IDBObjectStore) IDBObjectStore_() IDBObjectStore {
	return i
}

var singletonIDBObjectStore sync.Once

var idbobjectstoreinterface js.Value

func IDBObjectStoreGetInterface() js.Value {

	singletonIDBObjectStore.Do(func() {

		var err error
		if idbobjectstoreinterface, err = baseobject.Get(js.Global(), "IDBObjectStore"); err != nil {
			idbobjectstoreinterface = js.Undefined()
		}
		baseobject.Register(idbobjectstoreinterface, func(v js.Value) (interface{}, error) {
			return IDBObjectStoreNewFromJSObject(v)
		})
		IDBRequestGetInterface()
	})
	return idbobjectstoreinterface
}

func IDBObjectStoreNewFromJSObject(obj js.Value) (IDBObjectStore, error) {
	var i IDBObjectStore
	var err error
	if ai := IDBObjectStoreGetInterface(); !ai.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(ai) {
				i.BaseObject = i.SetObject(obj)
			} else {
				err = ErrNotAnIDBObjectStore
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return i, err
}

func (i IDBObjectStore) addput(method string, value interface{}, key ...string) (IDBRequest, error) {

	var obj js.Value
	var request IDBRequest
	var err error
	var arrayJS []interface{}
	arrayJS = append(arrayJS, baseobject.GetJsValueOf(value))

	if len(key) > 0 {
		arrayJS = append(arrayJS, js.ValueOf(key[0]))
	}

	if obj, err = i.Call(method, arrayJS...); err == nil {

		request, err = IDBRequestNewFromJSObject(obj)
	}

	return request, err

}
func (i IDBObjectStore) Add(value interface{}, key ...string) (IDBRequest, error) {
	return i.addput("add", value, key...)
}

func (i IDBObjectStore) CreateIndex(index string, keyname string, option ...map[string]interface{}) (IDBIndex, error) {

	var obj js.Value
	var o IDBIndex
	var err error
	var arrayJS []interface{}

	arrayJS = append(arrayJS, index)
	arrayJS = append(arrayJS, keyname)

	if len(option) > 0 {
		arrayJS = append(arrayJS, js.ValueOf(option[0]))
	}

	if obj, err = i.Call("createIndex", arrayJS...); err == nil {
		o, err = IDBDIndexNewFromJSObject(obj)
	}

	return o, err

}

func (i IDBObjectStore) Clear() (IDBRequest, error) {
	var obj js.Value
	var request IDBRequest
	var err error
	if obj, err = i.Call("clear"); err == nil {
		request, err = IDBRequestNewFromJSObject(obj)
	}

	return request, err
}

func (i IDBObjectStore) Count() (IDBRequest, error) {
	var obj js.Value
	var request IDBRequest
	var err error
	if obj, err = i.Call("count"); err == nil {
		request, err = IDBRequestNewFromJSObject(obj)
	}

	return request, err
}

func (i IDBObjectStore) Delete(key interface{}) (IDBRequest, error) {
	var obj js.Value
	var request IDBRequest
	var err error
	var objkey js.Value

	if rangekey, ok := key.(IDBKeyRange); ok {
		objkey = rangekey.JSObject()
	} else {
		objkey = js.ValueOf(key)
	}
	if obj, err = i.Call("delete", objkey); err == nil {
		request, err = IDBRequestNewFromJSObject(obj)
	}

	return request, err
}

func (i IDBObjectStore) DeleteIndex(key string) error {

	var err error
	_, err = i.Call("deleteIndex", js.ValueOf(key))

	return err
}

func (i IDBObjectStore) get(method string, key interface{}) (IDBRequest, error) {
	var obj js.Value
	var request IDBRequest
	var err error
	var objkey js.Value

	if rangekey, ok := key.(IDBKeyRange); ok {
		objkey = rangekey.JSObject()
	} else {
		objkey = js.ValueOf(key)
	}

	if obj, err = i.Call("get", objkey); err == nil {
		request, err = IDBRequestNewFromJSObject(obj)
	}

	return request, err
}

func (i IDBObjectStore) Get(key interface{}) (IDBRequest, error) {
	return i.get("get", key)
}

func (i IDBObjectStore) getAll(method string, option ...interface{}) (IDBRequest, error) {
	var obj js.Value
	var request IDBRequest
	var err error
	var objquery js.Value
	var arrayJS []interface{}

	if len(option) > 1 {
		if rangequery, ok := option[0].(IDBKeyRange); ok {
			objquery = rangequery.JSObject()
		} else {
			objquery = js.ValueOf(option[0])
		}
		arrayJS = append(arrayJS, objquery)
	}
	if len(option) > 2 {
		if count, ok := option[0].(int); ok {
			arrayJS = append(arrayJS, js.ValueOf(count))
		}

	}

	if obj, err = i.Call(method, arrayJS...); err == nil {
		request, err = IDBRequestNewFromJSObject(obj)
	}

	return request, err
}

func (i IDBObjectStore) GetAll(option ...interface{}) (IDBRequest, error) {
	return i.getAll("getAll", option...)
}

func (i IDBObjectStore) GetAllKeys(option ...interface{}) (IDBRequest, error) {
	return i.getAll("getAllKeys", option...)
}

func (i IDBObjectStore) GetKey(key interface{}) (IDBRequest, error) {
	return i.get("getKey", key)
}

func (i IDBObjectStore) Index(indexname string) (IDBIndex, error) {
	var obj js.Value
	var o IDBIndex
	var err error

	if obj, err = i.Call("index", js.ValueOf(indexname)); err == nil {
		o, err = IDBDIndexNewFromJSObject(obj)
	}

	return o, err
}

func (i IDBObjectStore) Put(value interface{}, key ...string) (IDBRequest, error) {
	return i.addput("put", value, key...)
}

func (i IDBObjectStore) openCursorWithMethod(method string, options ...interface{}) (IDBRequest, error) {
	var obj js.Value
	var request IDBRequest
	var err error
	var objquery js.Value
	var arrayJS []interface{}

	if len(options) > 1 {
		if rangequery, ok := options[0].(IDBKeyRange); ok {
			objquery = rangequery.JSObject()
			arrayJS = append(arrayJS, objquery)
		}

	}

	if len(options) > 2 {
		if direction, ok := options[1].(string); ok {
			objquery = js.ValueOf(direction)
			arrayJS = append(arrayJS, objquery)
		}

	}

	if obj, err = i.Call("openCursor", arrayJS...); err == nil {
		request, err = IDBRequestNewFromJSObject(obj)
	}

	return request, err
}

func (i IDBObjectStore) OpenCursor(options ...interface{}) (IDBRequest, error) {

	return i.openCursorWithMethod("openCursor", options...)
}

func (i IDBObjectStore) OpenKeyCursor(options ...interface{}) (IDBRequest, error) {

	return i.openCursorWithMethod("openKeyCursor", options...)
}
