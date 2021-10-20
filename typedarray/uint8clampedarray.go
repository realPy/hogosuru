package typedarray

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
)

var singletonuint8campledarray sync.Once

var uint8clampedarrayinterface js.Value

//Uint8ClampedArray struct
type Uint8ClampedArray struct {
	TypedArray
}

type Uint8ClampedArrayFrom interface {
	Uint8ClampedArray_() Uint8ClampedArray
}

func (u Uint8ClampedArray) Uint8ClampedArray_() Uint8ClampedArray {
	return u
}

//GetUint8ClampedArrayInterface get the JS interface of GetUint8ClampedArrayInterface
func GetUint8ClampedArrayInterface() js.Value {

	singletonuint8campledarray.Do(func() {

		var err error
		if uint8clampedarrayinterface, err = js.Global().GetWithErr("Uint8ClampedArray"); err != nil {
			uint8clampedarrayinterface = js.Undefined()
		}

	})

	return uint8clampedarrayinterface
}

func NewUint8ClampedArray(value interface{}) (Uint8ClampedArray, error) {

	var a Uint8ClampedArray
	var obj interface{}

	if ai := GetUint8ClampedArrayInterface(); !ai.IsUndefined() {
		if objGo, ok := value.(baseobject.ObjectFrom); ok {
			obj = objGo.JSObject()
		} else {
			obj = js.ValueOf(value)
		}

		a.BaseObject = a.SetObject(ai.New(obj))
		return a, nil
	}

	return a, ErrNotImplementedUint8ClampedArray
}

func NewUint8ClampedFromJSObject(obj js.Value) (Uint8ClampedArray, error) {
	var u Uint8ClampedArray

	if ui := GetUint8ClampedArrayInterface(); !ui.IsUndefined() {
		if obj.InstanceOf(ui) {
			u.BaseObject = u.SetObject(obj)
			return u, nil
		}
	}

	return u, ErrNotAUint8ClampedArray
}
