package eventtarget

// https://developer.mozilla.org/fr/docs/Web/API/EventTarget/EventTarget
import (
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/event"
)

var singleton sync.Once

var eventtargetinterface js.Value

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//GetJSInterface get teh JS interface of broadcast channel
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if eventtargetinterface, err = js.Global().GetWithErr("EventTarget"); err != nil {
			eventtargetinterface = js.Null()
		}

	})

	baseobject.Register(eventtargetinterface, func(v js.Value) (interface{}, error) {
		return NewFromJSObject(v)
	})

	return eventtargetinterface
}

type EventTarget struct {
	event.Event
	registerFunc map[string]js.Func
}

func New() (EventTarget, error) {

	var e EventTarget

	if eti := GetInterface(); !eti.IsNull() {
		e.BaseObject = e.SetObject(eti.New())
		e.registerFunc = make(map[string]js.Func)
		return e, nil
	}
	return e, ErrNotImplemented
}

func NewFromJSObject(obj js.Value) (EventTarget, error) {
	var e EventTarget

	if eti := GetInterface(); !eti.IsNull() {
		if obj.InstanceOf(eti) {
			e.BaseObject = e.SetObject(obj)
			e.registerFunc = make(map[string]js.Func)
			return e, nil
		}
	}

	return e, ErrNotAnEventTarget
}

func (e EventTarget) AddEventListener(name string, handler func(e event.Event)) error {

	var err error
	if handler != nil {
		cb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			if e, err := event.NewFromJSObject(args[0]); err == nil {
				handler(e)
			}
			return nil
		})
		if e.registerFunc == nil {
			e.registerFunc = make(map[string]js.Func)
		}
		e.registerFunc[name] = cb
		_, err = e.JSObject().CallWithErr("addEventListener", js.ValueOf(name), cb)
	}

	return err
}

func (e EventTarget) RemoveEventListener(name string, typeevent string) error {
	var err error
	_, err = e.JSObject().CallWithErr("removeEventListener", typeevent, e.registerFunc[name])
	cb := e.registerFunc[name]
	delete(e.registerFunc, name)
	cb.Release()
	return err
}

func (e EventTarget) DispatchEvent(event event.Event) error {
	var err error
	_, err = e.JSObject().CallWithErr("dispatchEvent", event.JSObject())
	return err
}
