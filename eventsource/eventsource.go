package eventsource

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/event"
	"github.com/realPy/hogosuru/eventtarget"
	"github.com/realPy/hogosuru/messageevent"
)

var singleton sync.Once

var sseinterface js.Value

//GetJSInterface Get the Event Source Interface If nil browser doesn't implement it
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if sseinterface, err = baseobject.Get(js.Global(), "EventSource"); err != nil {
			sseinterface = js.Undefined()
		}

		baseobject.Register(sseinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return sseinterface
}

//EventSource struct
type EventSource struct {
	eventtarget.EventTarget
}

type EventSourceFrom interface {
	EventSource_() EventSource
}

func (e EventSource) EventSource_() EventSource {
	return e
}

func NewFromJSObject(obj js.Value) (EventSource, error) {
	var e EventSource
	var err error
	if eventsourcei := GetInterface(); !eventsourcei.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(eventsourcei) {

				e.BaseObject = e.SetObject(obj)
			} else {
				err = ErrNotAnEventSource
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return e, err
}

func New(url string, opts ...interface{}) (EventSource, error) {
	var arrayJS []interface{}
	var e EventSource
	var err error
	arrayJS = append(arrayJS, url)
	for _, value := range opts {
		arrayJS = append(arrayJS, baseobject.GetJsValueOf(value))
	}
	if eventsourcei := GetInterface(); !eventsourcei.IsUndefined() {
		promisefetchobj := eventsourcei.Invoke(arrayJS...)
		e.BaseObject = e.SetObject(promisefetchobj)
	} else {
		err = ErrNotImplemented
	}
	return e, err

}

func (e EventSource) ReadyState() (int, error) {
	var err error
	var obj js.Value
	if obj, err = e.Get("readyState"); err == nil {

		return obj.Int(), nil
	}
	return 0, err
}

func (e EventSource) Url() (string, error) {

	var err error
	var obj js.Value
	if obj, err = e.Get("url"); err == nil {

		return obj.String(), nil
	}
	return "", err

}

func (e EventSource) WithCredentials() (bool, error) {
	return e.GetAttributeBool("withCredentials")
}

func (e EventSource) setHandler(jshandlername string, handler func(e event.Event)) {

	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		if e, err := event.NewFromJSObject(args[0]); err == nil {
			handler(e)
		}

		return nil
	})

	e.JSObject().Set(jshandlername, jsfunc)
}

//SetOnOpen Set onOpen Handler
func (e EventSource) SetOnOpen(handler func(e event.Event)) {

	e.setHandler("onopen", func(e event.Event) {
		handler(e)
	})
}

//SetOnClose Set onClose Handler
func (e EventSource) SetOnClose(handler func(e event.Event)) {
	e.setHandler("onclose", func(e event.Event) {
		handler(e)
	})
}

//SetOnClose Set onClose Handler
func (e EventSource) SetOnError(handler func(e event.Event)) {
	e.setHandler("onerror", func(e event.Event) {
		handler(e)
	})
}

//SetOnClose Set onClose Handler
func (e EventSource) SetOnMessage(handler func(e messageevent.MessageEvent)) {
	e.setHandler("onmessage", func(e event.Event) {

		if obj, err := baseobject.Discover(e.JSObject()); err == nil {

			if m, ok := obj.(messageevent.MessageEventFrom); ok {
				handler(m.MessageEvent_())
			}
		}
	})
}

//OnOpen Set onOpen Handler
func (e EventSource) OnOpen(handler func(e event.Event)) (js.Func, error) {

	return e.AddEventListener("open", handler)
}

//OnClose Set onClose Handler
func (e EventSource) OnClose(handler func(e event.Event)) (js.Func, error) {

	return e.AddEventListener("close", handler)
}

//OnError Set onError Handler
func (e EventSource) OnError(handler func(e event.Event)) (js.Func, error) {

	return e.AddEventListener("error", handler)
}
