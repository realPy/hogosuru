package typedarray

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
)

var singletonfloat64array sync.Once

var float64arrayinterface js.Value

//float64Array struct
type float64Array struct {
	TypedArray
}

type float64ArrayFrom interface {
	float64Array_() float64Array
}

func (u float64Array) float64Array_() float64Array {
	return u
}

//Getfloat64ArrayInterface get the JS interface of float64Array
func Getfloat64ArrayInterface() js.Value {

	singletonfloat64array.Do(func() {

		var err error
		if float64arrayinterface, err = js.Global().GetWithErr("float64Array"); err != nil {
			float64arrayinterface = js.Undefined()
		}
		baseobject.Register(float64arrayinterface, func(v js.Value) (interface{}, error) {
			return Newfloat64FromJSObject(v)
		})
	})

	return float64arrayinterface
}

func Newfloat64Array(value interface{}) (float64Array, error) {

	var a float64Array
	var obj interface{}

	if ai := Getfloat64ArrayInterface(); !ai.IsUndefined() {
		if objGo, ok := value.(baseobject.ObjectFrom); ok {
			obj = objGo.JSObject()
		} else {
			obj = js.ValueOf(value)
		}

		a.BaseObject = a.SetObject(ai.New(obj))
		return a, nil
	}

	return a, ErrNotImplementedFloat64Array
}

func Newfloat64FromJSObject(obj js.Value) (float64Array, error) {
	var u float64Array

	if ui := Getfloat64ArrayInterface(); !ui.IsUndefined() {
		if obj.InstanceOf(ui) {
			u.BaseObject = u.SetObject(obj)
			return u, nil
		}
	}

	return u, ErrNotAFloat64Array
}
