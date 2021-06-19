package arraybuffer

//partial implemented (herited from function)
// https://developer.mozilla.org/fr/docs/Web/JavaScript/Reference/Global_Objects/ArrayBuffer

import (
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
)

var singleton sync.Once

var arraybufferinterface js.Value

//GetInterface get teh JS interface of broadcast channel
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if arraybufferinterface, err = js.Global().GetWithErr("ArrayBuffer"); err != nil {
			arraybufferinterface = js.Null()
		}
	})
	baseobject.Register(arraybufferinterface, func(v js.Value) (interface{}, error) {
		return NewFromJSObject(v)
	})

	return arraybufferinterface
}

//ArrayBuffer struct
type ArrayBuffer struct {
	baseobject.BaseObject
}

func New(size int) (ArrayBuffer, error) {

	var a ArrayBuffer

	if ai := GetInterface(); !ai.IsNull() {

		a.BaseObject = a.SetObject(ai.New(js.ValueOf(size)))
		return a, nil
	}

	return a, ErrNotImplemented
}

func NewFromJSObject(obj js.Value) (ArrayBuffer, error) {
	var a ArrayBuffer

	if ai := GetInterface(); !ai.IsNull() {
		if obj.InstanceOf(ai) {
			a.BaseObject = a.SetObject(obj)
			return a, nil
		}
	}

	return a, ErrNotAnArrayBuffer
}

func (a ArrayBuffer) ByteLength() (int, error) {

	var byteLengthObject js.Value
	var err error
	if byteLengthObject, err = a.JSObject().GetWithErr("byteLength"); err == nil {
		if byteLengthObject.Type() == js.TypeNumber {
			return byteLengthObject.Int(), nil
		} else {
			return 0, baseobject.ErrObjectNotNumber
		}

	}
	return 0, err

}
