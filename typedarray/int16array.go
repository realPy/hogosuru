package typedarray

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
)

var singletonint16array sync.Once

var int16arrayinterface js.Value

//Int16Array struct
type Int16Array struct {
	TypedArray
}

type Int16ArrayFrom interface {
	Int16Array_() Int16Array
}

func (u Int16Array) Int16Array_() Int16Array {
	return u
}

//GetInt16ArrayInterface get the JS interface of Int16Array
func GetInt16ArrayInterface() js.Value {

	singletonint16array.Do(func() {

		var err error
		if int16arrayinterface, err = baseobject.Get(js.Global(), "Int16Array"); err != nil {
			int16arrayinterface = js.Undefined()
		}
		baseobject.Register(int16arrayinterface, func(v js.Value) (interface{}, error) {
			return NewInt16FromJSObject(v)
		})
	})

	return int16arrayinterface
}

func NewInt16Array(value interface{}) (Int16Array, error) {

	var a Int16Array
	var objnew js.Value
	var err error
	if ai := GetInt16ArrayInterface(); !ai.IsUndefined() {
		if objnew, err = baseobject.New(ai, baseobject.GetJsValueOf(value)); err == nil {
			a.BaseObject = a.SetObject(objnew)
		}

	} else {
		err = ErrNotImplementedInt16Array
	}

	return a, err
}

func NewInt16ArrayFrom(iterable interface{}) (Int16Array, error) {

	arr, err := newTypedArrayFrom(GetInt16ArrayInterface(), func(v js.Value) (interface{}, error) {
		return NewInt16FromJSObject(v)
	}, iterable)
	return arr.(Int16Array), err

}

func NewInt16ArrayOf(values ...interface{}) (Int16Array, error) {

	arr, err := newTypedArrayOf(GetInt16ArrayInterface(), func(v js.Value) (interface{}, error) {
		return NewInt16FromJSObject(v)
	}, values...)
	return arr.(Int16Array), err
}

func NewInt16FromJSObject(obj js.Value) (Int16Array, error) {
	var u Int16Array
	var err error
	if ui := GetInt16ArrayInterface(); !ui.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(ui) {
				u.BaseObject = u.SetObject(obj)

			} else {
				err = ErrNotAInt16Array

			}
		}

	} else {
		err = ErrNotImplementedInt16Array
	}

	return u, err
}
