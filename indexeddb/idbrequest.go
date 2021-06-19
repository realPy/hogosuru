package indexeddb

// https://developer.mozilla.org/fr/docs/Web/API/IDBRequest

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/event"
	"github.com/realPy/hogosuru/eventtarget"
)

var classIDBRequest string = "IDBRequest"

//IDBRequest struct
type IDBRequest struct {
	eventtarget.EventTarget
}

var singletonIDBRequest sync.Once

var idbrequestinterface js.Value

func IDBRequestGetInterface() js.Value {

	singletonIDBRequest.Do(func() {

		var err error
		if idbrequestinterface, err = js.Global().GetWithErr(classIDBRequest); err != nil {
			idbrequestinterface = js.Null()
		}

		baseobject.Register(idbrequestinterface, func(v js.Value) (interface{}, error) {
			return IDBRequestNewFromJSObject(v)
		})
	})
	return idbrequestinterface
}

func IDBRequestNewFromJSObject(obj js.Value) (IDBRequest, error) {
	var i IDBRequest
	var err error
	if ai := IDBRequestGetInterface(); !ai.IsNull() {
		if obj.InstanceOf(ai) {
			i.BaseObject = i.SetObject(obj)
		} else {
			err = ErrNotAnIDBRequest
		}
	} else {
		err = ErrNotImplemented
	}

	return i, err
}

func (i IDBRequest) OnError(handler func(e event.Event)) error {

	return i.AddEventListener("error", handler)
}

func (i IDBRequest) OnSuccess(handler func(e event.Event)) error {

	return i.AddEventListener("success", handler)
}

func (i IDBRequest) getStringAttribute(attribute string) (string, error) {

	var err error
	var obj js.Value
	var valueStr = ""

	if obj, err = i.JSObject().GetWithErr(attribute); err == nil {
		if obj.IsNull() {
			err = baseobject.ErrNotAnObject

		} else {

			valueStr, _ = baseobject.ToStringWithErr(obj)
		}
	}

	return valueStr, err

}

func (i IDBRequest) Error() (string, error) {
	return i.getStringAttribute("error")
}

func (i IDBRequest) ReadyState() (string, error) {
	return i.getStringAttribute("readystate")
}

func (i IDBRequest) getObjectAttribute(attribute string) (baseobject.BaseObject, error) {
	var err error
	var obj js.Value
	var bobj baseobject.BaseObject

	if obj, err = i.JSObject().GetWithErr(attribute); err == nil {
		if obj.IsNull() {
			err = baseobject.ErrNotAnObject

		} else {

			bobj, err = baseobject.NewFromJSObject(obj)
		}
	}

	return bobj, err
}

func (i IDBRequest) Result() (baseobject.BaseObject, error) {
	return i.getObjectAttribute("result")
}

func (i IDBRequest) Source() (baseobject.BaseObject, error) {
	return i.getObjectAttribute("source")
}

func (i IDBRequest) Transaction() (IDBTransaction, error) {
	var err error
	var obj baseobject.BaseObject
	var it IDBTransaction

	if obj, err = i.getObjectAttribute("transaction"); err == nil {
		it, err = IDBTransactionNewFromJSObject(obj.JSObject())

	}
	return it, err
}
