package webassembly

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/base/arraybuffer"
	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/initinterface"
	"github.com/realPy/hogosuru/base/promise"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var webassemblyinterface js.Value

// GetInterface get the JS interface
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if webassemblyinterface, err = baseobject.Get(js.Global(), "WebAssembly"); err != nil {
			webassemblyinterface = js.Undefined()
		}

		baseobject.Register(webassemblyinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})

	})

	return webassemblyinterface
}

// WebAssembly struct
type WebAssembly struct {
	baseobject.BaseObject
}

type WebAssemblyFrom interface {
	WebAssembly_() WebAssembly
}

func (w WebAssembly) WebAssembly_() WebAssembly {
	return w
}

func New() (WebAssembly, error) {

	var w WebAssembly

	if wi := GetInterface(); !wi.IsUndefined() {

		w.BaseObject = w.SetObject(wi)
		return w, nil
	}
	return w, ErrNotImplemented
}

func NewFromJSObject(obj js.Value) (WebAssembly, error) {
	var w WebAssembly
	var err error
	if wi := GetInterface(); !wi.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(wi) {
				w.BaseObject = w.SetObject(obj)
			} else {
				err = ErrNotAWebAssembly
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return w, err
}

func (w WebAssembly) InstantiateStreaming(source interface{}, imports js.Value) (promise.Promise, error) {

	var obj js.Value
	var err error
	var p promise.Promise

	if s, ok := source.(baseobject.ObjectFrom); ok {
		if obj, err = w.Call("instantiateStreaming", s.JSObject(), imports); err == nil {
			p, err = promise.NewFromJSObject(obj)

		}
	}

	return p, err
}

func (w WebAssembly) Instantiate(source arraybuffer.ArrayBuffer, imports js.Value) (promise.Promise, error) {
	var obj js.Value
	var err error
	var p promise.Promise

	if obj, err = w.Call("instantiate", source.JSObject(), imports); err == nil {
		p, err = promise.NewFromJSObject(obj)

	}
	return p, err
}
