package customevent

import (
	"sync"

	"github.com/realPy/hogosuru/event"
	"github.com/realPy/hogosuru/js"
)

var singleton sync.Once

var customeventinterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//JSCustomEvent JSCustomEvent struct
type JSCustomEvent struct {
	event.JSEvent
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

//New Create a newJSEvent
func New(message, detail string) (JSCustomEvent, error) {
	var event JSCustomEvent

	if eventi := GetJSInterface(); eventi != nil {
		event.Object = event.SetObject(eventi.objectInterface.New(js.ValueOf(message), js.ValueOf(map[string]interface{}{"detail": detail})))
		return event, nil
	}
	return event, ErrNotImplemented
}
