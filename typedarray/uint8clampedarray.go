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
		if uint8clampedarrayinterface, err = baseobject.Get(js.Global(), "Uint8ClampedArray"); err != nil {
			uint8clampedarrayinterface = js.Undefined()
		}
		baseobject.Register(uint8clampedarrayinterface, func(v js.Value) (interface{}, error) {
			return NewUint8ClampedFromJSObject(v)
		})
	})

	return uint8clampedarrayinterface
}

func NewUint8ClampedArray(value interface{}) (Uint8ClampedArray, error) {

	var a Uint8ClampedArray
	var obj interface{}
	var objnew js.Value
	var err error

	if ai := GetUint8ClampedArrayInterface(); !ai.IsUndefined() {
		if objGo, ok := value.(baseobject.ObjectFrom); ok {
			obj = objGo.JSObject()
		} else {
			obj = js.ValueOf(value)
		}

		if objnew, err = baseobject.New(ai, obj); err == nil {
			a.BaseObject = a.SetObject(objnew)
		}

	} else {
		err = ErrNotImplementedUint8ClampedArray
	}

	return a, err
}

func NewUint8ClampedArrayFrom(iterable interface{}) (Uint8ClampedArray, error) {

	arr, err := newTypedArrayFrom(GetUint8ClampedArrayInterface(), func(v js.Value) (interface{}, error) {
		return NewUint8ClampedFromJSObject(v)
	}, iterable)
	return arr.(Uint8ClampedArray), err
}

func NewUint8ClampedArrayOf(values ...interface{}) (Uint8ClampedArray, error) {

	arr, err := newTypedArrayOf(GetUint8ClampedArrayInterface(), func(v js.Value) (interface{}, error) {
		return NewUint8ClampedFromJSObject(v)
	}, values...)
	return arr.(Uint8ClampedArray), err
}

func NewUint8ClampedFromJSObject(obj js.Value) (Uint8ClampedArray, error) {
	var u Uint8ClampedArray
	var err error
	if ui := GetUint8ClampedArrayInterface(); !ui.IsUndefined() {
		if obj.IsUndefined() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(ui) {
				u.BaseObject = u.SetObject(obj)

			} else {
				err = ErrNotAUint8ClampedArray
			}
		}
	} else {
		err = ErrNotImplementedUint8ClampedArray
	}

	return u, err
}
