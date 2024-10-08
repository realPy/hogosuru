package indexeddb

// https://developer.mozilla.org/fr/docs/Web/API/IDBOpenDBRequest

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/event"
)

// IDBOpenRequest struct
type IDBOpenDBRequest struct {
	IDBRequest
}

type IDBOpenDBRequestFrom interface {
	IDBOpenDBRequest_() IDBOpenDBRequest
}

func (i IDBOpenDBRequest) IDBOpenDBRequest_() IDBOpenDBRequest {
	return i
}

var singletonIDBOpenRequest sync.Once

var idbopendbrequestinterface js.Value

func IDBOpenDBRequestGetInterface() js.Value {

	singletonIDBOpenRequest.Do(func() {

		var err error
		if idbopendbrequestinterface, err = baseobject.Get(js.Global(), "IDBOpenDBRequest"); err != nil {
			idbopendbrequestinterface = js.Undefined()
		}

		baseobject.Register(idbopendbrequestinterface, func(v js.Value) (interface{}, error) {
			return IDBOpenDBRequestNewFromJSObject(v)
		})
		IDBRequestGetInterface()
	})
	return idbopendbrequestinterface
}

func IDBOpenDBRequestNewFromJSObject(obj js.Value) (IDBOpenDBRequest, error) {
	var i IDBOpenDBRequest
	var err error
	if ai := IDBOpenDBRequestGetInterface(); !ai.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(ai) {
				i.BaseObject = i.SetObject(obj)
			} else {
				err = ErrNotAnIDBOpenDBRequest
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return i, err
}

func (i IDBOpenDBRequest) OnBlocked(handler func(e event.Event)) (js.Func, error) {

	return i.AddEventListener("blocked", handler)
}

func (i IDBOpenDBRequest) OnUpgradeNeeded(handler func(e event.Event)) (js.Func, error) {

	return i.AddEventListener("upgradeneeded", handler)
}
