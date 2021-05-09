package blob

import (
	"sync"

	"github.com/realPy/jswasm/arraybuffer"
	"github.com/realPy/jswasm/js"
	"github.com/realPy/jswasm/object"
	"github.com/realPy/jswasm/stream"
	readablestream "github.com/realPy/jswasm/stream"
)

var singleton sync.Once

var blobinterface *JSInterface

//JSInterface JSInterface struct
type JSInterface struct {
	objectInterface js.Value
}

//GetJSInterface get teh JS interface of broadcast channel
func GetJSInterface() *JSInterface {

	singleton.Do(func() {
		var blobinstance JSInterface
		var err error
		if blobinstance.objectInterface, err = js.Global().GetWithErr("Blob"); err == nil {
			blobinterface = &blobinstance
		}
	})

	return blobinterface
}

type Blob struct {
	object.Object
}

func New() (Blob, error) {

	var b Blob
	if bi := GetJSInterface(); bi != nil {

		b.Object = b.SetObject(bi.objectInterface.New())
		return b, nil
	}
	return b, ErrNotImplemented
}

func NewWithObject(o js.Value) (Blob, error) {

	var b Blob
	if bi := GetJSInterface(); bi != nil {
		b.Object = b.SetObject(bi.objectInterface.New(o))
		return b, nil
	}
	return b, ErrNotImplemented
}

func NewWithArrayBuffer(a arraybuffer.ArrayBuffer) (Blob, error) {

	var b Blob
	if bi := GetJSInterface(); bi != nil {

		b.Object = b.SetObject(bi.objectInterface.New([]interface{}{a.JSObject()}))
		return b, nil
	}
	return b, ErrNotImplemented
}

/*
func NewWithUint8Array(u uint8array.Uint8Array) (Blob, error) {

	var b Blob
	if bi := GetJSInterface(); bi != nil {

		b.Object = b.SetObject(bi.objectInterface.New(u.JSObject()))
		return b, nil
	}
	return b, ErrNotImplemented
}
*/

func NewWithBlob(bl Blob) (Blob, error) {

	var b Blob
	if bi := GetJSInterface(); bi != nil {
		b.Object = b.SetObject(bi.objectInterface.New(bl.JSObject()))
		return b, nil
	}
	return b, ErrNotImplemented
}

func NewFromJSObject(obj js.Value) (Blob, error) {
	var b Blob

	if object.String(obj) == "[object Blob]" {
		b.Object = b.SetObject(obj)
		return b, nil
	}

	return b, ErrNotABlob
}

func (b Blob) IsClosed() (bool, error) {
	var err error
	var obj js.Value

	if obj, err = b.JSObject().GetWithErr("isClosed"); err == nil {

		return obj.Bool(), nil
	}
	return true, err
}

func (b Blob) Size() (int, error) {
	var err error
	var obj js.Value
	if obj, err = b.JSObject().GetWithErr("size"); err == nil {

		return obj.Int(), nil
	}
	return 0, err
}
func (b Blob) Type() (string, error) {
	var err error
	var obj js.Value

	if obj, err = b.JSObject().GetWithErr("type"); err == nil {

		return obj.String(), nil
	}
	return "", err
}

func (b Blob) Close() error {
	_, err := b.JSObject().CallWithErr("close")
	return err
}

func (b Blob) Slice(begin, end int) (Blob, error) {
	var blob js.Value
	var err error
	if blob, err = b.JSObject().CallWithErr("slice", js.ValueOf(begin), js.ValueOf(end)); err == nil {
		var newblob Blob
		object := newblob.SetObject(blob)
		newblob.Object = object
		return newblob, nil
	}
	return Blob{}, err
}

func (b Blob) Stream() (stream.ReadableStream, error) {

	var err error
	var obj js.Value

	if obj, err = b.JSObject().CallWithErr("stream"); err == nil {
		return stream.NewReadableStreamFromJSObject(obj)

	}
	return readablestream.ReadableStream{}, err
}

func (b Blob) ArrayBuffer() (arraybuffer.ArrayBuffer, error) {

	var err error
	var promisebuffer js.Value
	var arrayb arraybuffer.ArrayBuffer

	if promisebuffer, err = b.JSObject().CallWithErr("arrayBuffer"); err == nil {
		waitsync := make(chan arraybuffer.ArrayBuffer)
		then := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			if arrayb, err = arraybuffer.NewFromJSObject(args[0]); err == nil {
				waitsync <- arrayb
			} else {
				waitsync <- arraybuffer.ArrayBuffer{}
			}

			return nil
		})

		promisebuffer.Call("then", then)
		arrayb = <-waitsync

	}
	return arrayb, err
}

func (b Blob) Text() (string, error) {
	var err error
	var promisetext js.Value
	var text string = ""

	if promisetext, err = b.JSObject().CallWithErr("text"); err == nil {
		waitsync := make(chan string)
		then := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			waitsync <- args[0].String()
			return nil
		})

		promisetext.Call("then", then)
		text = <-waitsync
	}
	return text, err
}

func (b Blob) Append(append object.Object) (Blob, error) {

	var blobObject js.Value
	var arrayblob []interface{} = []interface{}{b.JSObject(), append.JSObject()}
	if bi := GetJSInterface(); bi != nil {
		blobObject = bi.objectInterface.New(arrayblob)

		return NewFromJSObject(blobObject)
	}
	return Blob{}, ErrNotImplemented
}
