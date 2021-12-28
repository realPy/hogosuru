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
	var objnew js.Value
	var err error
	if ai := GetFloat64ArrayInterface(); !ai.IsUndefined() {
		if objnew, err = baseobject.New(ai, baseobject.GetJsValueOf(value)); err == nil {
			a.BaseObject = a.SetObject(objnew)
		}

	} else {
		err = ErrNotImplementedFloat64Array
	}

	return a, err
}

func NewFloat64ArrayFrom(iterable interface{}) (Float64Array, error) {

	arr, err := newTypedArrayFrom(GetFloat64ArrayInterface(), func(v js.Value) (interface{}, error) {
		return NewFloat64FromJSObject(v)
	}, iterable)
	return arr.(Float64Array), err
}

func NewFloat64ArrayOf(values ...interface{}) (Float64Array, error) {

	arr, err := newTypedArrayOf(GetFloat64ArrayInterface(), func(v js.Value) (interface{}, error) {
		return NewFloat64FromJSObject(v)
	}, values...)
	return arr.(Float64Array), err
}

func NewFloat64FromJSObject(obj js.Value) (Float64Array, error) {
	var u Float64Array
	var err error
	if ui := GetFloat64ArrayInterface(); !ui.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(ui) {
				u.BaseObject = u.SetObject(obj)

			} else {
				err = ErrNotAFloat64Array
			}
		}
	} else {
		err = ErrNotImplementedFloat64Array
	}

	return u, err
}
