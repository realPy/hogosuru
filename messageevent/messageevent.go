package messageevent

// https://developer.mozilla.org/fr/docs/Web/API/MessageEvent

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/event"
)

var singleton sync.Once

var messageeventinterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//GetJSInterface get the JS interface of formdata
func GetJSInterface() *JSInterface {

	singleton.Do(func() {
		var messageeventinstance JSInterface
		var err error
		if messageeventinstance.objectInterface, err = js.Global().GetWithErr("MessageEvent"); err == nil {
			messageeventinterface = &messageeventinstance
		}
	})

	return messageeventinterface
}

type MessageEvent struct {
	event.Event
}

func NewFromJSObject(obj js.Value) (MessageEvent, error) {
	var m MessageEvent

	if mi := GetJSInterface(); mi != nil {
		if obj.InstanceOf(mi.objectInterface) {
			m.Object = m.SetObject(obj)
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
