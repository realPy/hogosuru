package typedarray

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
)

var singletonFloat64array sync.Once

var Float64arrayinterface js.Value

//Float64Array struct
type Float64Array struct {
	TypedArray
}

type Float64ArrayFrom interface {
	Float64Array_() Float64Array
}

func (u Float64Array) Float64Array_() Float64Array {
	return u
}

//GetFloat64ArrayInterface get the JS interface of Float64Array
func GetFloat64ArrayInterface() js.Value {

	singletonFloat64array.Do(func() {

		var err error
		if Float64arrayinterface, err = baseobject.Get(js.Global(), "Float64Array"); err != nil {
			Float64arrayinterface = js.Undefined()
		}
		baseobject.Register(Float64arrayinterface, func(v js.Value) (interface{}, error) {
			return NewFloat64FromJSObject(v)
		})
	})

	return Float64arrayinterface
}

func NewFloat64Array(value interface{}) (Float64Array, error) {

	var a Float64Array
	var obj interface{}

	if ai := GetFloat64ArrayInterface(); !ai.IsUndefined() {
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

func NewFloat64FromJSObject(obj js.Value) (Float64Array, error) {
	var u Float64Array

	if ui := GetFloat64ArrayInterface(); !ui.IsUndefined() {
		if obj.InstanceOf(ui) {
			u.BaseObject = u.SetObject(obj)
			return u, nil
		}
	}

	return u, ErrNotAFloat64Array
}
