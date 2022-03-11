package clipboard

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/clipboarditem"
	"github.com/realPy/hogosuru/eventtarget"
	"github.com/realPy/hogosuru/initinterface"
	"github.com/realPy/hogosuru/promise"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var clipboardinterface js.Value

//BroadcastChannel struct
type Clipboard struct {
	eventtarget.EventTarget
}

type ClipboardFrom interface {
	Clipboard_() Clipboard
}

func (c Clipboard) Clipboard_() Clipboard {
	return c
}

//GetJSInterface get the JS interface
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if clipboardinterface, err = baseobject.Get(js.Global(), "Clipboard"); err != nil {
			clipboardinterface = js.Undefined()
		}

		baseobject.Register(clipboardinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)

		})

		clipboarditem.GetInterface()
		promise.GetInterface()

	})

	return clipboardinterface
}

func NewFromJSObject(obj js.Value) (Clipboard, error) {
	var c Clipboard

	if ci := GetInterface(); !ci.IsUndefined() {
		if obj.InstanceOf(ci) {
			c.BaseObject = c.SetObject(obj)
			return c, nil

		}
	}

	return c, ErrNotImplemented
}

func (c Clipboard) Read() (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj, err = c.Call("read"); err == nil {

		p, err = promise.NewFromJSObject(obj)
	}

	return p, err
}

func (c Clipboard) ReadText() (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj, err = c.Call("readText"); err == nil {

		p, err = promise.NewFromJSObject(obj)
	}

	return p, err
}

func (c Clipboard) Write(data []clipboarditem.ClipboardItem) (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise
	var arrayJS []interface{}

	for _, value := range data {
		arrayJS = append(arrayJS, value.JSObject())
	}

	if obj, err = c.Call("write", arrayJS); err == nil {

		p, err = promise.NewFromJSObject(obj)
	}

	return p, err
}

func (c Clipboard) WriteText(data string) (promise.Promise, error) {
	var err error
	var obj js.Value
	var p promise.Promise

	if obj, err = c.Call("writeText", js.ValueOf(data)); err == nil {

		p, err = promise.NewFromJSObject(obj)
	}

	return p, err
}
