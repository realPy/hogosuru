package indexeddb

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/event"
)

//IDBOpenRequest struct
type IDBOpenRequest struct {
	IDBRequest
}

var singletonIDBOpenRequest sync.Once

var idbopenrequestinterface js.Value

func IDBOpenRequestGetInterface() js.Value {

	singletonIDBRequest.Do(func() {

		var err error
		if idbopenrequestinterface, err = js.Global().GetWithErr("IDBOpenDBRequest"); err != nil {
			idbopenrequestinterface = js.Null()
		}
	})
	return idbrequestinterface
}

func IDBOpenRequestNewFromJSObject(obj js.Value) (IDBOpenRequest, error) {
	var i IDBOpenRequest
	var err error
	if ai := IDBOpenRequestGetInterface(); !ai.IsNull() {
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

func (i IDBOpenRequest) OnBlocked(handler func(e event.Event)) error {

	return i.AddEventListener("blocked", handler)
}

func (i IDBOpenRequest) OnUpgradeNeeded(handler func(e event.Event)) error {

	return i.AddEventListener("onupgradeneeded", handler)
}
