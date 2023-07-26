package typedarray

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/base/baseobject"
)

var singletonFloat32array sync.Once

var Float32arrayinterface js.Value

// Float32Array struct
type Float32Array struct {
	TypedArray
}

type Float32ArrayFrom interface {
	Float32Array_() Float32Array
}

func (u Float32Array) Float32Array_() Float32Array {
	return u
}

// GetFloat32ArrayInterface get the JS interface of Float32Array
func GetFloat32ArrayInterface() js.Value {

	singletonFloat32array.Do(func() {

		var err error
		if Float32arrayinterface, err = baseobject.Get(js.Global(), "Float32Array"); err != nil {
			Float32arrayinterface = js.Undefined()
		}
		baseobject.Register(Float32arrayinterface, func(v js.Value) (interface{}, error) {
			return NewFloat32FromJSObject(v)
		})
	})

	return Float32arrayinterface
}

func NewFloat32Array(value interface{}) (Float32Array, error) {

	var a Float32Array
	var objnew js.Value
	var err error
	if ai := GetFloat32ArrayInterface(); !ai.IsUndefined() {
		if objnew, err = baseobject.New(ai, baseobject.GetJsValueOf(value)); err == nil {
			a.BaseObject = a.SetObject(objnew)
		}

	} else {
		err = ErrNotImplementedFloat32Array
	}

	return a, err
}

func NewFloat32ArrayFrom(iterable interface{}) (Float32Array, error) {

	arr, err := newTypedArrayFrom(GetFloat32ArrayInterface(), func(v js.Value) (interface{}, error) {
		return NewFloat32FromJSObject(v)
	}, iterable)
	return arr.(Float32Array), err
}

func NewFloat32ArrayOf(values ...interface{}) (Float32Array, error) {

	arr, err := newTypedArrayOf(GetFloat32ArrayInterface(), func(v js.Value) (interface{}, error) {
		return NewFloat32FromJSObject(v)
	}, values...)
	return arr.(Float32Array), err
}

func NewFloat32FromJSObject(obj js.Value) (Float32Array, error) {
	var u Float32Array
	var err error
	if ui := GetFloat32ArrayInterface(); !ui.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(ui) {
				u.BaseObject = u.SetObject(obj)

			} else {
				err = ErrNotAFloat32Array
			}
		}
	} else {
		err = ErrNotImplementedFloat32Array
	}

	return u, err
}
