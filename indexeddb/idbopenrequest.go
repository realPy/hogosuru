package indexeddb

// https://developer.mozilla.org/fr/docs/Web/API/IDBOpenDBRequest

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/event"
)

//IDBOpenRequest struct
type IDBOpenDBRequest struct {
	IDBRequest
}

type IDBOpenDBRequestFrom interface {
	IDBOpenDBRequest() IDBOpenDBRequest
}

func (i IDBOpenDBRequest) IDBOpenDBRequest() IDBOpenDBRequest {
	return i
}

var singletonIDBOpenRequest sync.Once

var idbopendbrequestinterface js.Value

func IDBOpenDBRequestGetInterface() js.Value {

	singletonIDBOpenRequest.Do(func() {

		var err error
		if idbopendbrequestinterface, err = js.Global().GetWithErr("IDBOpenDBRequest"); err != nil {
			idbopendbrequestinterface = js.Null()
		}

		baseobject.Register(idbopendbrequestinterface, func(v js.Value) (interface{}, error) {
			return IDBOpenDBRequestNewFromJSObject(v)
		})
	})
	return idbopendbrequestinterface
}

func IDBOpenDBRequestNewFromJSObject(obj js.Value) (IDBOpenDBRequest, error) {
	var i IDBOpenDBRequest
	var err error
	if ai := IDBOpenDBRequestGetInterface(); !ai.IsNull() {
		if obj.InstanceOf(ai) {
			i.BaseObject = i.SetObject(obj)
		} else {
			err = ErrNotAnIDBOpenDBRequest
		}
	} else {
		err = ErrNotImplemented
	}

	return i, err
}

func (i IDBOpenDBRequest) OnBlocked(handler func(e event.Event)) error {

	return i.AddEventListener("blocked", handler)
}

func (i IDBOpenDBRequest) OnUpgradeNeeded(handler func(e event.Event)) error {

	return i.AddEventListener("upgradeneeded", handler)
}
