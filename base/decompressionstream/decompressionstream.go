package decompressionstream

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/base/initinterface"
	"github.com/realPy/hogosuru/base/response"
	"github.com/realPy/hogosuru/base/stream"
)

func init() {

	initinterface.RegisterInterface(GetInterface)
}

var singleton sync.Once

var decompressionstreaminterface js.Value

// GetJSInterface Get the JS Fetch Interface If nil browser doesn't implement it
func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if decompressionstreaminterface, err = baseobject.Get(js.Global(), "DecompressionStream"); err != nil {
			decompressionstreaminterface = js.Undefined()
		}

		response.GetInterface()
		baseobject.Register(decompressionstreaminterface, func(v js.Value) (interface{}, error) {
			return NewFromJSObject(v)
		})
	})

	return decompressionstreaminterface
}

// DecompressionStream struct
type DecompressionStream struct {
	baseobject.BaseObject
}

type DecompressionStreamFrom interface {
	DecompressionStream_() DecompressionStream
}

func (d DecompressionStream) DecompressionStream_() DecompressionStream {
	return d
}

func New(format string) (DecompressionStream, error) {
	var d DecompressionStream
	var err error
	var obj js.Value

	if decompressionstreami := GetInterface(); !decompressionstreami.IsUndefined() {

		if obj, err = baseobject.New(decompressionstreami, js.ValueOf(format)); err == nil {
			d.BaseObject = d.SetObject(obj)
		}

	} else {
		err = ErrNotImplemented
	}
	return d, err

}

func NewFromJSObject(obj js.Value) (DecompressionStream, error) {
	var d DecompressionStream
	var err error
	if decompressionstreami := GetInterface(); !decompressionstreami.IsUndefined() {
		if obj.IsUndefined() || obj.IsNull() {
			err = baseobject.ErrUndefinedValue
		} else {

			if obj.InstanceOf(decompressionstreami) {

				d.BaseObject = d.SetObject(obj)
			} else {
				err = ErrNotADecompressionStream
			}
		}
	} else {
		err = ErrNotImplemented
	}
	return d, err
}

func (d DecompressionStream) Readable() (stream.ReadableStream, error) {
	var err error
	var obj js.Value

	if obj, err = d.Get("readable"); err == nil {
		return stream.NewReadableStreamFromJSObject(obj)

	}
	return stream.ReadableStream{}, err
}

func (d DecompressionStream) Writable() (stream.WritableStream, error) {
	var err error
	var obj js.Value

	if obj, err = d.Get("writable"); err == nil {
		return stream.NewWriteableStreamFromJSObject(obj)

	}
	return stream.WritableStream{}, err
}
