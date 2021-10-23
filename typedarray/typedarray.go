package typedarray

import (
	"errors"
	"syscall/js"

	"github.com/realPy/hogosuru/array"
	"github.com/realPy/hogosuru/arraybuffer"
	"github.com/realPy/hogosuru/baseobject"
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
		if _, err = baseobject.CopyBytesToGo(buffer, t.JSObject()); err == nil {
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

	return baseobject.CopyBytesToGo(buffer, t.JSObject())

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

	return baseobject.CopyBytesToJS(t.JSObject(), buffer)
}

func (t TypedArray) Buffer() (arraybuffer.ArrayBuffer, error) {

	var err error
	var obj js.Value

	if obj, err = t.Get("buffer"); err == nil {

		return arraybuffer.NewFromJSObject(obj)

	}

	return arraybuffer.ArrayBuffer{}, err
}

func (t TypedArray) ByteLength() (int64, error) {

	return t.GetAttributeInt64("byteLength")
}

func (t TypedArray) ByteOffset() (int64, error) {

	return t.GetAttributeInt64("byteOffset")
}

func (t TypedArray) BYTES_PER_ELEMENT() (int, error) {

	return t.GetAttributeInt("BYTES_PER_ELEMENT")
}

func (t TypedArray) Subarray(opts ...int) (interface{}, error) {

	var err error
	var arrayJS []interface{}
	var obj js.Value
	var newArr interface{}

	if len(opts) < 3 {
		for _, opt := range opts {
			arrayJS = append(arrayJS, js.ValueOf(opt))
		}
	}

	if obj, err = t.Call("subarray", arrayJS...); err == nil {
		newArr, err = baseobject.Discover(obj)

	}

	return newArr, err
}
