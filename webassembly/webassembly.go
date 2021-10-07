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

//GetInterface get teh JS interface of broadcast channel
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if webassemblyinterface, err = js.Global().GetWithErr("WebAssembly"); err != nil {
			webassemblyinterface = js.Null()
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

	if wi := GetInterface(); !wi.IsNull() {

		w.BaseObject = w.SetObject(wi)
		return w, nil
	}
	return w, ErrNotImplemented
}

func NewFromJSObject(obj js.Value) (WebAssembly, error) {
	var w WebAssembly
	var err error
	if wi := GetInterface(); !wi.IsNull() {
		if obj.InstanceOf(wi) {
			w.BaseObject = w.SetObject(obj)
			return w, nil
		} else {
			err = ErrNotAWebAssembly
		}
	} else {
		err = ErrNotImplemented
	}

	return w, err
}

func (w WebAssembly) InstantiateStreaming(source promise.Promise, imports js.Value) (promise.Promise, error) {

	var obj js.Value
	var err error
	var p promise.Promise
	if obj, err = w.JSObject().CallWithErr("instantiateStreaming", source.JSObject(), imports); err == nil {
		p, err = promise.NewFromJSObject(obj)

	}
	return p, err
}

func (w WebAssembly) Instantiate(source arraybuffer.ArrayBuffer, imports js.Value) (baseobject.BaseObject, error) {
	var obj js.Value
	var err error
	var b baseobject.BaseObject

	if obj, err = w.JSObject().CallWithErr("instantiate", source.JSObject(), imports); err == nil {
		b, err = baseobject.NewFromJSObject(obj)

	}
	return b, err
}
