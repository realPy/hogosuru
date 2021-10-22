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

//GetJSInterface get teh JS interface of broadcast channel
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if eventtargetinterface, err = baseobject.Get(js.Global(), "EventTarget"); err != nil {
			eventtargetinterface = js.Undefined()
		}

		baseobject.Register(eventtargetinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return eventtargetinterface
}

type EventTarget struct {
	event.Event
}

type EventTargetFrom interface {
	EventTarget_() EventTarget
}

func (e EventTarget) EventTarget_() EventTarget {
	return e
}

func New() (EventTarget, error) {

	var e EventTarget

	if eti := GetInterface(); !eti.IsUndefined() {
		e.BaseObject = e.SetObject(eti.New())
		return e, nil
	}
	return e, ErrNotImplemented
}

func NewFromJSObject(obj js.Value) (EventTarget, error) {
	var e EventTarget

	if eti := GetInterface(); !eti.IsUndefined() {
		if obj.InstanceOf(eti) {
			e.BaseObject = e.SetObject(obj)
			return e, nil
		}
	}

	return e, ErrNotAnEventTarget
}

func (e EventTarget) AddEventListener(name string, handler func(e event.Event)) (js.Func, error) {

	var err error
	var cb js.Func
	if handler != nil {
		cb = js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			if e, err := event.NewFromJSObject(args[0]); err == nil {
				handler(e)
			}
			return nil
		})

		_, err = e.JSObject().CallWithErr("addEventListener", js.ValueOf(name), cb)
	}

	return cb, err
}

func (e EventTarget) RemoveEventListener(f js.Func, typeevent string) error {
	var err error
	_, err = e.JSObject().CallWithErr("removeEventListener", typeevent, f)
	f.Release()
	return err
}

func (e EventTarget) DispatchEvent(event event.Event) error {
	var err error
	_, err = e.JSObject().CallWithErr("dispatchEvent", event.JSObject())
	return err
}
