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

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//Event Event struct
type Event struct {
	baseobject.BaseObject
}

//GetInterface get the JS interface of event
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if eventinterface, err = js.Global().GetWithErr("Event"); err != nil {
			eventinterface = js.Null()
		}

	})

	return eventinterface
}

//New Create a event
func New(message string) (Event, error) {
	var event Event

	if eventi := GetInterface(); !eventi.IsNull() {
		event.BaseObject = event.SetObject(eventi.New(js.ValueOf(message)))
		return event, nil
	}
	return event, ErrNotImplemented
}

func NewFromJSObject(obj js.Value) (Event, error) {
	var e Event

	if eventi := GetInterface(); !eventi.IsNull() {
		if obj.InstanceOf(eventi) {
			e.BaseObject = e.SetObject(obj)
			return e, nil
		}
	}

	return e, ErrNotAnEvent
}

func (e Event) PreventDefault() error {
	var err error
	_, err = e.JSObject().CallWithErr("preventDefault")

	return err
}

func (e Event) StopImmediatePropagation() error {
	var err error
	_, err = e.JSObject().CallWithErr("stopImmediatePropagation")

	return err
}

func (e Event) StopPropagation() error {
	var err error
	_, err = e.JSObject().CallWithErr("stopPropagation")

	return err
}
