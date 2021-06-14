package indexeddb

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/event"
	"github.com/realPy/hogosuru/eventtarget"
)

//IDBRequest struct
type IDBRequest struct {
	eventtarget.EventTarget
}

var singletonIDBRequest sync.Once

var idbrequestinterface js.Value

func IDBRequestGetInterface() js.Value {

	singletonIDBRequest.Do(func() {

		var err error
		if idbrequestinterface, err = js.Global().GetWithErr("IDBBRequest"); err != nil {
			idbrequestinterface = js.Null()
		}
	})
	return idbrequestinterface
}

func IDBRequestNewFromJSObject(obj js.Value) (IDBRequest, error) {
	var i IDBRequest
	var err error
	if ai := IDBRequestGetInterface(); !ai.IsNull() {
		if obj.InstanceOf(ai) {
			i.BaseObject = i.SetObject(obj)
			return i, nil
		} else {
			err = ErrNotAnIDBOpenRequest
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

			valueStr = obj.String()
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
