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
		if messageeventinterface, err = baseobject.Get(js.Global(), "MessageEvent"); err != nil {
			messageeventinterface = js.Undefined()
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
	MessageEvent_() MessageEvent
}

func (m MessageEvent) MessageEvent_() MessageEvent {
	return m
}

func NewFromJSObject(obj js.Value) (MessageEvent, error) {

	var message MessageEvent
	var err error
	if mi := GetInterface(); !mi.IsUndefined() {
		if obj.IsUndefined() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(mi) {
				message.BaseObject = message.SetObject(obj)
			} else {
				err = ErrNotAMessageEvent
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return message, err

}

func (m MessageEvent) Data() (interface{}, error) {

	var jsObject js.Value
	var globalObj interface{}
	var err error
	if jsObject, err = m.Get("data"); err == nil {
		globalObj = baseobject.GoValue(jsObject)
	}

	return globalObj, err
}

func (m MessageEvent) Source() (js.Value, error) {
	return m.Get("source")
}
func (m MessageEvent) Origin() (string, error) {
	var err error
	var originObject js.Value

	if originObject, err = m.Get("origin"); err == nil {
		return originObject.String(), nil
	}
	return "", err
}

func (m MessageEvent) LastEventId() (string, error) {
	var err error
	var originObject js.Value

	if originObject, err = m.Get("lastEventId"); err == nil {
		return originObject.String(), nil
	}
	return "", err
}
