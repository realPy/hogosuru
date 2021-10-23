package event

// partial implemented
// https://developer.mozilla.org/fr/docs/Web/API/Event

import (
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
)

var singleton sync.Once

var eventinterface js.Value

//Event Event struct
type Event struct {
	baseobject.BaseObject
}

type EventFrom interface {
	Event_() Event
}

func (e Event) Event_() Event {
	return e
}

//GetInterface get the JS interface of event
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if eventinterface, err = baseobject.Get(js.Global(), "Event"); err != nil {
			eventinterface = js.Undefined()
		}
		baseobject.Register(eventinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

	})

	return eventinterface
}

//New Create a event
func New(message string) (Event, error) {
	var event Event

	if eventi := GetInterface(); !eventi.IsUndefined() {
		event.BaseObject = event.SetObject(eventi.New(js.ValueOf(message)))
		return event, nil
	}
	return event, ErrNotImplemented
}

func NewFromJSObject(obj js.Value) (Event, error) {
	var e Event

	if eventi := GetInterface(); !eventi.IsUndefined() {
		if obj.InstanceOf(eventi) {
			e.BaseObject = e.SetObject(obj)
			return e, nil
		}
	}

	return e, ErrNotAnEvent
}

func (e Event) Target() (interface{}, error) {
	var err error
	var obj js.Value
	var bobj interface{}

	if obj, err = e.Get("target"); err == nil {

		bobj, err = baseobject.Discover(obj)
	}
	return bobj, err
}
func (e Event) CurrentTarget() (interface{}, error) {
	var err error
	var obj js.Value
	var bobj interface{}

	if obj, err = e.Get("currentTarget"); err == nil {

		bobj, err = baseobject.Discover(obj)
	}
	return bobj, err
}

func (e Event) PreventDefault() error {
	var err error
	_, err = e.Call("preventDefault")

	return err
}

func (e Event) StopImmediatePropagation() error {
	var err error
	_, err = e.Call("stopImmediatePropagation")

	return err
}

func (e Event) StopPropagation() error {
	var err error
	_, err = e.Call("stopPropagation")

	return err
}
