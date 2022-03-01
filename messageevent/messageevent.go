package messageevent

// https://developer.mozilla.org/fr/docs/Web/API/MessageEvent

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/arraybuffer"
	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/blob"
	"github.com/realPy/hogosuru/event"
	"github.com/realPy/hogosuru/initinterface"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

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

func New(typeevent string, init ...map[string]interface{}) (MessageEvent, error) {

	var m MessageEvent
	var obj js.Value
	var err error
	var arrayJS []interface{}

	if pei := GetInterface(); !pei.IsUndefined() {
		arrayJS = append(arrayJS, js.ValueOf(typeevent))
		if len(init) > 0 {
			arrayJS = append(arrayJS, js.ValueOf(init[0]))
		}
		if obj, err = baseobject.New(pei, arrayJS...); err == nil {
			m.BaseObject = m.SetObject(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return m, err
}

func NewFromJSObject(obj js.Value) (MessageEvent, error) {

	var message MessageEvent
	var err error
	if mi := GetInterface(); !mi.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
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
		globalObj, err = baseobject.GoValue(jsObject)
	}

	return globalObj, err
}

func (m MessageEvent) Source() (interface{}, error) {
	var obj js.Value
	var err error
	var i interface{}

	if obj, err = m.Get("source"); err == nil {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {
			i, err = baseobject.Discover(obj)
		}

	}

	return i, err
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
