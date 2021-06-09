package customevent

// https://developer.mozilla.org/fr/docs/Web/API/CustomEvent
import (
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/event"
)

var singleton sync.Once

var customeventinterface js.Value

//CustomEvent CustomEvent struct
type CustomEvent struct {
	event.Event
}

//GetJSInterface get teh JS interface of event
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if customeventinterface, err = js.Global().GetWithErr("CustomEvent"); err != nil {
			customeventinterface = js.Null()
		}

	})

	return customeventinterface
}

//New Create a CustomEvent
func New(message, detail string) (CustomEvent, error) {
	var event CustomEvent

	if eventi := GetInterface(); !eventi.IsNull() {
		event.BaseObject = event.SetObject(eventi.New(js.ValueOf(message), js.ValueOf(map[string]interface{}{"detail": detail})))
		return event, nil
	}
	return event, ErrNotImplemented
}
