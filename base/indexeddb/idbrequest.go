package indexeddb

// https://developer.mozilla.org/fr/docs/Web/API/IDBRequest

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/domexception"
	"github.com/realPy/hogosuru/base/event"
	"github.com/realPy/hogosuru/base/eventtarget"
)

var classIDBRequest string = "IDBRequest"

// IDBRequest struct
type IDBRequest struct {
	eventtarget.EventTarget
}

type IDBRequestFrom interface {
	IDBRequest_() IDBRequest
}

func (i IDBRequest) IDBRequest_() IDBRequest {
	return i
}

var singletonIDBRequest sync.Once

var idbrequestinterface js.Value

func IDBRequestGetInterface() js.Value {

	singletonIDBRequest.Do(func() {

		var err error
		if idbrequestinterface, err = baseobject.Get(js.Global(), classIDBRequest); err != nil {
			idbrequestinterface = js.Undefined()
		}

		baseobject.Register(idbrequestinterface, func(v js.Value) (interface{}, error) {
			return IDBRequestNewFromJSObject(v)
		})
		IDBTransactionGetInterface()
		IDBDatabaseGetInterface()
		IDBCursorGetInterface()
	})
	return idbrequestinterface
}

func IDBRequestNewFromJSObject(obj js.Value) (IDBRequest, error) {
	var i IDBRequest
	var err error
	if ai := IDBRequestGetInterface(); !ai.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(ai) {
				i.BaseObject = i.SetObject(obj)
			} else {
				err = ErrNotAnIDBRequest
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return i, err
}

func (i IDBRequest) OnError(handler func(e event.Event)) (js.Func, error) {

	return i.AddEventListener("error", handler)
}

func (i IDBRequest) OnSuccess(handler func(e event.Event)) (js.Func, error) {

	return i.AddEventListener("success", handler)
}

func (i IDBRequest) Error() (domexception.DomException, error) {
	var err error
	var obj js.Value
	var e domexception.DomException
	if obj, err = i.Get("error"); err == nil {
		e, err = domexception.NewFromJSObject(obj)
	}
	return e, err
}

func (i IDBRequest) ReadyState() (string, error) {
	return i.GetAttributeString("readyState")
}

func (i IDBRequest) Result() (interface{}, error) {

	var err error
	var obj js.Value
	var ret interface{}

	if obj, err = i.Get("result"); err == nil {

		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {
			ret, err = baseobject.Discover(obj)
		}
	}

	return ret, err
}

func (i IDBRequest) Source() (interface{}, error) {

	var err error
	var obj js.Value
	var ret interface{}

	if obj, err = i.Get("source"); err == nil {

		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {
			ret, err = baseobject.Discover(obj)
		}

	}

	return ret, err

}

func (i IDBRequest) Transaction() (IDBTransaction, error) {

	var err error
	var obj js.Value
	var it IDBTransaction
	var ret interface{}

	if obj, err = i.Get("transaction"); err == nil {

		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {
			if ret, err = baseobject.Discover(obj); err == nil {

				if tfrom, ok := ret.(IDBTransactionFrom); ok {

					it = tfrom.IDBTransaction_()

				} else {
					err = ErrNotAnIDBTransaction
				}

			}
		}
	}
	return it, err
}
