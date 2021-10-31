package indexeddb

// https://developer.mozilla.org/fr/docs/Web/API/IDBTransaction

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/domstringlist"
	"github.com/realPy/hogosuru/event"
	"github.com/realPy/hogosuru/eventtarget"
)

//IDBTransaction struct
type IDBTransaction struct {
	eventtarget.EventTarget
}

type IDBTransactionFrom interface {
	IDBTransaction_() IDBTransaction
}

func (i IDBTransaction) IDBTransaction_() IDBTransaction {
	return i
}

var singletonIDBTransaction sync.Once

var idbtransactioninterface js.Value

func IDBTransactionGetInterface() js.Value {

	singletonIDBTransaction.Do(func() {

		var err error
		if idbtransactioninterface, err = baseobject.Get(js.Global(), "IDBTransaction"); err != nil {
			idbtransactioninterface = js.Undefined()
		}
		baseobject.Register(idbtransactioninterface, func(v js.Value) (interface{}, error) {
			return IDBTransactionNewFromJSObject(v)
		})

	})
	return idbtransactioninterface
}

func IDBTransactionNewFromJSObject(obj js.Value) (IDBTransaction, error) {
	var i IDBTransaction
	var err error
	if ai := IDBTransactionGetInterface(); !ai.IsUndefined() {
		if obj.IsUndefined() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(ai) {
				i.BaseObject = i.SetObject(obj)
			} else {
				err = ErrNotAnIDBTransaction
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return i, err
}

func (i IDBTransaction) Abort() error {
	var err error
	_, err = i.Call("abort")
	return err
}

func (i IDBTransaction) ObjectStore(name string) (IDBObjectStore, error) {
	var err error
	var obj js.Value
	var s IDBObjectStore
	if obj, err = i.Call("objectStore", js.ValueOf(name)); err == nil {
		s, err = IDBObjectStoreNewFromJSObject(obj)
	}

	return s, err
}

func (i IDBTransaction) DB() (IDBDatabase, error) {
	var err error
	var obj js.Value
	var t IDBDatabase
	if obj, err = i.Get("database"); err == nil {
		t, err = IDBDatabaseNewFromJSObject(obj)
	}

	return t, err
}

func (i IDBTransaction) Mode() (string, error) {
	return i.GetAttributeString("mode")
}

func (i IDBTransaction) ObjectStoreNames() (domstringlist.DOMStringList, error) {
	var err error
	var obj js.Value
	var list domstringlist.DOMStringList

	if obj, err = i.Get("objectStoreNames"); err == nil {
		list, err = domstringlist.NewFromJSObject(obj)
	}

	return list, err
}

func (i IDBTransaction) Error() (string, error) {
	return i.GetAttributeString("error")
}

func (i IDBTransaction) OnAbort(handler func(e event.Event)) (js.Func, error) {

	return i.AddEventListener("abort", handler)
}

func (i IDBOpenDBRequest) OnComplete(handler func(e event.Event)) (js.Func, error) {

	return i.AddEventListener("complete", handler)
}

func (i IDBOpenDBRequest) OnError(handler func(e event.Event)) (js.Func, error) {

	return i.AddEventListener("error", handler)
}
