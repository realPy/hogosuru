package messageevent

// https://developer.mozilla.org/fr/docs/Web/API/MessageEvent

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/arraybuffer"
	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/blob"
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
		//instance object for autodiscovery
		arraybuffer.GetInterface()
		blob.GetInterface()
		baseobject.Register(messageeventinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return messageeventinterface
}

type MessageEvent struct {
	event.Event
}

type MessageEventFrom interface {
	MessageEvent() MessageEvent
}

func (m MessageEvent) MessageEvent() MessageEvent {
	return m
}

func NewFromJSObject(obj js.Value) (MessageEvent, error) {

	var message MessageEvent

	if mi := GetInterface(); !mi.IsNull() {
		if obj.InstanceOf(mi) {
			message.BaseObject = message.SetObject(obj)
			return message, nil
		}
	}

	return message, ErrNotAMessageEvent

}

func (m MessageEvent) Data() (interface{}, error) {

	var jsObject js.Value
	var globalObj interface{}
	var err error
	if jsObject, err = m.JSObject().GetWithErr("data"); err == nil {
		globalObj, err = baseobject.Discover(jsObject)
	}

	return globalObj, err
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
