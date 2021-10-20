package typedarray

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
)

var singletonint32array sync.Once

var int32arrayinterface js.Value

//Int32Array struct
type Int32Array struct {
	TypedArray
}

type Int32ArrayFrom interface {
	Int32Array_() Int32Array
}

func (u Int32Array) Int32Array_() Int32Array {
	return u
}

//GetInt32ArrayInterface get the JS interface of Int32Array
func GetInt32ArrayInterface() js.Value {

	singletonint32array.Do(func() {

		var err error
		if int32arrayinterface, err = js.Global().GetWithErr("Int32Array"); err != nil {
			int32arrayinterface = js.Undefined()
		}
		baseobject.Register(int32arrayinterface, func(v js.Value) (interface{}, error) {
			return NewInt32FromJSObject(v)
		})
	})

	return int32arrayinterface
}

func NewInt32Array(value interface{}) (Int32Array, error) {

	var a Int32Array
	var obj interface{}

	if ai := GetInt32ArrayInterface(); !ai.IsUndefined() {
		if objGo, ok := value.(baseobject.ObjectFrom); ok {
			obj = objGo.JSObject()
		} else {
			obj = js.ValueOf(value)
		}

		a.BaseObject = a.SetObject(ai.New(obj))
		return a, nil
	}

	return a, ErrNotImplementedInt32Array
}

func NewInt32FromJSObject(obj js.Value) (Int32Array, error) {
	var u Int32Array

	if ui := GetInt32ArrayInterface(); !ui.IsUndefined() {
		if obj.InstanceOf(ui) {
			u.BaseObject = u.SetObject(obj)
			return u, nil
		}
	}

	return u, ErrNotAInt32Array
}
