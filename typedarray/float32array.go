package typedarray

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
)

var singletonfloat32array sync.Once

var float32arrayinterface js.Value

//float32Array struct
type float32Array struct {
	TypedArray
}

type float32ArrayFrom interface {
	float32Array_() float32Array
}

func (u float32Array) float32Array_() float32Array {
	return u
}

//Getfloat32ArrayInterface get the JS interface of float32Array
func Getfloat32ArrayInterface() js.Value {

	singletonfloat32array.Do(func() {

		var err error
		if float32arrayinterface, err = js.Global().GetWithErr("float32Array"); err != nil {
			float32arrayinterface = js.Undefined()
		}
		baseobject.Register(float32arrayinterface, func(v js.Value) (interface{}, error) {
			return Newfloat32FromJSObject(v)
		})
	})

	return float32arrayinterface
}

func Newfloat32Array(value interface{}) (float32Array, error) {

	var a float32Array
	var obj interface{}

	if ai := Getfloat32ArrayInterface(); !ai.IsUndefined() {
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

func Newfloat32FromJSObject(obj js.Value) (float32Array, error) {
	var u float32Array

	if ui := Getfloat32ArrayInterface(); !ui.IsUndefined() {
		if obj.InstanceOf(ui) {
			u.BaseObject = u.SetObject(obj)
			return u, nil
		}
	}

	return u, ErrNotAFloat32Array
}
