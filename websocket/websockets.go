package websocket

// https://developer.mozilla.org/fr/docs/Web/API/WebSocket

import (
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/arraybuffer"
	"github.com/realPy/hogosuru/blob"
	"github.com/realPy/hogosuru/event"
	"github.com/realPy/hogosuru/eventtarget"
	"github.com/realPy/hogosuru/messageevent"

	"github.com/realPy/hogosuru/baseobject"
)

var singleton sync.Once

var wsinterface js.Value

//Websocket struct
type WebSocket struct {
	eventtarget.EventTarget
}

type WebSocketFrom interface {
	WebSocket_() WebSocket
}

func (w WebSocket) WebSocket_() WebSocket {
	return w
}

const (
	BlobType        = "blob"
	ArrayBufferType = "arraybuffer"
)

//GetInterface get the JS interface
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if wsinterface, err = baseobject.Get(js.Global(), "WebSocket"); err != nil {
			wsinterface = js.Undefined()
		}

		messageevent.GetInterface()
		baseobject.Register(wsinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return wsinterface
}

func NewFromJSObject(obj js.Value) (WebSocket, error) {
	var w WebSocket
	var err error
	if si := GetInterface(); !si.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(si) {
				w.BaseObject = w.SetObject(obj)

			} else {
				err = ErrNotAWebSocket
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return w, err
}

//New Get a new channel broadcast
func New(url string) (WebSocket, error) {
	var ws WebSocket
	var err error
	var obj js.Value

	if wsi := GetInterface(); !wsi.IsUndefined() {
		if obj, err = baseobject.New(wsi, js.ValueOf(url)); err == nil {
			ws.BaseObject = ws.SetObject(obj)
		}
	} else {
		err = ErrNotImplemented
	}
	return ws, err
}

func (w WebSocket) setHandler(jshandlername string, handler func(e event.Event)) {

	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		if e, err := event.NewFromJSObject(args[0]); err == nil {
			handler(e)
		}

		return nil
	})

	w.JSObject().Set(jshandlername, jsfunc)
}

//SetOnOpen Set onOpen Handler
func (w WebSocket) SetOnOpen(handler func(e event.Event)) {

	w.setHandler("onopen", func(e event.Event) {
		handler(e)
	})
}

//SetOnClose Set onClose Handler
func (w WebSocket) SetOnClose(handler func(e event.Event)) {
	w.setHandler("onclose", func(e event.Event) {
		handler(e)
	})
}

//SetOnClose Set onClose Handler
func (w WebSocket) SetOnError(handler func(e event.Event)) {
	w.setHandler("onerror", func(e event.Event) {
		handler(e)
	})
}

//SetOnClose Set onClose Handler
func (w WebSocket) SetOnMessage(handler func(e messageevent.MessageEvent)) {
	w.setHandler("onmessage", func(e event.Event) {

		if obj, err := baseobject.Discover(e.JSObject()); err == nil {

			if m, ok := obj.(messageevent.MessageEventFrom); ok {
				handler(m.MessageEvent_())
			}
		}
	})
}

//OnOpen Set onOpen Handler
func (w WebSocket) OnOpen(handler func(e event.Event)) (js.Func, error) {

	return w.AddEventListener("open", handler)
}

//OnClose Set onClose Handler
func (w WebSocket) OnClose(handler func(e event.Event)) (js.Func, error) {

	return w.AddEventListener("close", handler)
}

//OnError Set onError Handler
func (w WebSocket) OnError(handler func(e event.Event)) (js.Func, error) {

	return w.AddEventListener("error", handler)
}

func (w WebSocket) BinaryType() (string, error) {

	var err error
	var obj js.Value
	if obj, err = w.Get("binaryType"); err == nil {

		return obj.String(), nil
	}
	return "", err

}

func (w WebSocket) SetBinaryType(binaryType string) error {

	switch binaryType {
	case BlobType:
	case ArrayBufferType:
	default:
		return ErrSetBadBinaryType
	}

	w.JSObject().Set("binaryType", js.ValueOf(binaryType))

	return nil

}

//OnError Set onError Handler
func (w WebSocket) OnMessage(handler func(m messageevent.MessageEvent)) (js.Func, error) {

	return w.AddEventListener("message", func(e event.Event) {

		if obj, err := baseobject.Discover(e.JSObject()); err == nil {
			if m, ok := obj.(messageevent.MessageEventFrom); ok {
				handler(m.MessageEvent_())
			}
		}
	})
}

func (w WebSocket) Send(data interface{}) error {
	var object js.Value

	var err error
	switch d := data.(type) {
	case string:
		object = js.ValueOf(d)
	case arraybuffer.ArrayBuffer:
		object = d.JSObject()
	case blob.Blob:
		object = d.JSObject()
	default:
		err = ErrSendUnknownType
	}

	_, err = w.Call("send", object)

	return err
}

func (w WebSocket) Close() error {

	var err error
	_, err = w.Call("close")
	return err
}

func (w WebSocket) Protocol() (string, error) {

	var err error
	var obj js.Value
	if obj, err = w.Get("protocol"); err == nil {

		return obj.String(), nil
	}
	return "", err

}

func (w WebSocket) BufferedAmount() (int, error) {
	var err error
	var obj js.Value
	if obj, err = w.Get("bufferedAmount"); err == nil {

		return obj.Int(), nil
	}
	return 0, err
}

func (w WebSocket) ReadyState() (int, error) {
	var err error
	var obj js.Value
	if obj, err = w.Get("readyState"); err == nil {

		return obj.Int(), nil
	}
	return 0, err
}

func (w WebSocket) Url() (string, error) {

	var err error
	var obj js.Value
	if obj, err = w.Get("url"); err == nil {

		return obj.String(), nil
	}
	return "", err

}
