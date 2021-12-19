package webassembly

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/arraybuffer"
	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/promise"
)

var singleton sync.Once

var webassemblyinterface js.Value

//GetInterface get the JS interface
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

//WebAssembly struct
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
	var wi js.Value
	if wi = GetInterface(); wi.IsUndefined() {
		return w, ErrNotImplemented
	}
	if obj.IsUndefined() || obj.IsNull() {
		return w, baseobject.ErrUndefinedValue
	}
	if !obj.InstanceOf(wi) {
		return w, ErrNotAWebAssembly
	}
	w.BaseObject = w.SetObject(obj)
	return w, nil
}

func (w WebAssembly) InstantiateStreaming(source promise.Promise, imports js.Value) (promise.Promise, error) {
	var obj js.Value
	var err error
	var p promise.Promise
	if obj, err = w.Call("instantiateStreaming", source.JSObject(), imports); err == nil {
		return promise.NewFromJSObject(obj)
	}
	return p, err
}

func (w WebAssembly) Instantiate(source arraybuffer.ArrayBuffer, imports js.Value) (promise.Promise, error) {
	var obj js.Value
	var err error
	var p promise.Promise
	if obj, err = w.Call("instantiate", source.JSObject(), imports); err == nil {
		return promise.NewFromJSObject(obj)
	}
	return p, err
}
