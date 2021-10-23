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
	var obj interface{}

	if ai := GetInt16ArrayInterface(); !ai.IsUndefined() {
		if objGo, ok := value.(baseobject.ObjectFrom); ok {
			obj = objGo.JSObject()
		} else {
			obj = js.ValueOf(value)
		}

		a.BaseObject = a.SetObject(ai.New(obj))
		return a, nil
	}

	return a, ErrNotImplementedInt16Array
}

func NewInt16FromJSObject(obj js.Value) (Int16Array, error) {
	var u Int16Array

	if ui := GetInt16ArrayInterface(); !ui.IsUndefined() {
		if obj.InstanceOf(ui) {
			u.BaseObject = u.SetObject(obj)
			return u, nil
		}
	}

	return u, ErrNotAInt16Array
}
