package typedarray

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
)

var singletonint8array sync.Once

var int8arrayinterface js.Value

//Uint8Array struct
type Int8Array struct {
	TypedArray
}

type Int8ArrayFrom interface {
	Uint8Array_() Uint8Array
}

func (a Int8Array) Int8Array_() Int8Array {
	return a
}

//GetInterface get teh JS interface of broadcast channel
func GetInt8ArrayInterface() js.Value {

	singletonint8array.Do(func() {

		var err error
		if int8arrayinterface, err = baseobject.Get(js.Global(), "Int8Array"); err != nil {
			int8arrayinterface = js.Undefined()
		}
		baseobject.Register(int8arrayinterface, func(v js.Value) (interface{}, error) {
			return NewInt8FromJSObject(v)
		})

	})

	return int8arrayinterface
}

func NewInt8Array(value interface{}) (Int8Array, error) {

	var a Int8Array
	var obj interface{}
	if ai := GetInt8ArrayInterface(); !ai.IsUndefined() {

		if objGo, ok := value.(baseobject.ObjectFrom); ok {
			obj = objGo.JSObject()
		} else {
			obj = js.ValueOf(value)
		}

		a.BaseObject = a.SetObject(ai.New(obj))
		return a, nil
	}

	return a, ErrNotImplementedInt8Array
}

func NewInt8ArrayFrom(iterable interface{}) (Int8Array, error) {

	arr, err := newTypedArrayFrom(GetInt8ArrayInterface(), func(v js.Value) (interface{}, error) {
		return NewInt8FromJSObject(v)
	}, iterable)
	return arr.(Int8Array), err

}

func NewInt8ArrayOf(values ...interface{}) (Int8Array, error) {

	arr, err := newTypedArrayOf(GetInt8ArrayInterface(), func(v js.Value) (interface{}, error) {
		return NewInt8FromJSObject(v)
	}, values...)
	return arr.(Int8Array), err
}

func NewInt8FromJSObject(obj js.Value) (Int8Array, error) {
	var u Int8Array

	if ui := GetInt8ArrayInterface(); !ui.IsUndefined() {
		if obj.InstanceOf(ui) {
			u.BaseObject = u.SetObject(obj)
			return u, nil
		}
	}

	return u, ErrNotAInt8Array
}
