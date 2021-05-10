package websocket

//https://developer.mozilla.org/fr/docs/Web/API/WebSocket

import (
	"sync"

	"github.com/realPy/hogosuru/arraybuffer"
	"github.com/realPy/hogosuru/blob"
	"github.com/realPy/hogosuru/js"
	"github.com/realPy/hogosuru/messageevent"

	"github.com/realPy/hogosuru/object"
)

var singleton sync.Once

var wsinterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//Websocket struct
type WebSocket struct {
	object.Object
}

//GetJSInterface get teh JS interface of broadcast channel
func GetJSInterface() *JSInterface {

	singleton.Do(func() {
		var wsinstance JSInterface
		var err error
		if wsinstance.objectInterface, err = js.Global().GetWithErr("WebSocket"); err == nil {
			wsinterface = &wsinstance
		}
	})

	return wsinterface
}

//New Get a new channel broadcast
func New(url string) (WebSocket, error) {
	var ws WebSocket

	if wsi := GetJSInterface(); wsi != nil {
		ws.Object = ws.SetObject(wsi.objectInterface.New(js.ValueOf(url)))
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

func (w WebSocket) SetOnMessage(handler func(WebSocket, interface{})) {

	w.setHandler("onmessage", func(ws WebSocket, v []js.Value) {

		if len(v) > 0 {

			if m, err := messageevent.NewFromJSObject(v[0]); err == nil {

				if data, err := m.Data(); err == nil {

					if btype, err := w.BinaryType(); err == nil {

						switch btype {
						case "blob":
							handler(w, data.String())
						case "arraybuffer":
							if a, err := arraybuffer.NewFromJSObject(data); err == nil {
								handler(w, a)
							}
						}

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
