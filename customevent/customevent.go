package customevent

// https://developer.mozilla.org/fr/docs/Web/API/CustomEvent
import (
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/event"
)

var singleton sync.Once

var customeventinterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//CustomEvent CustomEvent struct
type CustomEvent struct {
	event.Event
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

//New Create a CustomEvent
func New(message, detail string) (CustomEvent, error) {
	var event CustomEvent

	if eventi := GetJSInterface(); eventi != nil {
		event.BaseObject = event.SetObject(eventi.objectInterface.New(js.ValueOf(message), js.ValueOf(map[string]interface{}{"detail": detail})))
		return event, nil
	}
	return event, ErrNotImplemented
}
