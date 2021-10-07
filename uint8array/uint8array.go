package uint8array

import (
	"errors"
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/arraybuffer"
	"github.com/realPy/hogosuru/baseobject"
)

var singleton sync.Once

var uint8arrayinterface js.Value

//Uint8Array struct
type Uint8Array struct {
	baseobject.BaseObject
}

type Uint8ArrayFrom interface {
	Uint8Array_() Uint8Array
}

func (u Uint8Array) Uint8Array_() Uint8Array {
	return u
}

//GetInterface get teh JS interface of broadcast channel
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if uint8arrayinterface, err = js.Global().GetWithErr("Uint8Array"); err != nil {
			uint8arrayinterface = js.Null()
		}

	})

	return uint8arrayinterface
}

func NewFromArrayBuffer(a arraybuffer.ArrayBuffer) (Uint8Array, error) {
	var arr Uint8Array
	var err error
	if ai := GetInterface(); !ai.IsNull() {
		uint8arrayObject := GetInterface().New(a.JSObject())
		arr, err = NewFromJSObject(uint8arrayObject)

	} else {
		err = ErrNotImplemented
	}

	return arr, err

}

func NewFromJSObject(obj js.Value) (Uint8Array, error) {
	var u Uint8Array

	if ui := GetInterface(); !ui.IsNull() {
		if obj.InstanceOf(ui) {
			u.BaseObject = u.SetObject(obj)
			return u, nil
		}
	}

	return u, ErrNotAUint8Array
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
