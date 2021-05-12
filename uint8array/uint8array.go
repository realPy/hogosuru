package uint8array

import (
	"errors"
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/arraybuffer"
	"github.com/realPy/hogosuru/object"
)

var singleton sync.Once

var uint8arrayinterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//Uint8Array struct
type Uint8Array struct {
	object.Object
}

//GetJSInterface get teh JS interface of broadcast channel
func GetJSInterface() *JSInterface {

	singleton.Do(func() {
		var uint8arrayinstance JSInterface
		var err error
		if uint8arrayinstance.objectInterface, err = js.Global().GetWithErr("Uint8Array"); err == nil {
			uint8arrayinterface = &uint8arrayinstance
		}
	})

	return uint8arrayinterface
}

/*
func (j *JSInterface) New(obj js.Value) js.Value {
	return j.objectInterface.New(obj)
}
*/

func NewFromArrayBuffer(a arraybuffer.ArrayBuffer) (Uint8Array, error) {

	uint8arrayObject := GetJSInterface().objectInterface.New(a.JSObject())
	return NewFromJSObject(uint8arrayObject)

}

func NewFromJSObject(obj js.Value) (Uint8Array, error) {
	var u Uint8Array

	u.Object = u.SetObject(obj)
	return u, nil

}

func (u Uint8Array) Bytes() ([]byte, error) {
	var err error
	var buffer []byte
	buffer = make([]byte, u.Length())
	if _, err = js.CopyBytesToGoWithErr(buffer, u.JSObject()); err == nil {
		return buffer, nil
	}
	return buffer, err
}

func (u Uint8Array) CopyBytes(buffer []byte) (int, error) {

	if len(buffer) < u.Length() {
		return 0, errors.New("Increase your buffer size")
	}

	return js.CopyBytesToGoWithErr(buffer, u.JSObject())

}

func (u Uint8Array) CopyFromBytes(buffer []byte) (int, error) {

	if len(buffer) < u.Length() {
		return 0, errors.New("Increase your buffer size")
	}

	return js.CopyBytesToJSWithErr(u.JSObject(), buffer)
}

func (u Uint8Array) ArrayBuffer() (arraybuffer.ArrayBuffer, error) {

	var err error
	var obj js.Value

	if obj, err = u.JSObject().GetWithErr("buffer"); err == nil {

		return arraybuffer.NewFromJSObject(obj)

	}

	return arraybuffer.ArrayBuffer{}, err
}
