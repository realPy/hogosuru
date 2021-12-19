package typedarray

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
)

var singletonuint16array sync.Once

var uint16arrayinterface js.Value

//Uint16Array struct
type Uint16Array struct {
	TypedArray
}

type Uint16ArrayFrom interface {
	Uint16Array_() Uint16Array
}

func (u Uint16Array) Uint16Array_() Uint16Array {
	return u
}

//GetInterface get the JS interface
func GetUint16ArrayInterface() js.Value {

	singletonuint16array.Do(func() {

		var err error
		if uint16arrayinterface, err = baseobject.Get(js.Global(), "Uint16Array"); err != nil {
			uint16arrayinterface = js.Undefined()
		}
		baseobject.Register(uint16arrayinterface, func(v js.Value) (interface{}, error) {
			return NewUint16FromJSObject(v)
		})
	})

	return uint16arrayinterface
}

func NewUint16Array(value interface{}) (Uint16Array, error) {

	var a Uint16Array
	var objnew js.Value
	var err error
	if ai := GetUint16ArrayInterface(); !ai.IsUndefined() {
		if objnew, err = baseobject.New(ai, baseobject.GetJsValueOf(value)); err == nil {
			a.BaseObject = a.SetObject(objnew)
		}

	} else {
		err = ErrNotImplementedUint16Array
	}

	return a, err
}

func NewUint16ArrayFrom(iterable interface{}) (Uint16Array, error) {

	arr, err := newTypedArrayFrom(GetUint16ArrayInterface(), func(v js.Value) (interface{}, error) {
		return NewUint16FromJSObject(v)
	}, iterable)
	return arr.(Uint16Array), err

}

func NewUint16ArrayOf(values ...interface{}) (Uint16Array, error) {

	arr, err := newTypedArrayOf(GetUint16ArrayInterface(), func(v js.Value) (interface{}, error) {
		return NewUint16FromJSObject(v)
	}, values...)
	return arr.(Uint16Array), err
}

func NewUint16FromJSObject(obj js.Value) (Uint16Array, error) {
	var u Uint16Array
	var err error
	if ui := GetUint16ArrayInterface(); !ui.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(ui) {
				u.BaseObject = u.SetObject(obj)

			} else {
				err = ErrNotAUint16Array
			}
		}
	} else {
		err = ErrNotImplementedUint16Array
	}

	return u, err
}
