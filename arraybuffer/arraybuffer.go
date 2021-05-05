package arraybuffer

import (
	"errors"
	"sync"

	"github.com/realPy/jswasm/js"
	"github.com/realPy/jswasm/object"
	"github.com/realPy/jswasm/uint8array"
)

var (
	ErrNotAnArrayBuffer = errors.New("The given value must be an arrayBuffer")
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
		if arraybufferinstance.objectInterface, err = js.Global().GetWithErr("arrayBuffer"); err == nil {
			arraybufferinterface = &arraybufferinstance
		}
	})

	return arraybufferinterface
}

//ArrayBuffer struct
type ArrayBuffer struct {
	object js.Value
}

func (j *JSInterface) New(obj js.Value) js.Value {
	return j.objectInterface.New(obj)
}

func NewArrayBuffer(obj js.Value) (ArrayBuffer, error) {
	var arraybuffer ArrayBuffer

	if object.String(obj) == "[object ArrayBuffer]" {

		arraybuffer.object = obj
		return arraybuffer, nil
	}
	return arraybuffer, ErrNotAnArrayBuffer
}

func (a ArrayBuffer) ByteLength() (int, error) {

	var byteLengthObject js.Value
	var err error
	if byteLengthObject, err = a.object.GetWithErr("byteLength"); err == nil {
		if byteLengthObject.Type() == js.TypeNumber {
			return byteLengthObject.Int(), nil
		} else {
			return 0, object.ErrObjectNotNumber
		}

	}
	return 0, err

}

func (a ArrayBuffer) Bytes() ([]byte, error) {
	var err error
	var buffer []byte

	if arrayConstructor := uint8array.GetJSInterface(); arrayConstructor != nil {
		uint8arrayBuffer := arrayConstructor.New(a.object)
		len, _ := a.ByteLength()
		buffer = make([]byte, len)
		if _, err = js.CopyBytesToGoWithErr(buffer, uint8arrayBuffer); err == nil {
			return buffer, nil
		}
	}
	//convert arraybuffer JS TO uint8array
	return buffer, err
}
