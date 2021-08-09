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

type ArrayBufferFrom interface {
	ArrayBuffer() ArrayBuffer
}

func (a ArrayBuffer) ArrayBuffer() ArrayBuffer {
	return a
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

	return a.GetAttributeInt("byteLength")
}
