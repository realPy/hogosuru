package event

// partial implemented
// https://developer.mozilla.org/fr/docs/Web/API/Event

import (
	"sync"

	"github.com/realPy/hogosuru/js"
	"github.com/realPy/hogosuru/object"
)

var singleton sync.Once

var eventinterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//Event Event struct
type Event struct {
	object.Object
}

//GetJSInterface get the JS interface of event
func GetJSInterface() *JSInterface {

	singleton.Do(func() {
		var eventinstance JSInterface
		var err error
		if eventinstance.objectInterface, err = js.Global().GetWithErr("Event"); err == nil {
			eventinterface = &eventinstance
		}
	})

	return eventinterface
}

//New Create a event
func New(message string) (Event, error) {
	var event Event

	if eventi := GetJSInterface(); eventi != nil {
		event.Object = event.SetObject(eventi.objectInterface.New(js.ValueOf(message)))
		return event, nil
	}
	return event, ErrNotImplemented
}

func NewFromJSObject(obj js.Value) (Event, error) {
	var e Event

	if object.String(obj) == "[object Event]" {
		e.Object = e.SetObject(obj)
		return e, nil
	}

	return e, ErrNotAnEvent
}

//DispatchEvent to the object
func (e Event) DispatchEvent(obj js.Value) error {
	var err error
	_, err = obj.CallWithErr("dispatchEvent", e.JSObject())
	return err
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
