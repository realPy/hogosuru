package customevent

import (
	"sync"

	"github.com/realPy/jswasm/js"
)

var singleton sync.Once

var customeventinterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//JSCustomEvent JSCustomEvent struct
type JSCustomEvent struct {
	object js.Value
}

//GetJSInterface get teh JS interface of event
func GetJSInterface() *JSInterface {

	singleton.Do(func() {
		var customeventinstance JSInterface
		var err error
		if customeventinstance.objectInterface, err = js.Global().GetWithErr("CustomEvent"); err == nil {
			customeventinterface = &customeventinstance
		}
	})

	return customeventinterface
}

//NewJSCustomEvent Create a newJSEvent
func NewJSCustomEvent(message, detail string) (JSCustomEvent, error) {
	var event JSCustomEvent

	if eventi := GetJSInterface(); eventi != nil {
		event.object = eventi.objectInterface.New(js.ValueOf(message), js.ValueOf(map[string]interface{}{"detail": detail}))
		return event, nil
	}
	return event, ErrNotImplemented
}

//DispatchEvent to the object
func (j JSCustomEvent) DispatchEvent(obj js.Value) error {
	var err error
	_, err = obj.CallWithErr("dispatchEvent", j.object)
	return err
}
