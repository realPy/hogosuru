package typedarray

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
)

var singletonFloat32array sync.Once

var Float32arrayinterface js.Value

//Float32Array struct
type Float32Array struct {
	TypedArray
}

type Float32ArrayFrom interface {
	Float32Array_() Float32Array
}

func (u Float32Array) Float32Array_() Float32Array {
	return u
}

//GetFloat32ArrayInterface get the JS interface of Float32Array
func GetFloat32ArrayInterface() js.Value {

	singletonFloat32array.Do(func() {

		var err error
		if Float32arrayinterface, err = js.Global().GetWithErr("Float32Array"); err != nil {
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
	var obj interface{}

	if ai := GetFloat32ArrayInterface(); !ai.IsUndefined() {
		if objGo, ok := value.(baseobject.ObjectFrom); ok {
			obj = objGo.JSObject()
		} else {
			obj = js.ValueOf(value)
		}

		a.BaseObject = a.SetObject(ai.New(obj))
		return a, nil
	}

	return a, ErrNotImplementedFloat32Array
}

func NewFloat32FromJSObject(obj js.Value) (Float32Array, error) {
	var u Float32Array

	if ui := GetFloat32ArrayInterface(); !ui.IsUndefined() {
		if obj.InstanceOf(ui) {
			u.BaseObject = u.SetObject(obj)
			return u, nil
		}
	}

	return u, ErrNotAFloat32Array
}
