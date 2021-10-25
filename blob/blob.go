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

//GetJSInterface get teh JS interface of broadcast channel
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

func New() (Blob, error) {

	var b Blob
	if bi := GetInterface(); !bi.IsUndefined() {

		b.BaseObject = b.SetObject(bi.New())
		return b, nil
	}
	return b, ErrNotImplemented
}

func NewWithObject(o js.Value) (Blob, error) {

	var b Blob
	if bi := GetInterface(); !bi.IsUndefined() {
		b.BaseObject = b.SetObject(bi.New(o))
		return b, nil
	}
	return b, ErrNotImplemented
}

func NewWithArrayBuffer(a arraybuffer.ArrayBuffer) (Blob, error) {

	var b Blob
	if bi := GetInterface(); !bi.IsUndefined() {

		b.BaseObject = b.SetObject(bi.New([]interface{}{a.JSObject()}))
		return b, nil
	}
	return b, ErrNotImplemented
}

func NewFromJSObject(obj js.Value) (Blob, error) {
	var b Blob

	if bi := GetInterface(); !bi.IsUndefined() {
		if obj.InstanceOf(bi) {
			b.BaseObject = b.SetObject(obj)
			return b, nil
		}
	}

	return b, ErrNotABlob
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

func (b Blob) ArrayBuffer() (arraybuffer.ArrayBuffer, error) {

	var err error
	var promisebuffer js.Value
	var arrayb arraybuffer.ArrayBuffer
	var p promise.Promise
	var binaryObj interface{}

	if promisebuffer, err = b.Call("arrayBuffer"); err == nil {

		if p, err = promise.NewFromJSObject(promisebuffer); err == nil {

			if binaryObj, err = p.Await(); err == nil {
				if binary, ok := binaryObj.(arraybuffer.ArrayBufferFrom); ok {
					arrayb = binary.ArrayBuffer_()
				} else {
					err = arraybuffer.ErrNotAnArrayBuffer
				}

			}
		}
	}

	return arrayb, err
}

func (b Blob) Text() (string, error) {
	var err error
	var promisetext js.Value
	var p promise.Promise
	var jsTxtObj interface{}
	var text string = ""

	if promisetext, err = b.Call("text"); err == nil {
		if p, err = promise.NewFromJSObject(promisetext); err == nil {

			if jsTxtObj, err = p.Await(); err == nil {

				if jsTxt, ok := jsTxtObj.(baseobject.ObjectFrom); ok {
					text = jsTxt.JSObject().String()
				} else {
					err = baseobject.ErrNotABaseObject
				}

			}

		}
	}

	return text, err
}

func (b Blob) Append(append baseobject.BaseObject) (Blob, error) {

	var blobObject js.Value
	var arrayblob []interface{} = []interface{}{b.JSObject(), append.JSObject()}
	if bi := GetInterface(); !bi.IsUndefined() {
		blobObject = bi.New(arrayblob)

		return NewFromJSObject(blobObject)
	}
	return Blob{}, ErrNotImplemented
}
