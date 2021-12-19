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
	var eventi, obj js.Value
	var err error
	if eventi = GetInterface(); eventi.IsUndefined() {
		return event, ErrNotImplemented
	}
	if obj, err = baseobject.New(eventi, js.ValueOf(message), js.ValueOf(map[string]interface{}{"detail": baseobject.GetJsValueOf(detail)})); err != nil {
		return event, err
	}
	event.BaseObject = event.SetObject(obj)
	return event, nil
}

func NewFromJSObject(obj js.Value) (CustomEvent, error) {
	var c CustomEvent
	var bi js.Value
	if bi = GetInterface(); bi.IsUndefined() {
		return c, ErrNotImplemented
	}
	if obj.IsUndefined() || obj.IsNull() {
		return c, baseobject.ErrUndefinedValue
	}
	if !obj.InstanceOf(bi) {
		return c, ErrNotAnCustomEvent
	}
	c.BaseObject = c.SetObject(obj)
	return c, nil
}

func (c CustomEvent) Detail() (interface{}, error) {
	var obj js.Value
	var err error
	var i interface{}
	if obj, err = c.Get("detail"); err != nil {
		return i, err
	}
	return baseobject.GoValue(obj)
}
