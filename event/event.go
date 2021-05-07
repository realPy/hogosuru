package event

import (
	"sync"

	"github.com/realPy/jswasm/js"
	"github.com/realPy/jswasm/object"
)

var singleton sync.Once

var eventinterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//JSEvent JSEvent struct
type JSEvent struct {
	object.Object
}

//GetJSInterface get teh JS interface of event
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

//NewJSEvent Create a newJSEvent
func NewJSEvent(message string) (JSEvent, error) {
	var event JSEvent

	if eventi := GetJSInterface(); eventi != nil {
		event.Object = event.SetObject(eventi.objectInterface.New(js.ValueOf(message)))
		return event, nil
	}
	return event, ErrNotImplemented
}

//DispatchEvent to the object
func (j JSEvent) DispatchEvent(obj js.Value) error {
	var err error
	_, err = obj.CallWithErr("dispatchEvent", j.JSObject())
	return err
}
