package typedarray

import (
	"errors"
	"syscall/js"

	"github.com/realPy/hogosuru/array"
	"github.com/realPy/hogosuru/arraybuffer"
)

//TypedArray struct
type TypedArray struct {
	array.Array
}

type TypedArrayFrom interface {
	TypedArray_() TypedArray
}

func (t TypedArray) TypedArray_() TypedArray {
	return t
}

func (t TypedArray) Bytes() ([]byte, error) {
	var err error
	var buffer []byte
	var l int
	if l, err = t.Length(); err == nil {
		buffer = make([]byte, l)
		if _, err = js.CopyBytesToGoWithErr(buffer, t.JSObject()); err == nil {
			return buffer, nil
		}
	}

	return buffer, err
}

func (t TypedArray) CopyBytes(buffer []byte) (int, error) {

	var err error
	var l int
	if l, err = t.Length(); err == nil {
		if len(buffer) < l {
			return 0, errors.New("Increase your buffer size")
		}

	} else {
		return 0, err
	}

	return js.CopyBytesToGoWithErr(buffer, t.JSObject())

}

func (t TypedArray) CopyFromBytes(buffer []byte) (int, error) {

	var err error
	var l int
	if l, err = t.Length(); err == nil {
		if len(buffer) < l {
			return 0, errors.New("Increase your buffer size")
		}

	} else {
		return 0, err
	}

	return js.CopyBytesToJSWithErr(t.JSObject(), buffer)
}

func (t TypedArray) Buffer() (arraybuffer.ArrayBuffer, error) {

	var err error
	var obj js.Value

	if obj, err = t.JSObject().GetWithErr("buffer"); err == nil {

		return arraybuffer.NewFromJSObject(obj)

	}

	return arraybuffer.ArrayBuffer{}, err
}

func (t TypedArray) ByteLength() (int64, error) {

	return t.GetAttributeInt64("byteLength")
}
