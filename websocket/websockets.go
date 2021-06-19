package websocket

// https://developer.mozilla.org/fr/docs/Web/API/WebSocket

import (
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/arraybuffer"
	"github.com/realPy/hogosuru/blob"
	"github.com/realPy/hogosuru/messageevent"

	"github.com/realPy/hogosuru/baseobject"
)

var singleton sync.Once

var wsinterface js.Value

//Websocket struct
type WebSocket struct {
	baseobject.BaseObject
}

const (
	BlobType        = "blob"
	ArrayBufferType = "arraybuffer"
)

//GetInterface get teh JS interface of broadcast channel
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if wsinterface, err = js.Global().GetWithErr("WebSocket"); err != nil {
			wsinterface = js.Null()
		}

	})
	baseobject.Register(wsinterface, func(v js.Value) (interface{}, error) {
		return NewFromJSObject(v)
	})
	return wsinterface
}

func NewFromJSObject(obj js.Value) (WebSocket, error) {
	var w WebSocket
	var err error
	if si := GetInterface(); !si.IsNull() {
		if obj.InstanceOf(si) {
			w.BaseObject = w.SetObject(obj)

		}
	} else {
		err = ErrNotAWebSocket
	}

	return w, err
}

//New Get a new channel broadcast
func New(url string) (WebSocket, error) {
	var ws WebSocket

	if wsi := GetInterface(); !wsi.IsNull() {
		ws.BaseObject = ws.SetObject(wsi.New(js.ValueOf(url)))
		return ws, nil
	}
	return ws, ErrNotImplemented
}

func (w WebSocket) setHandler(jshandlername string, handler func(WebSocket, []js.Value)) {

	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		handler(w, args)
		return nil
	})

	w.JSObject().Set(jshandlername, jsfunc)
}

//SetOnOpen Set onOpen Handler
func (w WebSocket) SetOnOpen(handler func(WebSocket)) {

	w.setHandler("onopen", func(ws WebSocket, v []js.Value) {
		handler(ws)
	})
}

//SetOnClose Set onClose Handler
func (w WebSocket) SetOnClose(handler func(WebSocket)) {

	w.setHandler("onclose", func(ws WebSocket, v []js.Value) {
		handler(ws)
	})
}

//SetOnError Set onError Handler
func (w WebSocket) SetOnError(handler func(WebSocket)) {

	w.setHandler("onerror", func(ws WebSocket, v []js.Value) {
		handler(ws)
	})
}

func (w WebSocket) BinaryType() (string, error) {

	var err error
	var obj js.Value
	if obj, err = w.JSObject().GetWithErr("binaryType"); err == nil {

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

func (w WebSocket) SetOnMessage(handler func(WebSocket, interface{})) {

	w.setHandler("onmessage", func(ws WebSocket, v []js.Value) {

		if len(v) > 0 {

			if m, err := messageevent.NewFromJSObject(v[0]); err == nil {

				if data, err := m.Data(); err == nil {
					switch baseobject.String(data) {
					case "[object Blob]":
						if b, err := blob.NewFromJSObject(data); err == nil {
							handler(w, b)
						}
					case "[object ArrayBuffer]":
						if a, err := arraybuffer.NewFromJSObject(data); err == nil {
							handler(w, a)
						}
					default:
						handler(w, data.String())
					}

				}

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

	_, err = w.JSObject().CallWithErr("send", object)

	return err
}

func (w WebSocket) Close() error {

	var err error
	_, err = w.JSObject().CallWithErr("close")
	return err
}

func (w WebSocket) Protocol() (string, error) {

	var err error
	var obj js.Value
	if obj, err = w.JSObject().GetWithErr("protocol"); err == nil {

		return obj.String(), nil
	}
	return "", err

}

func (w WebSocket) BufferedAmount() (int, error) {
	var err error
	var obj js.Value
	if obj, err = w.JSObject().GetWithErr("bufferedAmount"); err == nil {

		return obj.Int(), nil
	}
	return 0, err
}

func (w WebSocket) ReadyState() (int, error) {
	var err error
	var obj js.Value
	if obj, err = w.JSObject().GetWithErr("readyState"); err == nil {

		return obj.Int(), nil
	}
	return 0, err
}

func (w WebSocket) Url() (string, error) {

	var err error
	var obj js.Value
	if obj, err = w.JSObject().GetWithErr("url"); err == nil {

		return obj.String(), nil
	}
	return "", err

}
