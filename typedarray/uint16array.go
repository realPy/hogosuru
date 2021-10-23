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

//GetInterface get teh JS interface of broadcast channel
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
	var obj interface{}

	if ai := GetUint16ArrayInterface(); !ai.IsUndefined() {
		if objGo, ok := value.(baseobject.ObjectFrom); ok {
			obj = objGo.JSObject()
		} else {
			obj = js.ValueOf(value)
		}

		a.BaseObject = a.SetObject(ai.New(obj))
		return a, nil
	}

	return a, ErrNotImplementedUint16Array
}

func NewUint16FromJSObject(obj js.Value) (Uint16Array, error) {
	var u Uint16Array

	if ui := GetUint16ArrayInterface(); !ui.IsUndefined() {
		if obj.InstanceOf(ui) {
			u.BaseObject = u.SetObject(obj)
			return u, nil
		}
	}

	return u, ErrNotAUint16Array
}
