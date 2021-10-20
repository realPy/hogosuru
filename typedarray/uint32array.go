package typedarray

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
)

var singletonuint32array sync.Once

var uint32arrayinterface js.Value

//Uint32Array struct
type Uint32Array struct {
	TypedArray
}

type Uint32ArrayFrom interface {
	Uint32Array_() Uint32Array
}

func (u Uint32Array) Uint32Array_() Uint32Array {
	return u
}

//GetInterface get teh JS interface of broadcast channel
func GetUint32ArrayInterface() js.Value {

	singletonuint32array.Do(func() {

		var err error
		if uint32arrayinterface, err = js.Global().GetWithErr("Uint32Array"); err != nil {
			uint32arrayinterface = js.Undefined()
		}
		baseobject.Register(uint32arrayinterface, func(v js.Value) (interface{}, error) {
			return NewUint32FromJSObject(v)
		})
	})

	return uint32arrayinterface
}

func NewUint32Array(value interface{}) (Uint32Array, error) {

	var a Uint32Array
	var obj interface{}

	if ai := GetUint32ArrayInterface(); !ai.IsUndefined() {
		if objGo, ok := value.(baseobject.ObjectFrom); ok {
			obj = objGo.JSObject()
		} else {
			obj = js.ValueOf(value)
		}

		a.BaseObject = a.SetObject(ai.New(obj))
		return a, nil
	}

	return a, ErrNotImplementedUint32Array
}

func NewUint32FromJSObject(obj js.Value) (Uint32Array, error) {
	var u Uint32Array

	if ui := GetUint32ArrayInterface(); !ui.IsUndefined() {
		if obj.InstanceOf(ui) {
			u.BaseObject = u.SetObject(obj)
			return u, nil
		}
	}

	return u, ErrNotAUint32Array
}
