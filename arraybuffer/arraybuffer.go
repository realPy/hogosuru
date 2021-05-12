package arraybuffer

//partial implemented (herited from function)
// https://developer.mozilla.org/fr/docs/Web/JavaScript/Reference/Global_Objects/ArrayBuffer

import (
	"sync"

	"github.com/realPy/hogosuru/js"
	"github.com/realPy/hogosuru/object"
)

var singleton sync.Once

var arraybufferinterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//GetJSInterface get teh JS interface of broadcast channel
func GetJSInterface() *JSInterface {

	singleton.Do(func() {
		var arraybufferinstance JSInterface
		var err error
		if arraybufferinstance.objectInterface, err = js.Global().GetWithErr("ArrayBuffer"); err == nil {
			arraybufferinterface = &arraybufferinstance
		}
	})

	return arraybufferinterface
}

//ArrayBuffer struct
type ArrayBuffer struct {
	object.Object
}

func New(size int) (ArrayBuffer, error) {

	var a ArrayBuffer
	if ai := GetJSInterface(); ai != nil {

		a.Object = a.SetObject(ai.objectInterface.New(js.ValueOf(size)))
		return a, nil
	}

	return a, ErrNotImplemented
}

func NewFromJSObject(obj js.Value) (ArrayBuffer, error) {
	var a ArrayBuffer

	if object.String(obj) == "[object ArrayBuffer]" {
		a.Object = a.SetObject(obj)
		return a, nil
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
			return 0, object.ErrObjectNotNumber
		}

	}
	return 0, err

}
