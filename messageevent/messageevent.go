package messageevent

// https://developer.mozilla.org/fr/docs/Web/API/MessageEvent

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/event"
)

var singleton sync.Once

var messageeventinterface js.Value

//GetJSInterface get the JS interface of formdata
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if messageeventinterface, err = js.Global().GetWithErr("MessageEvent"); err != nil {
			messageeventinterface = js.Null()
		}

	})

	return messageeventinterface
}

type MessageEvent struct {
	event.Event
}

func NewFromJSObject(obj js.Value) (MessageEvent, error) {
	var m MessageEvent

	if mi := GetInterface(); !mi.IsNull() {
		if obj.InstanceOf(mi) {
			m.BaseObject = m.SetObject(obj)
			return m, nil

		}
	}

	return m, ErrNotAMessageEvent
}

func (m MessageEvent) Data() (js.Value, error) {
	return m.JSObject().GetWithErr("data")
}

func (m MessageEvent) Source() (js.Value, error) {
	return m.JSObject().GetWithErr("source")
}
func (m MessageEvent) Origin() (string, error) {
	var err error
	var originObject js.Value

	if originObject, err = m.JSObject().GetWithErr("origin"); err == nil {
		return originObject.String(), nil
	}
	return "", err
}

func (m MessageEvent) LastEventId() (string, error) {
	var err error
	var originObject js.Value

	if originObject, err = m.JSObject().GetWithErr("lastEventId"); err == nil {
		return originObject.String(), nil
	}
	return "", err
}
