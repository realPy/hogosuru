package blob

// Full implemented
// https://developer.mozilla.org/fr/docs/Web/API/Blob

import (
	"sync"

	"syscall/js"

	"github.com/realPy/hogosuru/arraybuffer"
	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/promise"
	"github.com/realPy/hogosuru/stream"
	readablestream "github.com/realPy/hogosuru/stream"
)

var singleton sync.Once

var blobinterface js.Value

//GetInterface get the JS interface Blob
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if blobinterface, err = baseobject.Get(js.Global(), "Blob"); err != nil {
			blobinterface = js.Undefined()
		}
		//autodiscover
		arraybuffer.GetInterface()
		baseobject.Register(blobinterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return blobinterface
}

type Blob struct {
	baseobject.BaseObject
}

type BlobFrom interface {
	Blob_() Blob
}

func (b Blob) Blob_() Blob {
	return b
}

func New(values ...interface{}) (Blob, error) {

	var b Blob
	var obj js.Value
	var err error
	var arrayJS []interface{}

	for _, value := range values {
		arrayJS = append(arrayJS, baseobject.GetJsValueOf(value))
	}

	if bi := GetInterface(); !bi.IsUndefined() {

		if obj, err = baseobject.New(bi, arrayJS); err == nil {
			b.BaseObject = b.SetObject(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return b, err
}

func NewWithObject(o js.Value) (Blob, error) {

	var b Blob
	var obj js.Value
	var err error
	if bi := GetInterface(); !bi.IsUndefined() {

		if obj, err = baseobject.New(bi, o); err == nil {
			b.BaseObject = b.SetObject(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return b, err
}

func NewWithArrayBuffer(a arraybuffer.ArrayBuffer) (Blob, error) {

	var b Blob
	var obj js.Value
	var err error
	if bi := GetInterface(); !bi.IsUndefined() {

		if obj, err = baseobject.New(bi, []interface{}{a.JSObject()}); err == nil {
			b.BaseObject = b.SetObject(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return b, err
}

func NewFromJSObject(obj js.Value) (Blob, error) {
	var b Blob
	var err error
	if bi := GetInterface(); !bi.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(bi) {
				b.BaseObject = b.SetObject(obj)

			} else {
				err = ErrNotABlob
			}
		}
	} else {
		err = ErrNotImplemented
	}

	return b, err
}

func (b Blob) IsClosed() (bool, error) {
	var err error
	var obj js.Value
	var ret bool

	if obj, err = b.Get("isClosed"); err == nil {
		if !obj.IsUndefined() {
			ret = obj.Bool()
		} else {
			err = baseobject.ErrNotImplementedFunc
		}

	}
	return ret, err
}

func (b Blob) Size() (int64, error) {

	return b.GetAttributeInt64("size")
}
func (b Blob) Type() (string, error) {
	var err error
	var obj js.Value

	if obj, err = b.Get("type"); err == nil {

		return obj.String(), nil
	}
	return "", err
}

func (b Blob) Close() error {
	_, err := b.Call("close")
	return err
}

func (b Blob) Slice(begin, end int64) (Blob, error) {
	var blob js.Value
	var err error
	if blob, err = b.Call("slice", js.ValueOf(begin), js.ValueOf(end)); err == nil {
		var newblob Blob
		object := newblob.SetObject(blob)
		newblob.BaseObject = object
		return newblob, nil
	}
	return Blob{}, err
}

func (b Blob) Stream() (stream.ReadableStream, error) {

	var err error
	var obj js.Value

	if obj, err = b.Call("stream"); err == nil {
		return stream.NewFromJSObject(obj)

	}
	return readablestream.ReadableStream{}, err
}

func (b Blob) ArrayBuffer() (promise.Promise, error) {

	var err error
	var promisebuffer js.Value
	var p promise.Promise

	if promisebuffer, err = b.Call("arrayBuffer"); err == nil {

		p, err = promise.NewFromJSObject(promisebuffer)

	}

	return p, err
}

func (b Blob) Text() (promise.Promise, error) {
	var err error
	var promisetext js.Value
	var p promise.Promise

	if promisetext, err = b.Call("text"); err == nil {
		p, err = promise.NewFromJSObject(promisetext)
	}

	return p, err
}
