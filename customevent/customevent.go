package customevent

// https://developer.mozilla.org/fr/docs/Web/API/CustomEvent
import (
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/event"
)

var singleton sync.Once

var customeventinterface js.Value

//CustomEvent CustomEvent struct
type CustomEvent struct {
	event.Event
}

type CustomEventFrom interface {
	CustomEvent_() CustomEvent
}

func (c CustomEvent) CustomEvent_() CustomEvent {
	return c
}

//GetInterface get teh JS interface of event
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if customeventinterface, err = baseobject.Get(js.Global(), "CustomEvent"); err != nil {
			customeventinterface = js.Undefined()
		}

		baseobject.Register(customeventinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

	})

	return customeventinterface
}

//New Create a CustomEvent
func New(message, detail interface{}) (CustomEvent, error) {
	var event CustomEvent
	var jsObj js.Value
	var obj js.Value
	var err error
	if objGo, ok := detail.(baseobject.ObjectFrom); ok {
		jsObj = objGo.JSObject()
	} else {
		jsObj = js.ValueOf(detail)
	}

	if eventi := GetInterface(); !eventi.IsUndefined() {

		if obj, err = baseobject.New(eventi, js.ValueOf(message), js.ValueOf(map[string]interface{}{"detail": jsObj})); err == nil {
			event.BaseObject = event.SetObject(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return event, err
}

func NewFromJSObject(obj js.Value) (CustomEvent, error) {
	var c CustomEvent
	var err error

	if bi := GetInterface(); !bi.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(bi) {
				c.BaseObject = c.SetObject(obj)

			} else {
				err = ErrNotAnCustomEvent
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return c, err
}

func (c CustomEvent) Detail() (interface{}, error) {
	var obj js.Value
	var err error
	var i interface{}

	if obj, err = c.Get("detail"); err == nil {
		i = baseobject.GoValue(obj)
	}
	return i, err
}
