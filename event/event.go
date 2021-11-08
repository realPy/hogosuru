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
func New(typeevent string, init ...map[string]interface{}) (Event, error) {
	var e Event
	var obj js.Value
	var err error
	var arrayJS []interface{}

	if ei := GetInterface(); !ei.IsUndefined() {
		arrayJS = append(arrayJS, js.ValueOf(typeevent))
		if len(init) > 0 {
			arrayJS = append(arrayJS, js.ValueOf(init[0]))
		}
		if obj, err = baseobject.New(ei, arrayJS...); err == nil {
			e.BaseObject = e.SetObject(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return e, err
}

func NewFromJSObject(obj js.Value) (Event, error) {
	var e Event
	var err error
	if eventi := GetInterface(); !eventi.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(eventi) {
				e.BaseObject = e.SetObject(obj)

			} else {
				err = ErrNotAnEvent
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return e, err
}

func (e Event) Target() (interface{}, error) {
	var err error
	var obj js.Value
	var bobj interface{}

	if obj, err = e.Get("target"); err == nil {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {
			bobj, err = baseobject.Discover(obj)
		}

	}
	return bobj, err
}
func (e Event) CurrentTarget() (interface{}, error) {
	var err error
	var obj js.Value
	var bobj interface{}

	if obj, err = e.Get("currentTarget"); err == nil {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {
			bobj, err = baseobject.Discover(obj)
		}

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

func (e Event) Bubbles() (bool, error) {
	return e.GetAttributeBool("bubbles")
}

func (e Event) Cancelable() (bool, error) {
	return e.GetAttributeBool("cancelable")
}

func (e Event) Composed() (bool, error) {
	return e.GetAttributeBool("composed")
}

func (e Event) EventPhase() (int, error) {
	return e.GetAttributeInt("eventPhase")
}

func (e Event) Type() (string, error) {
	return e.GetAttributeString("type")
}

func (e Event) IsTrusted() (bool, error) {
	return e.GetAttributeBool("isTrusted")
}
