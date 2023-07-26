package eventsource

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/event"
	"github.com/realPy/hogosuru/base/eventtarget"
	"github.com/realPy/hogosuru/base/messageevent"
)

var singleton sync.Once

var sseinterface js.Value

// GetJSInterface Get the Event Source Interface If nil browser doesn't implement it
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if sseinterface, err = baseobject.Get(js.Global(), "EventSource"); err != nil {
			sseinterface = js.Undefined()
		}

		messageevent.GetInterface()
		baseobject.Register(sseinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return sseinterface
}

// EventSource struct
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
	var obj js.Value

	arrayJS = append(arrayJS, url)
	for _, value := range opts {
		arrayJS = append(arrayJS, baseobject.GetJsValueOf(value))
	}

	if eventsourcei := GetInterface(); !eventsourcei.IsUndefined() {

		if obj, err = baseobject.New(eventsourcei, arrayJS...); err == nil {
			e.BaseObject = e.SetObject(obj)
		}

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

func (e EventSource) Close() error {
	var err error
	_, err = e.Call("close")
	return err
}

func (e EventSource) WithCredentials() (bool, error) {
	return e.GetAttributeBool("withCredentials")
}

func (sse EventSource) setHandler(jshandlername string, handler func(e event.Event)) {

	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		if e, err := event.NewFromJSObject(args[0]); err == nil {
			handler(e)
		}

		return nil
	})

	sse.JSObject().Set(jshandlername, jsfunc)
}

// SetOnOpen Set onOpen Handler
func (sse EventSource) SetOnOpen(handler func(e event.Event)) {

	sse.setHandler("onopen", func(e event.Event) {
		handler(e)
	})
}

// SetOnClose Set onClose Handler
func (sse EventSource) SetOnClose(handler func(e event.Event)) {
	sse.setHandler("onclose", func(e event.Event) {
		handler(e)
	})
}

// SetOnClose Set onClose Handler
func (sse EventSource) SetOnError(handler func(e event.Event)) {
	sse.setHandler("onerror", func(e event.Event) {
		handler(e)
	})
}

// SetOnClose Set onClose Handler
func (sse EventSource) SetOnMessage(handler func(e messageevent.MessageEvent)) {
	sse.setHandler("onmessage", func(e event.Event) {

		if obj, err := baseobject.Discover(e.JSObject()); err == nil {

			if m, ok := obj.(messageevent.MessageEventFrom); ok {
				handler(m.MessageEvent_())
			}
		}
	})
}

// OnOpen Set onOpen Handler
func (e EventSource) OnOpen(handler func(e event.Event)) (js.Func, error) {

	return e.AddEventListener("open", handler)
}

// OnClose Set onClose Handler
func (e EventSource) OnClose(handler func(e event.Event)) (js.Func, error) {

	return e.AddEventListener("close", handler)
}

// OnError Set onError Handler
func (e EventSource) OnError(handler func(e event.Event)) (js.Func, error) {

	return e.AddEventListener("error", handler)
}
