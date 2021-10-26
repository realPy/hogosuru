package typedarray

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
)

var singletonuint8array sync.Once

var uint8arrayinterface js.Value

//Uint8Array struct
type Uint8Array struct {
	TypedArray
}

type Uint8ArrayFrom interface {
	Uint8Array_() Uint8Array
}

func (u Uint8Array) Uint8Array_() Uint8Array {
	return u
}

//GetInterface get teh JS interface of broadcast channel
func GetUint8ArrayInterface() js.Value {

	singletonuint8array.Do(func() {

		var err error
		if uint8arrayinterface, err = baseobject.Get(js.Global(), "Uint8Array"); err != nil {
			uint8arrayinterface = js.Undefined()
		}
		baseobject.Register(uint8arrayinterface, func(v js.Value) (interface{}, error) {
			return NewUint8FromJSObject(v)
		})

	})

	return uint8arrayinterface
}

func NewUint8Array(value interface{}) (Uint8Array, error) {

	var a Uint8Array
	var obj interface{}

	if ai := GetUint8ArrayInterface(); !ai.IsUndefined() {
		if objGo, ok := value.(baseobject.ObjectFrom); ok {
			obj = objGo.JSObject()
		} else {
			obj = js.ValueOf(value)
		}

		a.BaseObject = a.SetObject(ai.New(obj))
		return a, nil
	}

	return a, ErrNotImplementedUint8Array
}

func NewUint8ArrayFrom(iterable interface{}) (Uint8Array, error) {

	arr, err := newTypedArrayFrom(GetUint8ArrayInterface(), func(v js.Value) (interface{}, error) {
		return NewUint8FromJSObject(v)
	}, iterable)
	return arr.(Uint8Array), err
}

func NewUint8ArrayOf(values ...interface{}) (Uint8Array, error) {

	arr, err := newTypedArrayOf(GetUint8ArrayInterface(), func(v js.Value) (interface{}, error) {
		return NewUint8FromJSObject(v)
	}, values...)
	return arr.(Uint8Array), err
}

func NewUint8FromJSObject(obj js.Value) (Uint8Array, error) {
	var u Uint8Array

	if ui := GetUint8ArrayInterface(); !ui.IsUndefined() {
		if obj.InstanceOf(ui) {
			u.BaseObject = u.SetObject(obj)
			return u, nil
		}
	}

	return u, ErrNotAUint8Array
}
